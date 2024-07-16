package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseInboundService struct {
}

func NewStorehouseInboundService() *StorehouseInboundService {
	return &StorehouseInboundService{}
}

func (s *StorehouseInboundService) CreateInbound(ctx *app.Context, userId string, req *model.StorehouseInboundReq) error {
	nowstr := time.Now().Format("2006-01-02 15:04:05")
	inbound := &model.StorehouseInbound{
		StorehouseUuid: req.StorehouseUuid,
		Title:          req.Title,
		InboundType:    req.InboundType,
		Status:         req.Status,
		InboundOrderNo: uuid.New().String(),
		InboundDate:    time.Now().Format("2006-01-02"),
		InboundBy:      userId,
		CreatedAt:      nowstr,
		UpdatedAt:      nowstr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(inbound).Error; err != nil {
			return err
		}

		for _, detailReq := range req.Detail {
			detail := &model.StorehouseInboundDetail{
				InboundOrderNo: inbound.InboundOrderNo,
				ProductUuid:    detailReq.ProductUuid,
				SkuUuid:        detailReq.SkuUuid,
				Quantity:       detailReq.Quantity,
				CreatedAt:      nowstr,
				UpdatedAt:      nowstr,
			}
			if err := tx.Create(detail).Error; err != nil {
				return err
			}

			// 更新库存
			// 先获取库存
			stock := &model.StorehouseProduct{}
			err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", req.StorehouseUuid, detailReq.ProductUuid, detailReq.SkuUuid).First(stock).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					stock.Uuid = uuid.New().String()
					stock.StorehouseUuid = req.StorehouseUuid
					stock.ProductUuid = detailReq.ProductUuid
					stock.SkuUuid = detailReq.SkuUuid
					stock.Quantity = detailReq.Quantity
					stock.CreatedAt = nowstr
					stock.UpdatedAt = nowstr
					if err := tx.Create(stock).Error; err != nil {
						return err
					}

					// 创建库存记录
					stockopLog := &model.StorehouseProductOpLog{
						Uuid:                  uuid.New().String(),
						StorehouseProductUuid: stock.Uuid,
						StorehouseUuid:        inbound.StorehouseUuid,
						BeforeQuantity:        0,
						Quantity:              detail.Quantity,
						OpQuantity:            detail.Quantity,
						OpType:                1,
						OpDesc:                "仓库第一次入库",
						OpBy:                  userId,
						CreatedAt:             nowstr,
					}
					if err := tx.Create(stockopLog).Error; err != nil {
						ctx.Logger.Error("Failed to create stockop log", err)
						return errors.New("failed to create stockop log")
					}

				}
			} else {
				beforQuantity := stock.Quantity
				stock.Quantity += detailReq.Quantity
				stock.UpdatedAt = nowstr
				if err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", req.StorehouseUuid, detailReq.ProductUuid, detailReq.SkuUuid).Updates(stock).Error; err != nil {
					return err
				}

				// 创建库存记录
				stockopLog := &model.StorehouseProductOpLog{
					Uuid:                  uuid.New().String(),
					StorehouseProductUuid: stock.Uuid,
					StorehouseUuid:        inbound.StorehouseUuid,
					BeforeQuantity:        beforQuantity,
					Quantity:              stock.Quantity,
					OpQuantity:            detailReq.Quantity,
					OpType:                1,
					OpDesc:                "仓库入库,增加库存",
					OpBy:                  userId,
					CreatedAt:             nowstr,
				}
				if err := tx.Create(stockopLog).Error; err != nil {
					ctx.Logger.Error("Failed to create stockop log", err)
					return errors.New("failed to create stockop log")
				}
			}
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to create inbound", err)
		return errors.New("failed to create inbound")
	}

	return nil
}

func (s *StorehouseInboundService) GetInbound(ctx *app.Context, uuid string) (*model.StorehouseInbound, error) {
	inbound := &model.StorehouseInbound{}
	err := ctx.DB.Where("inbound_order_no = ?", uuid).First(inbound).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("inbound not found")
		}
		ctx.Logger.Error("Failed to get inbound by UUID", err)
		return nil, errors.New("failed to get inbound by UUID")
	}
	return inbound, nil
}

func (s *StorehouseInboundService) UpdateInbound(ctx *app.Context, req *model.StorehouseInboundUpdateReq) error {
	inbound := &model.StorehouseInbound{
		StorehouseUuid: req.StorehouseUuid,
		Title:          req.Title,
		InboundType:    req.InboundType,
		Status:         req.Status,
		InboundDate:    time.Now().Format("2006-01-02"),
		UpdatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("inbound_order_no = ?", req.InboundOrderNo).Updates(inbound).Error; err != nil {
			return err
		}

		if err := tx.Where("inbound_order_no = ?", req.InboundOrderNo).Delete(&model.StorehouseInboundDetail{}).Error; err != nil {
			return err
		}

		for _, detailReq := range req.Detail {
			detail := &model.StorehouseInboundDetail{
				InboundOrderNo: req.InboundOrderNo,
				ProductUuid:    detailReq.ProductUuid,
				SkuUuid:        detailReq.SkuUuid,
				Quantity:       detailReq.Quantity,
			}
			if err := tx.Create(detail).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to update inbound", err)
		return errors.New("failed to update inbound")
	}

	return nil
}

func (s *StorehouseInboundService) DeleteInbound(ctx *app.Context, uuid string) error {
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("inbound_order_no = ?", uuid).Delete(&model.StorehouseInbound{}).Error; err != nil {
			return err
		}
		if err := tx.Where("inbound_order_no = ?", uuid).Delete(&model.StorehouseInboundDetail{}).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to delete inbound", err)
		return errors.New("failed to delete inbound")
	}

	return nil
}

func (s *StorehouseInboundService) ListInbounds(ctx *app.Context, param *model.ReqStorehouseInboundQueryParam) (r *model.PagedResponse, err error) {
	var (
		inboundList []*model.StorehouseInbound
		total       int64
	)

	db := ctx.DB.Model(&model.StorehouseInbound{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&inboundList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	userUuids := make([]string, 0)
	for _, inbound := range inboundList {
		userUuids = append(userUuids, inbound.InboundBy)
	}

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
		return
	}

	storehouseUuids := make([]string, 0)
	for _, inbound := range inboundList {
		storehouseUuids = append(storehouseUuids, inbound.StorehouseUuid)
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}
	res := make([]*model.StorehouseInboundRes, 0)
	for _, inbound := range inboundList {
		inboundRes := &model.StorehouseInboundRes{
			StorehouseInbound: *inbound,
		}
		if storehouse, ok := storehouseMap[inbound.StorehouseUuid]; ok {
			inboundRes.Storehouse = *storehouse
		}

		if user, ok := userMap[inbound.InboundBy]; ok {
			inboundRes.InboundByUser = *user
		}

		res = append(res, inboundRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
