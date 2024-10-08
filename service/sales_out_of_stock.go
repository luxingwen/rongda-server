package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesOutOfStockService struct {
}

func NewSalesOutOfStockService() *SalesOutOfStockService {
	return &SalesOutOfStockService{}
}

func (s *SalesOutOfStockService) CreateSalesOutOfStock(ctx *app.Context, userId string, req *model.SalesOutOfStockReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	outOfStock := &model.SalesOutOfStock{
		UUID:           uuid.New().String(),
		OutOfStockDate: req.OutOfStockDate,
		SalesOrderNo:   req.SalesOrderNo,
		CustomerUuid:   req.CustomerUuid,
		BatchNo:        req.BatchNo,
		Registrant:     req.Registrant,
		StorehouseUuid: req.StorehouseUuid,
		Remark:         req.Remark,
		Status:         req.Status,
		CreatedAt:      nowStr,
		UpdatedAt:      nowStr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(outOfStock).Error; err != nil {
			ctx.Logger.Error("Failed to create sales out of stock", err)
			return errors.New("failed to create sales out of stock")
		}

		for _, itemReq := range req.Items {
			item := &model.SalesOutOfStockItem{
				SalesOutOfStockNo: outOfStock.UUID,
				ProductUuid:       itemReq.ProductUuid,
				SkuUuid:           itemReq.SkuUuid,
				Quantity:          itemReq.Quantity,
				Price:             itemReq.Price,
				TotalAmount:       itemReq.TotalAmount,
				CreatedAt:         nowStr,
				UpdatedAt:         nowStr,
			}

			if err := tx.Create(item).Error; err != nil {
				ctx.Logger.Error("Failed to create sales out of stock item", err)
				return errors.New("failed to create sales out of stock item")
			}

			stock := &model.StorehouseProduct{}
			err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", req.StorehouseUuid, itemReq.ProductUuid, itemReq.SkuUuid).First(stock).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.New("仓库中没有该商品")
				}
				ctx.Logger.Error("Failed to get stock", err)
				return err
			}

			if stock.Quantity < itemReq.Quantity {
				return errors.New("库存不足")
			}

			stock.Quantity -= itemReq.Quantity
			stock.UpdatedAt = nowStr
			if err := tx.Save(stock).Error; err != nil {
				ctx.Logger.Error("Failed to update stock", err)
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

func (s *SalesOutOfStockService) GetSalesOutOfStock(ctx *app.Context, uuid string) (*model.SalesOutOfStockRes, error) {
	outOfStock := &model.SalesOutOfStock{}
	err := ctx.DB.Where("uuid = ?", uuid).First(outOfStock).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("sales out of stock not found")
		}
		ctx.Logger.Error("Failed to get sales out of stock by ID", err)
		return nil, errors.New("failed to get sales out of stock by ID")
	}

	user, err := NewUserService().GetUserByUUID(ctx, outOfStock.Registrant)

	if err != nil && err.Error() != "user not found" {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, err
	}

	customer, err := NewCustomerService().GetCustomerByUUID(ctx, outOfStock.CustomerUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get customer by UUID", err)
	}

	storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, outOfStock.StorehouseUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
	}

	outOfStockRes := &model.SalesOutOfStockRes{
		SalesOutOfStock: *outOfStock,
		RegistrantInfo:  user,
		CustomerInfo:    customer,
		StorehouseInfo:  storehouse,
	}

	return outOfStockRes, nil
}

// 获取出库单商品明细
func (s *SalesOutOfStockService) GetSalesOutOfStockItems(ctx *app.Context, outOfStockNo string) ([]*model.SalesOutOfStockItemRes, error) {
	var items []*model.SalesOutOfStockItem
	err := ctx.DB.Where("sales_out_of_stock_no = ?", outOfStockNo).Find(&items).Error
	if err != nil {
		ctx.Logger.Error("Failed to get sales out of stock items", err)
		return nil, errors.New("failed to get sales out of stock items")
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	for _, item := range items {
		productUuids = append(productUuids, item.ProductUuid)
		skuUuids = append(skuUuids, item.SkuUuid)
	}

	products, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get products by UUIDs", err)
		return nil, err
	}

	skus, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get skus by UUIDs", err)
		return nil, err
	}

	res := make([]*model.SalesOutOfStockItemRes, 0)

	for _, item := range items {
		outOfStockItem := &model.SalesOutOfStockItemRes{
			SalesOutOfStockItem: *item,
		}
		if product, ok := products[item.ProductUuid]; ok {
			outOfStockItem.ProductInfo = product
		}
		if sku, ok := skus[item.SkuUuid]; ok {
			outOfStockItem.SkuInfo = sku
		}
		res = append(res, outOfStockItem)
	}

	return res, nil
}

func (s *SalesOutOfStockService) UpdateSalesOutOfStock(ctx *app.Context, outOfStock *model.SalesOutOfStock) error {
	outOfStock.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(outOfStock).Error
	if err != nil {
		ctx.Logger.Error("Failed to update sales out of stock", err)
		return errors.New("failed to update sales out of stock")
	}
	return nil
}

func (s *SalesOutOfStockService) DeleteSalesOutOfStock(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.SalesOutOfStock{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete sales out of stock", err)
		return errors.New("failed to delete sales out of stock")
	}
	return nil
}

func (s *SalesOutOfStockService) ListSalesOutOfStocks(ctx *app.Context, param *model.ReqSalesOutOfStockQueryParam) (r *model.PagedResponse, err error) {
	var (
		outOfStockList []*model.SalesOutOfStock
		total          int64
	)

	db := ctx.DB.Model(&model.SalesOutOfStock{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&outOfStockList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     outOfStockList,
	}
	return
}
