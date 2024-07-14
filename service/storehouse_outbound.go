package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseOutboundService struct {
}

func NewStorehouseOutboundService() *StorehouseOutboundService {
	return &StorehouseOutboundService{}
}

func (s *StorehouseOutboundService) CreateOutbound(ctx *app.Context, userId string, req *model.StorehouseOutboundReq) error {

	nowstr := time.Now().Format("2006-01-02 15:04:05")
	outbound := &model.StorehouseOutbound{
		StorehouseUuid:  req.StorehouseUuid,
		OutboundType:    req.OutboundType,
		Status:          req.Status,
		OutboundOrderNo: uuid.New().String(), // Generating a unique order number
		OutboundDate:    time.Now().Format("2006-01-02"),
		OutboundBy:      userId, // Assuming the user ID is available in the context
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:       time.Now().Format("2006-01-02 15:04:05"),
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(outbound).Error; err != nil {
			ctx.Logger.Error("Failed to create outbound", err)
			return errors.New("failed to create outbound")
		}

		for _, detailReq := range req.Detail {
			detail := &model.StorehouseOutboundDetail{
				OutboundOrderNo: outbound.OutboundOrderNo,
				ProductUuid:     detailReq.ProductUuid,
				SkuUuid:         detailReq.SkuUuid,
				Quantity:        detailReq.Quantity,
				CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt:       time.Now().Format("2006-01-02 15:04:05"),
			}

			// 获取库存

			// 先获取库存
			stock := &model.StorehouseProduct{}
			err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", req.StorehouseUuid, detailReq.ProductUuid, detailReq.SkuUuid).First(stock).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.New("仓库中没有该商品")
				}
				ctx.Logger.Error("Failed to get stock", err)
				return err
			}

			// 出库数量大于库存数量
			if stock.Quantity < detailReq.Quantity {
				return errors.New("库存不足")
			}

			if err := tx.Create(detail).Error; err != nil {
				ctx.Logger.Error("Failed to create outbound detail", err)
				return errors.New("failed to create outbound detail")
			}

			stock.Quantity -= detailReq.Quantity
			stock.UpdatedAt = nowstr
			if err := tx.Where("uuid = ?", stock.Uuid).Updates(stock).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *StorehouseOutboundService) GetOutbound(ctx *app.Context, requuid string) (*model.StorehouseOutbound, error) {
	outbound := &model.StorehouseOutbound{}
	err := ctx.DB.Where("uuid = ?", requuid).First(outbound).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("outbound not found")
		}
		ctx.Logger.Error("Failed to get outbound by ID", err)
		return nil, errors.New("failed to get outbound by ID")
	}
	return outbound, nil
}

func (s *StorehouseOutboundService) UpdateOutbound(ctx *app.Context, outbound *model.StorehouseOutbound) error {
	outbound.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(outbound).Error
	if err != nil {
		ctx.Logger.Error("Failed to update outbound", err)
		return errors.New("failed to update outbound")
	}
	return nil
}

func (s *StorehouseOutboundService) DeleteOutbound(ctx *app.Context, requuid string) error {
	err := ctx.DB.Where("uuid = ?", requuid).Delete(&model.StorehouseOutbound{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete outbound", err)
		return errors.New("failed to delete outbound")
	}
	return nil
}

func (s *StorehouseOutboundService) ListOutbounds(ctx *app.Context, param *model.ReqStorehouseOutboundQueryParam) (r *model.PagedResponse, err error) {
	var (
		outboundList []*model.StorehouseOutbound
		total        int64
	)

	db := ctx.DB.Model(&model.StorehouseOutbound{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&outboundList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	storehouseUuids := make([]string, 0)
	for _, outbound := range outboundList {
		storehouseUuids = append(storehouseUuids, outbound.StorehouseUuid)
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseOutboundRes, 0)
	for _, outbound := range outboundList {
		outboundRes := &model.StorehouseOutboundRes{
			StorehouseOutbound: *outbound,
		}
		if storehouse, ok := storehouseMap[outbound.StorehouseUuid]; ok {
			outboundRes.Storehouse = *storehouse
		}
		res = append(res, outboundRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
