package service

import (
	"encoding/json"
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

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

	if req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:                 req.Title,
		OrderNo:               utils.GenerateOrderID(), // Generating a unique order number
		EntrustOrderId:        req.EntrustOrderId,
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
		DepositAmount:         req.DepositAmount,
		DepositRatio:          req.DepositRatio,

		Purchaser:  userId,
		Status:     model.PurchaseOrderStatusPending, // Assuming 1 is the initial status
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
				PurchaseOrderNo:        order.OrderNo,
				PurchaseOrderProductNo: uuid.New().String(),
				ProductUuid:            detailReq.ProductUuid,
				SkuUuid:                detailReq.SkuUuid,
				ProductName:            detailReq.ProductName,
				SkuName:                detailReq.SkuName,
				Quantity:               detailReq.Quantity,
				Price:                  detailReq.Price,
				TotalAmount:            detailReq.TotalAmount,

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
				CIResidualAmount:     detailReq.CIResidualAmount,
				ProductionDate:       detailReq.ProductionDate,
				EstimatedArrivalDate: detailReq.EstimatedArrivalDate,
				RMBDepositAmount:     detailReq.RMBDepositAmount,
				RMBResidualAmount:    detailReq.RMBResidualAmount,
				DepositExchangeRate:  detailReq.DepositExchangeRate,
				ResidualExchangeRate: detailReq.ResidualExchangeRate,
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

	if req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:                 req.Title,
		OrderNo:               utils.GenerateOrderID(), // Generating a unique order number
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
		DepositAmount:         req.DepositAmount,
		DepositRatio:          req.DepositRatio,

		Purchaser:  userId,
		Status:     model.PurchaseOrderStatusPending, // Assuming 1 is the initial status
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
				PurchaseOrderNo:        order.OrderNo,
				PurchaseOrderProductNo: uuid.New().String(),
				ProductUuid:            detailReq.ProductUuid,
				SkuUuid:                detailReq.SkuUuid,
				ProductName:            detailReq.ProductName,
				SkuName:                detailReq.SkuName,
				Quantity:               detailReq.Quantity,
				BoxNum:                 detailReq.BoxNum,
				Price:                  detailReq.Price,
				TotalAmount:            detailReq.TotalAmount,
				PIBoxNum:               detailReq.PIBoxNum,
				PIQuantity:             detailReq.PIQuantity,
				PIUnitPrice:            detailReq.PIUnitPrice,
				PITotalAmount:          detailReq.PITotalAmount,
				CabinetNo:              detailReq.CabinetNo,
				BillOfLadingNo:         detailReq.BillOfLadingNo,
				ShipName:               detailReq.ShipName,
				Voyage:                 detailReq.Voyage,
				CIInvoiceNo:            detailReq.CIInvoiceNo,
				CIBoxNum:               detailReq.CIBoxNum,
				CIQuantity:             detailReq.CIQuantity,
				CIUnitPrice:            detailReq.CIUnitPrice,
				CITotalAmount:          detailReq.CITotalAmount,
				ProductionDate:         detailReq.ProductionDate,
				EstimatedArrivalDate:   detailReq.EstimatedArrivalDate,
				Tariff:                 detailReq.Tariff,
				VAT:                    detailReq.VAT,
				PaymentDate:            detailReq.PaymentDate,
				Desc:                   detailReq.Desc,

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

func (s *PurchaseOrderService) GetPurchaseOrderRecord(ctx *app.Context, orderNo string) (*model.PurchaseOrder, error) {
	order := &model.PurchaseOrder{}
	err := ctx.DB.Where("order_no = ?", orderNo).First(order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("purchase order not found")
		}
		ctx.Logger.Error("Failed to get purchase order by order no", err)
		return nil, errors.New("failed to get purchase order by order no")
	}
	return order, nil
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
	err := ctx.DB.Where("order_no = ?", order.OrderNo).Updates(order).Error
	if err != nil {
		ctx.Logger.Error("Failed to update purchase order", err)
		return errors.New("failed to update purchase order")
	}
	return nil
}

func (s *PurchaseOrderService) UpdatePurchaseOrderFutures(ctx *app.Context, userId string, req *model.PurchaseOrderReq) error {

	attachment := ""
	if len(req.Attachment) > 0 {
		bdata, _ := json.Marshal(req.Attachment)
		attachment = string(bdata)
	}

	if req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:                 req.Title,
		OrderNo:               req.OrderNo,
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
		DepositAmount:         req.DepositAmount,
		DepositRatio:          req.DepositRatio,

		Updater:    userId,
		Status:     req.Status, // Assuming 1 is the initial status
		UpdatedAt:  nowStr,
		OrderType:  model.OrderTypeFutures,
		Attachment: attachment,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("order_no = ?", order.OrderNo).Updates(order).Error; err != nil {
			ctx.Logger.Error("Failed to update purchase order", err)
			return errors.New("failed to update purchase order")
		}

		// 先获取原来的订单明细
		oldItems := make([]*model.PurchaseOrderItem, 0)
		if err := tx.Where("purchase_order_no = ?", order.OrderNo).Find(&oldItems).Error; err != nil {
			ctx.Logger.Error("Failed to get purchase order items", err)
			return errors.New("failed to get purchase order items")
		}

		oldItemsMap := make(map[string]*model.PurchaseOrderItem)
		for _, item := range oldItems {
			oldItemsMap[item.PurchaseOrderProductNo] = item
		}

		for _, detailReq := range req.Details {
			itemPurchaseOrderProductNo := detailReq.PurchaseOrderProductNo
			// 如果有原来的订单明细，就更新原来的订单明细
			if itemPurchaseOrderProductNo == "" {
				itemPurchaseOrderProductNo = uuid.New().String()
			}

			detail := &model.PurchaseOrderItem{
				PurchaseOrderNo:        order.OrderNo,
				PurchaseOrderProductNo: itemPurchaseOrderProductNo,
				ProductUuid:            detailReq.ProductUuid,
				SkuUuid:                detailReq.SkuUuid,
				ProductName:            detailReq.ProductName,
				SkuName:                detailReq.SkuName,
				Quantity:               detailReq.Quantity,
				Price:                  detailReq.Price,
				TotalAmount:            detailReq.TotalAmount,

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
				CIResidualAmount:     detailReq.CIResidualAmount,
				ProductionDate:       detailReq.ProductionDate,
				EstimatedArrivalDate: detailReq.EstimatedArrivalDate,
				RMBDepositAmount:     detailReq.RMBDepositAmount,
				RMBResidualAmount:    detailReq.RMBResidualAmount,
				DepositExchangeRate:  detailReq.DepositExchangeRate,
				ResidualExchangeRate: detailReq.ResidualExchangeRate,
				Tariff:               detailReq.Tariff,
				VAT:                  detailReq.VAT,
				PaymentDate:          detailReq.PaymentDate,

				CreatedAt: nowStr,
				UpdatedAt: nowStr,
			}

			if detailReq.PurchaseOrderProductNo != "" {
				delete(oldItemsMap, itemPurchaseOrderProductNo)

				colsitem := []string{"product_uuid", "sku_uuid", "product_name", "sku_name", "quantity", "box_num", "price",
					"total_amount", "pi_box_num", "pi_quantity", "pi_unit_price", "pi_total_amount", "cabinet_no", "bill_of_lading_no",
					"ship_name", "voyage", "ci_invoice_no", "ci_box_num", "ci_quantity", "ci_unit_price", "ci_total_amount", "ci_residual_amount",
					"production_date", "estimated_arrival_date", "rmb_deposit_amount", "rmb_residual_amount", "deposit_exchange_rate", "residual_exchange_rate",
					"tariff", "vat", "payment_date", "desc", "updated_at"}

				if err := tx.Where("purchase_order_product_no = ?", detailReq.PurchaseOrderProductNo).Select(colsitem).Updates(detail).Error; err != nil {
					ctx.Logger.Error("Failed to update purchase order item", err)
					return errors.New("failed to update purchase order item")
				}

				// 获取销售订单商品
				var salesOrderItem model.SalesOrderItem
				if err := tx.Where("purchase_order_product_no = ?", detailReq.PurchaseOrderProductNo).First(&salesOrderItem).Error; err != nil && err != gorm.ErrRecordNotFound {
					ctx.Logger.Error("Failed to get sales order item by purchase order product no", err)
					return errors.New("failed to get sales order item by purchase order product no")
				}

				flagUpdate := false
				if salesOrderItem.ProductQuantity > detailReq.CIQuantity {
					salesOrderItem.ProductQuantity = detailReq.CIQuantity
					flagUpdate = true
				}

				if salesOrderItem.BoxNum > detailReq.CIBoxNum {
					salesOrderItem.BoxNum = detailReq.CIBoxNum
					flagUpdate = true
				}
				if flagUpdate {
					if err := tx.Where("purchase_order_product_no = ?", detailReq.PurchaseOrderProductNo).Select("product_quantity", "box_num").Updates(salesOrderItem).Error; err != nil {
						ctx.Logger.Error("Failed to update sales order item", err)
						return errors.New("failed to update sales order item")
					}
				}

			} else {

				if err := tx.Create(detail).Error; err != nil {
					ctx.Logger.Error("Failed to create purchase order item", err)
					return errors.New("failed to create purchase order item")
				}
			}

		}

		// 删除原来的订单明细
		deleteOldPurchaseOrderItem := []string{}
		for _, item := range oldItemsMap {
			deleteOldPurchaseOrderItem = append(deleteOldPurchaseOrderItem, item.PurchaseOrderProductNo)
		}

		if len(deleteOldPurchaseOrderItem) > 0 {
			if err := tx.Where("purchase_order_product_no IN ?", deleteOldPurchaseOrderItem).Delete(&model.PurchaseOrderItem{}).Error; err != nil {
				ctx.Logger.Error("Failed to delete purchase order item", err)
				return errors.New("failed to delete purchase order item")
			}

			// 删除销售订单商品
			if err := tx.Where("purchase_order_product_no IN ?", deleteOldPurchaseOrderItem).Delete(&model.SalesOrderItem{}).Error; err != nil {
				ctx.Logger.Error("Failed to delete sales order item", err)
				return errors.New("failed to delete sales order item")
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *PurchaseOrderService) UpdatePurchaseOrderSpot(ctx *app.Context, userId string, req *model.PurchaseOrderReq) error {

	attachment := ""
	if len(req.Attachment) > 0 {
		bdata, _ := json.Marshal(req.Attachment)
		attachment = string(bdata)
	}

	if req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:                 req.Title,
		OrderNo:               req.OrderNo, // Generating a unique order number
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
		DepositAmount:         req.DepositAmount,
		DepositRatio:          req.DepositRatio,

		Updater:    userId,
		Status:     req.Status, // Assuming 1 is the initial status
		UpdatedAt:  nowStr,
		OrderType:  model.OrderTypeSpot,
		Attachment: attachment,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		cols := []string{"title", "supplier_uuid", "customer_uuid", "date", "pi_agreement_no", "order_currency", "settlement_currency", "departure", "destination", "estimated_shipping_date", "estimated_warehouse", "actual_warehouse", "deposit_amount", "deposit_ratio", "updater", "updated_at", "order_type", "attachment"}
		if err := tx.Where("order_no = ?", order.OrderNo).Select(cols).Updates(order).Error; err != nil {
			ctx.Logger.Error("Failed to update purchase order", err)
			return errors.New("failed to update purchase order")
		}

		// 先获取原来的订单明细
		oldItems := make([]*model.PurchaseOrderItem, 0)
		if err := tx.Where("purchase_order_no = ?", order.OrderNo).Find(&oldItems).Error; err != nil {
			ctx.Logger.Error("Failed to get purchase order items", err)
			return errors.New("failed to get purchase order items")
		}

		oldItemsMap := make(map[string]*model.PurchaseOrderItem)
		for _, item := range oldItems {
			oldItemsMap[item.PurchaseOrderProductNo] = item
		}

		for _, detailReq := range req.Details {
			itemPurchaseOrderProductNo := detailReq.PurchaseOrderProductNo
			// 如果有原来的订单明细，就更新原来的订单明细
			if itemPurchaseOrderProductNo == "" {
				itemPurchaseOrderProductNo = uuid.New().String()
			}

			detail := &model.PurchaseOrderItem{
				PurchaseOrderNo:        order.OrderNo,
				PurchaseOrderProductNo: itemPurchaseOrderProductNo,
				ProductUuid:            detailReq.ProductUuid,
				SkuUuid:                detailReq.SkuUuid,
				ProductName:            detailReq.ProductName,
				SkuName:                detailReq.SkuName,
				Quantity:               detailReq.Quantity,
				BoxNum:                 detailReq.BoxNum,
				Price:                  detailReq.Price,
				TotalAmount:            detailReq.TotalAmount,
				PIBoxNum:               detailReq.PIBoxNum,
				PIQuantity:             detailReq.PIQuantity,
				PIUnitPrice:            detailReq.PIUnitPrice,
				PITotalAmount:          detailReq.PITotalAmount,
				CabinetNo:              detailReq.CabinetNo,
				BillOfLadingNo:         detailReq.BillOfLadingNo,
				ShipName:               detailReq.ShipName,
				Voyage:                 detailReq.Voyage,
				CIInvoiceNo:            detailReq.CIInvoiceNo,
				CIBoxNum:               detailReq.CIBoxNum,
				CIQuantity:             detailReq.CIQuantity,
				CIUnitPrice:            detailReq.CIUnitPrice,
				CITotalAmount:          detailReq.CITotalAmount,
				ProductionDate:         detailReq.ProductionDate,
				EstimatedArrivalDate:   detailReq.EstimatedArrivalDate,
				Tariff:                 detailReq.Tariff,
				VAT:                    detailReq.VAT,
				PaymentDate:            detailReq.PaymentDate,
				Desc:                   detailReq.Desc,

				CreatedAt: nowStr,
				UpdatedAt: nowStr,
			}

			if detailReq.PurchaseOrderProductNo != "" {
				delete(oldItemsMap, itemPurchaseOrderProductNo)

				colsitem := []string{"product_uuid", "sku_uuid", "product_name", "sku_name", "quantity", "box_num", "price",
					"total_amount", "pi_box_num", "pi_quantity", "pi_unit_price", "pi_total_amount", "cabinet_no", "bill_of_lading_no",
					"ship_name", "voyage", "ci_invoice_no", "ci_box_num", "ci_quantity", "ci_unit_price", "ci_total_amount", "production_date", "estimated_arrival_date", "tariff", "vat", "payment_date", "desc", "updated_at"}

				if err := tx.Where("purchase_order_product_no = ?", detailReq.PurchaseOrderProductNo).Select(colsitem).Updates(detail).Error; err != nil {
					ctx.Logger.Error("Failed to update purchase order item", err)
					return errors.New("failed to update purchase order item")
				}

				// 获取销售订单商品
				var salesOrderItem model.SalesOrderItem
				if err := tx.Where("purchase_order_product_no = ?", detailReq.PurchaseOrderProductNo).First(&salesOrderItem).Error; err != nil && err != gorm.ErrRecordNotFound {
					ctx.Logger.Error("Failed to get sales order item by purchase order product no", err)
					return errors.New("failed to get sales order item by purchase order product no")
				}

				flagUpdate := false
				if salesOrderItem.ProductQuantity > detailReq.Quantity {
					salesOrderItem.ProductQuantity = detailReq.Quantity
					flagUpdate = true
				}

				if salesOrderItem.BoxNum > detailReq.BoxNum {
					salesOrderItem.BoxNum = detailReq.BoxNum
					flagUpdate = true
				}
				if flagUpdate {
					if err := tx.Where("purchase_order_product_no = ?", detailReq.PurchaseOrderProductNo).Select("product_quantity", "box_num").Updates(salesOrderItem).Error; err != nil {
						ctx.Logger.Error("Failed to update sales order item", err)
						return errors.New("failed to update sales order item")
					}
				}

			} else {
				if err := tx.Create(detail).Error; err != nil {
					ctx.Logger.Error("Failed to create purchase order item", err)
					return errors.New("failed to create purchase order item")
				}
			}

		}

		// 删除原来的订单明细
		deleteOldPurchaseOrderItem := []string{}
		for _, item := range oldItemsMap {
			deleteOldPurchaseOrderItem = append(deleteOldPurchaseOrderItem, item.PurchaseOrderProductNo)
		}

		if len(deleteOldPurchaseOrderItem) > 0 {
			if err := tx.Where("purchase_order_product_no IN ?", deleteOldPurchaseOrderItem).Delete(&model.PurchaseOrderItem{}).Error; err != nil {
				ctx.Logger.Error("Failed to delete purchase order item", err)
				return errors.New("failed to delete purchase order item")
			}

			// 删除销售订单商品
			if err := tx.Where("purchase_order_product_no IN ?", deleteOldPurchaseOrderItem).Delete(&model.SalesOrderItem{}).Error; err != nil {
				ctx.Logger.Error("Failed to delete sales order item", err)
				return errors.New("failed to delete sales order item")
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *PurchaseOrderService) UpdatePurchaseOrderStatus(ctx *app.Context, orderNo string, status string) error {
	err := ctx.DB.Model(&model.PurchaseOrder{}).Where("order_no = ?", orderNo).Update("status", status).Error
	if err != nil {
		ctx.Logger.Error("Failed to update purchase order status", err)
		return errors.New("failed to update purchase order status")
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

	if param.CustomerUuid != "" {
		db = db.Where("customer_uuid = ?", param.CustomerUuid)
	}

	if param.Title != "" {
		db = db.Where("title LIKE ?", "%"+param.Title+"%")
	}

	if param.OrderNo != "" {
		db = db.Where("order_no = ?", param.OrderNo)
	}

	if err = db.Order("id DESC").Offset(param.GetOffset()).Limit(param.PageSize).Find(&orderList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	supplierUuids := make([]string, 0)
	customerUuids := make([]string, 0)
	userUuids := make([]string, 0)
	for _, order := range orderList {
		supplierUuids = append(supplierUuids, order.SupplierUuid)
		customerUuids = append(customerUuids, order.CustomerUuid)
		userUuids = append(userUuids, order.Purchaser)
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

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
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
		if user, ok := userMap[order.Purchaser]; ok {
			purchaseOrderItem.PurchaserInfo = *user
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

// 获取所有可用的采购订单
func (s *PurchaseOrderService) GetAvailablePurchaseOrderList(ctx *app.Context) (r []*model.PurchaseOrder, err error) {
	err = ctx.DB.Where("status = ?", model.PurchaseOrderStatusDone).Find(&r).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available purchase order list", err)
		return nil, errors.New("failed to get available purchase order list")
	}
	return
}

// 根据订单状态获取采购订单列表
func (s *PurchaseOrderService) GetPurchaseOrderListByStatus(ctx *app.Context, status []string) (r []*model.PurchaseOrder, err error) {
	err = ctx.DB.Where("status IN ?", status).Find(&r).Error
	if err != nil {
		ctx.Logger.Error("Failed to get purchase order list by status", err)
		return nil, errors.New("failed to get purchase order list by status")
	}
	return
}

// 补全model.PurchaseOrderItemReq 信息
func (s *PurchaseOrderService) CompletePurchaseOrderItem(ctx *app.Context, item []model.PurchaseOrderItemReq) (r []model.PurchaseOrderItemReq, err error) {
	produckNames := make([]string, 0)
	skuCodes := make([]string, 0)
	skuSpecs := make([]string, 0)
	for _, v := range item {
		produckNames = append(produckNames, v.ProductName)
		skuCodes = append(skuCodes, v.SkuCode)
		skuSpecs = append(skuSpecs, v.SkuSpec)
	}

	productlist, err := NewProductService().GetProductListByNames(ctx, produckNames)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by names", err)
		return
	}

	skulist, err := NewSkuService().GetSkuListByCodesSpecs(ctx, skuCodes, skuSpecs)

	if err != nil {
		ctx.Logger.Error("Failed to get sku list by names", err)
		return
	}

	findproduct := func(item model.PurchaseOrderItemReq) (*model.Product, error) {
		for _, v := range productlist {
			if v.Name == item.ProductName {
				return v, nil
			}
		}
		return nil, errors.New("product not found")
	}

	findsku := func(productUuid string, item model.PurchaseOrderItemReq) (*model.Sku, error) {
		for _, v := range skulist {
			if v.ProductUuid == productUuid && v.Code == item.SkuCode && v.Specification == item.SkuSpec {
				return v, nil
			}
		}
		return nil, errors.New("sku not found")
	}

	for _, v := range item {
		product, err := findproduct(v)
		if err != nil {
			ctx.Logger.Error("Failed to find product", err)
			return nil, err
		}
		sku, err := findsku(product.Uuid, v)
		if err != nil {
			ctx.Logger.Error("Failed to find sku", err)
			return nil, err
		}
		v.ProductUuid = product.Uuid
		v.SkuUuid = sku.UUID
		r = append(r, v)
	}
	return
}
