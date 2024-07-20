package service

import (
	"encoding/json"
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseOrderService struct {
}

func NewPurchaseOrderService() *PurchaseOrderService {
	return &PurchaseOrderService{}
}

func (s *PurchaseOrderService) CreatePurchaseOrderFutures(ctx *app.Context, userId string, req *model.PurchaseOrderReq) error {

	attachment := ""
	if len(req.Attachment) > 0 {
		bdata, _ := json.Marshal(req.Attachment)
		attachment = string(bdata)
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:                 req.Title,
		OrderNo:               uuid.New().String(), // Generating a unique order number
		SupplierUuid:          req.SupplierUuid,
		CustomerUuid:          req.CustomerUuid,
		Date:                  req.Date,
		PIAgreementNo:         req.PIAgreementNo,
		OrderCurrency:         req.OrderCurrency,
		SettlementCurrency:    req.SettlementCurrency,
		Departure:             req.Departure,
		Destination:           req.Destination,
		EstimatedShippingDate: req.EstimatedShippingDate,
		EstimatedWarehouse:    req.EstimatedWarehouse,

		Purchaser:  userId,
		Status:     1, // Assuming 1 is the initial status
		CreatedAt:  nowStr,
		UpdatedAt:  nowStr,
		OrderType:  model.OrderTypeFutures,
		Attachment: attachment,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			ctx.Logger.Error("Failed to create purchase order", err)
			return errors.New("failed to create purchase order")
		}

		for _, detailReq := range req.Details {
			detail := &model.PurchaseOrderItem{
				PurchaseOrderNo: order.OrderNo,
				ProductUuid:     detailReq.ProductUuid,
				SkuUuid:         detailReq.SkuUuid,
				ProductName:     detailReq.ProductName,
				SkuName:         detailReq.SkuName,
				Quantity:        detailReq.Quantity,
				Price:           detailReq.Price,
				TotalAmount:     detailReq.TotalAmount,

				PIBoxNum:             detailReq.PIBoxNum,
				PIQuantity:           detailReq.PIQuantity,
				PIUnitPrice:          detailReq.PIUnitPrice,
				PITotalAmount:        detailReq.PITotalAmount,
				CabinetNo:            detailReq.CabinetNo,
				BillOfLadingNo:       detailReq.BillOfLadingNo,
				ShipName:             detailReq.ShipName,
				Voyage:               detailReq.Voyage,
				CIInvoiceNo:          detailReq.CIInvoiceNo,
				CIBoxNum:             detailReq.CIBoxNum,
				CIQuantity:           detailReq.CIQuantity,
				CIUnitPrice:          detailReq.CIUnitPrice,
				CITotalAmount:        detailReq.CITotalAmount,
				ProductionDate:       detailReq.ProductionDate,
				EstimatedArrivalDate: detailReq.EstimatedArrivalDate,
				Tariff:               detailReq.Tariff,
				VAT:                  detailReq.VAT,
				PaymentDate:          detailReq.PaymentDate,

				CreatedAt: nowStr,
				UpdatedAt: nowStr,
			}

			if err := tx.Create(detail).Error; err != nil {
				ctx.Logger.Error("Failed to create purchase order item", err)
				return errors.New("failed to create purchase order item")
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *PurchaseOrderService) CreatePurchaseOrderSpot(ctx *app.Context, userId string, req *model.PurchaseOrderReq) error {

	attachment := ""
	if len(req.Attachment) > 0 {
		bdata, _ := json.Marshal(req.Attachment)
		attachment = string(bdata)
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:                 req.Title,
		OrderNo:               uuid.New().String(), // Generating a unique order number
		SupplierUuid:          req.SupplierUuid,
		CustomerUuid:          req.CustomerUuid,
		Date:                  req.Date,
		PIAgreementNo:         req.PIAgreementNo,
		OrderCurrency:         req.OrderCurrency,
		SettlementCurrency:    req.SettlementCurrency,
		Departure:             req.Departure,
		Destination:           req.Destination,
		EstimatedShippingDate: req.EstimatedShippingDate,
		EstimatedWarehouse:    req.EstimatedWarehouse,
		ActualWarehouse:       req.ActualWarehouse,

		Purchaser:  userId,
		Status:     1, // Assuming 1 is the initial status
		CreatedAt:  nowStr,
		UpdatedAt:  nowStr,
		OrderType:  model.OrderTypeSpot,
		Attachment: attachment,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			ctx.Logger.Error("Failed to create purchase order", err)
			return errors.New("failed to create purchase order")
		}

		for _, detailReq := range req.Details {
			detail := &model.PurchaseOrderItem{
				PurchaseOrderNo: order.OrderNo,
				ProductUuid:     detailReq.ProductUuid,
				SkuUuid:         detailReq.SkuUuid,
				ProductName:     detailReq.ProductName,
				SkuName:         detailReq.SkuName,
				Quantity:        detailReq.Quantity,
				Price:           detailReq.Price,
				TotalAmount:     detailReq.TotalAmount,

				PIBoxNum:             detailReq.PIBoxNum,
				PIQuantity:           detailReq.PIQuantity,
				PIUnitPrice:          detailReq.PIUnitPrice,
				PITotalAmount:        detailReq.PITotalAmount,
				CabinetNo:            detailReq.CabinetNo,
				BillOfLadingNo:       detailReq.BillOfLadingNo,
				ShipName:             detailReq.ShipName,
				Voyage:               detailReq.Voyage,
				CIInvoiceNo:          detailReq.CIInvoiceNo,
				CIBoxNum:             detailReq.CIBoxNum,
				CIQuantity:           detailReq.CIQuantity,
				CIUnitPrice:          detailReq.CIUnitPrice,
				CITotalAmount:        detailReq.CITotalAmount,
				ProductionDate:       detailReq.ProductionDate,
				EstimatedArrivalDate: detailReq.EstimatedArrivalDate,
				Tariff:               detailReq.Tariff,
				VAT:                  detailReq.VAT,
				PaymentDate:          detailReq.PaymentDate,

				CreatedAt: nowStr,
				UpdatedAt: nowStr,
			}

			if err := tx.Create(detail).Error; err != nil {
				ctx.Logger.Error("Failed to create purchase order item", err)
				return errors.New("failed to create purchase order item")
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *PurchaseOrderService) GetPurchaseOrder(ctx *app.Context, orderNo string) (*model.PurchaseOrderRes, error) {
	order := &model.PurchaseOrderRes{}
	err := ctx.DB.Where("order_no = ?", orderNo).First(&order.PurchaseOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("purchase order not found")
		}
		ctx.Logger.Error("Failed to get purchase order by order no", err)
		return nil, errors.New("failed to get purchase order by order no")
	}

	// Get supplier information
	err = ctx.DB.Where("uuid = ?", order.SupplierUuid).First(&order.Supplier).Error
	if err != nil {
		ctx.Logger.Error("Failed to get supplier by uuid", err)
		return nil, errors.New("failed to get supplier by uuid")
	}

	user, err := NewUserService().GetUserByUUID(ctx, order.Purchaser)
	if err != nil && err.Error() != "user not found" {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, err
	}

	curuuids := make([]string, 0)
	curuuids = append(curuuids, order.OrderCurrency)
	curuuids = append(curuuids, order.SettlementCurrency)

	currencyMap, err := NewSettlementCurrencyService().GetSettlementCurrencyByUuids(ctx, curuuids)
	if err != nil {
		ctx.Logger.Error("Failed to get settlement currency by uuids", err)
		return nil, err
	}

	if currency, ok := currencyMap[order.OrderCurrency]; ok {
		order.OrderCurrencyInfo = currency
	}

	if currency, ok := currencyMap[order.SettlementCurrency]; ok {
		order.SettlementCurrencyInfo = currency
	}

	if order.EstimatedWarehouse != "" {
		storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, order.EstimatedWarehouse)
		if err != nil {
			ctx.Logger.Error("Failed to get storehouse by uuid", err)
		}
		order.EstimatedWarehouseInfo = storehouse
	}

	if order.ActualWarehouse != "" {
		storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, order.ActualWarehouse)
		if err != nil {
			ctx.Logger.Error("Failed to get storehouse by uuid", err)
		}
		order.ActualWarehouseInfo = storehouse
	}

	customer, err := NewCustomerService().GetCustomerByUUID(ctx, order.CustomerUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get customer by uuid", err)
	}
	order.CustomerInfo = customer

	if user != nil {

		order.PurchaserInfo = *user
	}
	return order, nil
}

// 获取采购单商品明细
func (s *PurchaseOrderService) GetPurchaseOrderItems(ctx *app.Context, orderNo string) ([]*model.PurchaseOrderItemRes, error) {
	var items []*model.PurchaseOrderItem
	err := ctx.DB.Where("purchase_order_no = ?", orderNo).Find(&items).Error
	if err != nil {
		ctx.Logger.Error("Failed to get purchase order items", err)
		return nil, errors.New("failed to get purchase order items")
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	for _, item := range items {
		productUuids = append(productUuids, item.ProductUuid)
		skuUuids = append(skuUuids, item.SkuUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return nil, err

	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return nil, err
	}

	res := make([]*model.PurchaseOrderItemRes, 0)
	for _, item := range items {
		purchaseOrderItem := &model.PurchaseOrderItemRes{
			PurchaseOrderItem: *item,
		}
		if product, ok := productMap[item.ProductUuid]; ok {
			purchaseOrderItem.Product = *product
		}
		if sku, ok := skuMap[item.SkuUuid]; ok {
			purchaseOrderItem.Sku = *sku
		}
		res = append(res, purchaseOrderItem)
	}

	return res, nil
}

func (s *PurchaseOrderService) UpdatePurchaseOrder(ctx *app.Context, order *model.PurchaseOrder) error {
	order.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(order).Error
	if err != nil {
		ctx.Logger.Error("Failed to update purchase order", err)
		return errors.New("failed to update purchase order")
	}
	return nil
}

func (s *PurchaseOrderService) DeletePurchaseOrder(ctx *app.Context, orderNo string) error {
	err := ctx.DB.Where("order_no = ?", orderNo).Delete(&model.PurchaseOrder{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete purchase order", err)
		return errors.New("failed to delete purchase order")
	}
	return nil
}

func (s *PurchaseOrderService) ListPurchaseOrders(ctx *app.Context, param *model.ReqPurchaseOrderQueryParam) (r *model.PagedResponse, err error) {
	var (
		orderList []*model.PurchaseOrder
		total     int64
	)

	db := ctx.DB.Model(&model.PurchaseOrder{})

	if param.SupplierUuid != "" {
		db = db.Where("supplier_uuid = ?", param.SupplierUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&orderList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	supplierUuids := make([]string, 0)
	customerUuids := make([]string, 0)
	for _, order := range orderList {
		supplierUuids = append(supplierUuids, order.SupplierUuid)
		customerUuids = append(customerUuids, order.CustomerUuid)
	}

	supplierMap, err := NewSupplierService().GetSupplierListByUUIDs(ctx, supplierUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get supplier list by UUIDs", err)
		return
	}

	customerMap, err := NewCustomerService().GetCustomerListByUUIDs(ctx, customerUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get customer list by UUIDs", err)
		return
	}

	res := make([]*model.PurchaseOrderRes, 0)
	for _, order := range orderList {
		purchaseOrderItem := &model.PurchaseOrderRes{
			PurchaseOrder: *order,
		}
		if supplier, ok := supplierMap[order.SupplierUuid]; ok {
			purchaseOrderItem.Supplier = *supplier
		}
		if customer, ok := customerMap[order.CustomerUuid]; ok {
			purchaseOrderItem.CustomerInfo = customer
		}
		res = append(res, purchaseOrderItem)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
