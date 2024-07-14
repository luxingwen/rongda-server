package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseInventoryCheckService struct {
}

func NewStorehouseInventoryCheckService() *StorehouseInventoryCheckService {
	return &StorehouseInventoryCheckService{}
}

func (s *StorehouseInventoryCheckService) CreateInventoryCheck(ctx *app.Context, userId string, req *model.StorehouseInventoryCheckReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	check := &model.StorehouseInventoryCheck{
		StorehouseUuid: req.StorehouseUuid,
		CheckOrderNo:   uuid.New().String(),
		CheckDate:      req.CheckDate,
		Status:         req.Status,
		CheckBy:        userId,
		CreatedAt:      nowStr,
		UpdatedAt:      nowStr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(check).Error; err != nil {
			ctx.Logger.Error("Failed to create inventory check", err)
			return errors.New("failed to create inventory check")
		}

		for _, detailReq := range req.Detail {

			// 获取库存
			stock := &model.StorehouseProduct{}
			err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", req.StorehouseUuid, detailReq.ProductUuid, detailReq.SkuUuid).First(stock).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.New("仓库中没有该商品")
				}
				ctx.Logger.Error("Failed to get stock", err)
				return errors.New("failed to get stock")
			}

			differenceQuantity := 0
			detailReq.DifferenceOp = "0" // 正常

			if detailReq.Quantity > stock.Quantity {
				detailReq.DifferenceOp = "1"
				differenceQuantity = detailReq.Quantity - stock.Quantity //盘盈
			}

			if detailReq.Quantity < stock.Quantity {
				detailReq.DifferenceOp = "2"
				differenceQuantity = stock.Quantity - detailReq.Quantity //盘亏
			}

			detail := &model.StorehouseInventoryCheckDetail{
				CheckOrderNo:       check.CheckOrderNo,
				ProductUuid:        detailReq.ProductUuid,
				SkuUuid:            detailReq.SkuUuid,
				Quantity:           detailReq.Quantity,
				DifferenceOp:       detailReq.DifferenceOp,
				CreatedAt:          nowStr,
				UpdatedAt:          nowStr,
				DifferenceQuantity: differenceQuantity,
			}

			if err := tx.Create(detail).Error; err != nil {
				ctx.Logger.Error("Failed to create inventory check detail", err)
				return errors.New("failed to create inventory check detail")
			}

			if detailReq.DifferenceOp != "0" {
				stock.Quantity = detailReq.Quantity
				stock.UpdatedAt = nowStr
				if err := tx.Where("uuid = ?", stock.Uuid).Updates(stock).Error; err != nil {
					ctx.Logger.Error("Failed to update stock", err)
					return err
				}
			}

		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *StorehouseInventoryCheckService) GetInventoryCheck(ctx *app.Context, requuid string) (*model.StorehouseInventoryCheck, error) {
	check := &model.StorehouseInventoryCheck{}
	err := ctx.DB.Where("uuid = ?", requuid).First(check).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("inventory check not found")
		}
		ctx.Logger.Error("Failed to get inventory check by ID", err)
		return nil, errors.New("failed to get inventory check by ID")
	}
	return check, nil
}

func (s *StorehouseInventoryCheckService) UpdateInventoryCheck(ctx *app.Context, check *model.StorehouseInventoryCheck) error {
	check.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(check).Error
	if err != nil {
		ctx.Logger.Error("Failed to update inventory check", err)
		return errors.New("failed to update inventory check")
	}
	return nil
}

func (s *StorehouseInventoryCheckService) DeleteInventoryCheck(ctx *app.Context, requuid string) error {
	err := ctx.DB.Where("uuid = ?", requuid).Delete(&model.StorehouseInventoryCheck{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete inventory check", err)
		return errors.New("failed to delete inventory check")
	}
	return nil
}

func (s *StorehouseInventoryCheckService) ListInventoryChecks(ctx *app.Context, param *model.ReqInventoryCheckQueryParam) (r *model.PagedResponse, err error) {
	var (
		checkList []*model.StorehouseInventoryCheck
		total     int64
	)

	db := ctx.DB.Model(&model.StorehouseInventoryCheck{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&checkList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	storehouseUuids := make([]string, 0)
	for _, check := range checkList {
		storehouseUuids = append(storehouseUuids, check.StorehouseUuid)
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseInventoryCheckRes, 0)
	for _, check := range checkList {
		checkitem := &model.StorehouseInventoryCheckRes{
			StorehouseInventoryCheck: *check,
		}
		if storehouse, ok := storehouseMap[check.StorehouseUuid]; ok {
			checkitem.Storehouse = *storehouse
		}
		res = append(res, checkitem)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
