package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseService struct {
}

func NewStorehouseService() *StorehouseService {
	return &StorehouseService{}
}

func (s *StorehouseService) CreateStorehouse(ctx *app.Context, storehouse *model.Storehouse) error {
	storehouse.Uuid = uuid.New().String()
	storehouse.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	storehouse.UpdatedAt = storehouse.CreatedAt

	err := ctx.DB.Create(storehouse).Error
	if err != nil {
		ctx.Logger.Error("Failed to create storehouse", err)
		return errors.New("failed to create storehouse")
	}
	return nil
}

func (s *StorehouseService) GetStorehouseByUUID(ctx *app.Context, uuid string) (*model.Storehouse, error) {
	storehouse := &model.Storehouse{}
	err := ctx.DB.Where("uuid = ?", uuid).First(storehouse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("storehouse not found")
		}
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, errors.New("failed to get storehouse by UUID")
	}
	return storehouse, nil
}

func (s *StorehouseService) UpdateStorehouse(ctx *app.Context, storehouse *model.Storehouse) error {
	storehouse.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", storehouse.Uuid).Updates(storehouse).Error
	if err != nil {
		ctx.Logger.Error("Failed to update storehouse", err)
		return errors.New("failed to update storehouse")
	}
	return nil
}

func (s *StorehouseService) DeleteStorehouse(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Storehouse{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete storehouse", err)
		return errors.New("failed to delete storehouse")
	}
	return nil
}

// 获取仓库列表
func (s *StorehouseService) GetStorehouseList(ctx *app.Context, param *model.ReqStorehouseQueryParam) (r *model.PagedResponse, err error) {
	var (
		storehouseList []*model.Storehouse
		total          int64
	)

	db := ctx.DB.Model(&model.Storehouse{})

	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&storehouseList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     storehouseList,
	}
	return
}

// 获取所有可用仓库
func (s *StorehouseService) GetAvailableStorehouses(ctx *app.Context) (storehouseList []*model.Storehouse, err error) {
	err = ctx.DB.Where("status = ?", model.StorehouseStatusEnabled).Find(&storehouseList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available storehouses", err)
		return nil, errors.New("failed to get available storehouses")
	}
	return
}

// 根据uuids获取仓库列表
func (s *StorehouseService) GetStorehousesByUUIDs(ctx *app.Context, uuids []string) (r map[string]*model.Storehouse, err error) {
	storehouseList := make([]*model.Storehouse, 0)
	r = make(map[string]*model.Storehouse)
	err = ctx.DB.Where("uuid in (?)", uuids).Find(&storehouseList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get storehouses by UUIDs", err)
		return nil, errors.New("failed to get storehouses by UUIDs")
	}
	for _, storehouse := range storehouseList {
		r[storehouse.Uuid] = storehouse
	}
	return
}

// CreateStorehouseOutboundOrder
func (s *StorehouseService) CreateStorehouseOutboundOrder(ctx *app.Context, params *model.ReqStorehouseOutboundOrder) error {

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		// 先获取库存商品列表信息
		storehouseGoodsList := make([]*model.StorehouseProduct, 0)

		err := tx.Where("uuid IN ?", params.StorehouseProductUuids).Find(&storehouseGoodsList).Error
		if err != nil {
			tx.Rollback()
			ctx.Logger.Error("Failed to get storehouse products by UUIDs", err)
			return errors.New("failed to get storehouse products by UUIDs")
		}

		// 判断是否是整柜出库
		if params.IsFullCabinet {
			// 遍历storehouseGoodsList
			mPurchaseStorehouse := make(map[string]map[string]string, 0)

			for _, storehouseGoods := range storehouseGoodsList {

				if mdata, ok := mPurchaseStorehouse[storehouseGoods.PurchaseOrderNo]; ok {

					if _, ok := mdata[storehouseGoods.CabinetNo]; !ok {
						mdata[storehouseGoods.CabinetNo] = storehouseGoods.CabinetNo
						mPurchaseStorehouse[storehouseGoods.PurchaseOrderNo] = mdata
					}
					continue
				}

				mdata := make(map[string]string, 0)
				mdata[storehouseGoods.CabinetNo] = storehouseGoods.CabinetNo
				mPurchaseStorehouse[storehouseGoods.PurchaseOrderNo] = mdata

			}

			storehouseGoodsList := make([]*model.StorehouseProduct, 0)

			// 遍历mPurchaseStorehouse
			for purchaseOrderNo, mdata := range mPurchaseStorehouse {

				storeProductListItems := make([]*model.StorehouseProduct, 0)

				cabinetNos := make([]string, 0)
				for _, cabinetNo := range mdata {
					cabinetNos = append(cabinetNos, cabinetNo)
				}

				err := tx.Where("purchase_order_no = ? AND cabinet_no IN ?", purchaseOrderNo, cabinetNos).Find(&storeProductListItems).Error
				if err != nil {
					ctx.Logger.Error("Failed to get storehouse products by UUIDs", err)
					return errors.New("failed to get storehouse products by UUIDs")
				}

				storehouseGoodsList = append(storehouseGoodsList, storeProductListItems...)
			}

		}

		// 出库操作
		for _, storehouseGoods := range storehouseGoodsList {

			// 添加出库记录
			storehouseOutboundRecord := &model.StorehouseOutboundOrderDetail{
				OutboundOrderNo:       uuid.New().String(),
				StorehouseProductUuid: storehouseGoods.Uuid,
				StorehouseUuid:        storehouseGoods.StorehouseUuid,
				ProductUuid:           storehouseGoods.ProductUuid,
				SkuUuid:               storehouseGoods.SkuUuid,
				Quantity:              storehouseGoods.Quantity,
				BoxNum:                storehouseGoods.BoxNum,
				PurchaseOrderNo:       storehouseGoods.PurchaseOrderNo,
				CabinetNo:             storehouseGoods.CabinetNo,
				Status:                "申请中",
				ApplyBy:               params.ApplyBy,
				CustomerUuid:          storehouseGoods.CustomerUuid,
				CreatedAt:             time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt:             time.Now().Format("2006-01-02 15:04:05"),
			}

			err = tx.Create(storehouseOutboundRecord).Error
			if err != nil {
				tx.Rollback()
				ctx.Logger.Error("Failed to create storehouse outbound record", err)
				return errors.New("failed to create storehouse outbound record")
			}

			// 创建出库记录

			beforeQuantity := storehouseGoods.Quantity
			beforBoxNum := storehouseGoods.BoxNum

			storehouseGoods.Quantity = 0
			storehouseGoods.BoxNum = 0

			err := tx.Model(&model.StorehouseProduct{}).Where("uuid = ?", storehouseGoods.Uuid).Updates(map[string]interface{}{
				"quantity":   storehouseGoods.Quantity,
				"box_num":    storehouseGoods.BoxNum,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			}).Error

			// 添加出库log
			storehouseOutboundLog := &model.StorehouseProductOpLog{
				OutboundOrderNo:       storehouseOutboundRecord.OutboundOrderNo,
				PurchaseOrderNo:       storehouseGoods.PurchaseOrderNo,
				StorehouseProductUuid: storehouseGoods.Uuid,
				StorehouseUuid:        storehouseGoods.StorehouseUuid,
				ProductUuid:           storehouseGoods.ProductUuid,
				SkuUuid:               storehouseGoods.SkuUuid,
				BeforeQuantity:        beforeQuantity,
				BeforeBoxNum:          beforBoxNum,
				Quantity:              storehouseGoods.Quantity,
				BoxNum:                storehouseGoods.BoxNum,
				TeamUuid:              storehouseGoods.CustomerUuid,
				OpQuantity:            beforeQuantity,
				OpBoxNum:              beforBoxNum,
				CreatedAt:             time.Now().Format("2006-01-02 15:04:05"),
				OpType:                model.StorehouseProductOpLogOpTypeOutbound,
				OpBy:                  params.ApplyBy,
				CabinetNo:             storehouseGoods.CabinetNo,
			}

			err = tx.Create(storehouseOutboundLog).Error
			if err != nil {
				ctx.Logger.Error("Failed to create storehouse outbound log", err)
				tx.Rollback()
				return errors.New("failed to create storehouse outbound log")
			}

		}

		return nil
	})

	if err != nil {

		ctx.Logger.Error("Failed to create storehouse outbound order", err)
		return errors.New("failed to create storehouse outbound order")
	}

	return nil

}

// GetStorehouseOutboundOrderList
func (s *StorehouseService) GetStorehouseOutboundOrderList(ctx *app.Context, param *model.ReqStorehouseOutboundOrderQueryParam) (r *model.PagedResponse, err error) {
	var (
		storehouseOutboundOrderList []*model.StorehouseOutboundOrderDetail
		total                       int64
	)

	db := ctx.DB.Model(&model.StorehouseOutboundOrderDetail{})

	// if param.OutboundOrderNo != "" {
	// 	db = db.Where("outbound_order_no like ?", "%"+param.OutboundOrderNo+"%")
	// }

	// if param.PurchaseOrderNo != "" {
	// 	db = db.Where("purchase_order_no like ?", "%"+param.PurchaseOrderNo+"%")
	// }

	if param.TeamUuid != "" {
		db = db.Where("customer_uuid = ?", param.TeamUuid)
	}

	if param.Status != "" {
		db = db.Where("status = ?", param.Status)
	}

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&storehouseOutboundOrderList).Error; err != nil {
		return
	}

	purchaseOrderNos := make([]string, 0)
	productUuids := make([]string, 0)
	storehouseUuids := make([]string, 0)
	skuuuids := make([]string, 0)
	customerUuids := make([]string, 0)
	for _, storehouseOutboundOrder := range storehouseOutboundOrderList {
		purchaseOrderNos = append(purchaseOrderNos, storehouseOutboundOrder.PurchaseOrderNo)
		productUuids = append(productUuids, storehouseOutboundOrder.ProductUuid)
		skuuuids = append(skuuuids, storehouseOutboundOrder.SkuUuid)
		storehouseUuids = append(storehouseUuids, storehouseOutboundOrder.StorehouseUuid)
		customerUuids = append(customerUuids, storehouseOutboundOrder.CustomerUuid)
	}

	purchaseOrderMap, err := NewPurchaseOrderService().GetPurchaseOrderListByOrderNos(ctx, purchaseOrderNos)

	if err != nil {
		ctx.Logger.Error("Failed to get purchase order list by order nos", err)
		return nil, errors.New("failed to get purchase order list by order nos")
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return nil, errors.New("failed to get product list by UUIDs")
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuuuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return nil, errors.New("failed to get sku list by UUIDs")
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return nil, errors.New("failed to get storehouse list by UUIDs")
	}

	customerMap, err := NewCustomerService().GetCustomerListByUUIDs(ctx, customerUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get customer list by UUIDs", err)
		return nil, errors.New("failed to get customer list by UUIDs")
	}

	res := make([]*model.StorehouseOutboundOrderDetailRes, 0)
	for _, storehouseOutboundOrder := range storehouseOutboundOrderList {
		storehouseOutboundOrderRes := &model.StorehouseOutboundOrderDetailRes{
			StorehouseOutboundOrderDetail: *storehouseOutboundOrder,
		}

		if purchaseOrder, ok := purchaseOrderMap[storehouseOutboundOrder.PurchaseOrderNo]; ok {
			storehouseOutboundOrderRes.PurchaseOrderInfo = *purchaseOrder
		}

		if product, ok := productMap[storehouseOutboundOrder.ProductUuid]; ok {
			storehouseOutboundOrderRes.Product = *product
		}

		if sku, ok := skuMap[storehouseOutboundOrder.SkuUuid]; ok {
			storehouseOutboundOrderRes.Sku = *sku
		}

		if storehouse, ok := storehouseMap[storehouseOutboundOrder.StorehouseUuid]; ok {
			storehouseOutboundOrderRes.StorehouseInfo = *storehouse
		}

		if customer, ok := customerMap[storehouseOutboundOrder.CustomerUuid]; ok {
			storehouseOutboundOrderRes.CustomerInfo = *customer
		}

		res = append(res, storehouseOutboundOrderRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}

// GetStorehouseOutboundOrderInfo
func (s *StorehouseService) GetStorehouseOutboundOrderInfo(ctx *app.Context, uuid string) (*model.StorehouseOutboundOrderDetailRes, error) {

	storehouseOutboundOrder := &model.StorehouseOutboundOrderDetail{}
	err := ctx.DB.Where("outbound_order_no = ?", uuid).First(storehouseOutboundOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("storehouse outbound order not found")
		}
		ctx.Logger.Error("Failed to get storehouse outbound order by UUID", err)
		return nil, errors.New("failed to get storehouse outbound order by UUID")
	}

	purchaseOrder, err := NewPurchaseOrderService().GetPurchaseOrderRecord(ctx, storehouseOutboundOrder.PurchaseOrderNo)
	if err != nil {
		ctx.Logger.Error("Failed to get purchase order by order no", err)
		return nil, errors.New("failed to get purchase order by order no")
	}

	product, err := NewProductService().GetProductByUUID(ctx, storehouseOutboundOrder.ProductUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get product by UUID", err)
		return nil, errors.New("failed to get product by UUID")
	}

	sku, err := NewSkuService().GetSkuByUUID(ctx, storehouseOutboundOrder.SkuUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get sku by UUID", err)
		return nil, errors.New("failed to get sku by UUID")
	}

	storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, storehouseOutboundOrder.StorehouseUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, errors.New("failed to get storehouse by UUID")
	}

	customer, err := NewCustomerService().GetCustomerByUUID(ctx, storehouseOutboundOrder.CustomerUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get customer by UUID", err)
		return nil, errors.New("failed to get customer by UUID")
	}

	storehouseOutboundOrderRes := &model.StorehouseOutboundOrderDetailRes{
		StorehouseOutboundOrderDetail: *storehouseOutboundOrder,
		PurchaseOrderInfo:             *purchaseOrder,
		Product:                       *product,
		Sku:                           *sku,
		StorehouseInfo:                *storehouse,
		CustomerInfo:                  *customer,
	}

	return storehouseOutboundOrderRes, nil
}

// DeleteStorehouseOutboundOrder
func (s *StorehouseService) DeleteStorehouseOutboundOrder(ctx *app.Context, uuid string) error {
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		storehouseOutboundOrder := &model.StorehouseOutboundOrderDetail{}
		err := tx.Where("outbound_order_no = ?", uuid).First(storehouseOutboundOrder).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("storehouse outbound order not found")
			}
			ctx.Logger.Error("Failed to get storehouse outbound order by UUID", err)
			return errors.New("failed to get storehouse outbound order by UUID")
		}

		// 删除出库记录
		err = tx.Where("outbound_order_no = ?", uuid).Delete(&model.StorehouseOutboundOrderDetail{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete storehouse outbound order", err)
			return errors.New("failed to delete storehouse outbound order")
		}

		// 删除出库log
		err = tx.Where("outbound_order_no = ?", uuid).Delete(&model.StorehouseProductOpLog{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete storehouse outbound log", err)
			return errors.New("failed to delete storehouse outbound log")
		}

		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to delete storehouse outbound order", err)
		return errors.New("failed to delete storehouse outbound order")
	}

	return nil
}

// UpdateStorehouseOutboundOrderDetailStatus
func (s *StorehouseService) UpdateStorehouseOutboundOrderDetailStatus(ctx *app.Context, param *model.ReqUpdateStatus) error {

	err := ctx.DB.Model(&model.StorehouseOutboundOrderDetail{}).Where("outbound_order_no = ?", param.Uuid).Updates(map[string]interface{}{
		"status":     param.Status,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}).Error

	if err != nil {
		ctx.Logger.Error("Failed to update storehouse outbound order detail status", err)
		return errors.New("failed to update storehouse outbound order detail status")
	}

	return nil
}
