package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesOrderService struct {
}

func NewSalesOrderService() *SalesOrderService {
	return &SalesOrderService{}
}

func (s *SalesOrderService) CreateSalesOrder(ctx *app.Context, userId string, req *model.SalesOrderReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	orderNo := utils.GenerateOrderID()
	salesOrder := &model.SalesOrder{
		OrderNo:            orderNo,
		OrderType:          req.OrderType,
		Title:              req.Title,
		OrderDate:          req.OrderDate,
		DepositAmount:      float64(req.Deposit),
		OrderAmount:        float64(req.OrderAmount),
		Salesman:           userId,
		CustomerUuid:       req.CustomerUuid,
		SettlementCurrency: req.SettlementCurrency,
		DepositRatio:       req.DepositRatio,
		FinalPaymentAmount: req.FinalPaymentAmount,
		Remarks:            req.Remarks,
		OrderStatus:        model.OrderStatusPending,
		PurchaseOrderNo:    req.PurchaseOrderNo,
		EntrustOrderId:     req.EntrustOrderId,
		CreatedAt:          nowStr,
		UpdatedAt:          nowStr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(salesOrder).Error; err != nil {
			ctx.Logger.Error("Failed to create sales order", err)
			return errors.New("failed to create sales order")
		}

		if salesOrder.EntrustOrderId != "" {
			// 更新委托订单销售单号
			err := tx.Model(&model.EntrustOrder{}).Where("order_id = ?", salesOrder.EntrustOrderId).Updates(map[string]interface{}{
				"sales_order_no": salesOrder.OrderNo,
				"status":         "已处理",
				"updated_at":     nowStr,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update entrust order", err)
				return errors.New("failed to update entrust order")
			}
		}

		// 获取采购单信息
		var purchaseOrder model.PurchaseOrder
		if salesOrder.PurchaseOrderNo != "" {
			if err := tx.Where("order_no = ?", salesOrder.PurchaseOrderNo).First(&purchaseOrder).Error; err != nil {
				ctx.Logger.Error("Failed to get purchase order by order number", err)
				return errors.New("failed to get purchase order by order number")
			}
		}

		for _, itemReq := range req.ProductList {
			item := &model.SalesOrderItem{
				Uuid:                   uuid.New().String(),
				OrderNo:                orderNo,
				ProductUuid:            itemReq.ProductUuid,
				SkuUuid:                itemReq.SkuUuid,
				PurchaseOrderProductNo: itemReq.PurchaseOrderProductNo,
				ProductQuantity:        itemReq.ProductQuantity,
				ProductPrice:           itemReq.ProductPrice,
				ProductAmount:          itemReq.ProductAmount,
				BoxNum:                 itemReq.BoxNum,
				CreatedAt:              nowStr,
				UpdatedAt:              nowStr,
			}

			var purchaseOrderProduct model.PurchaseOrderItem
			if item.PurchaseOrderProductNo != "" {
				if err := tx.Where("purchase_order_product_no = ?", item.PurchaseOrderProductNo).First(&purchaseOrderProduct).Error; err != nil {
					ctx.Logger.Error("Failed to get purchase order product by purchase order product no", err)
					return errors.New("failed to get purchase order product by purchase order product no")
				}
				productQuantity := purchaseOrderProduct.Quantity
				productBoxNum := purchaseOrderProduct.BoxNum
				if purchaseOrder.OrderType == model.OrderTypeFutures {
					productQuantity = purchaseOrderProduct.CIQuantity
					productBoxNum = purchaseOrderProduct.CIBoxNum
				}
				if productQuantity < item.ProductQuantity {
					return errors.New("采购单商品数量不足，无法创建销售订单")
				}
				if productBoxNum < item.BoxNum {
					return errors.New("采购单箱数不足，无法创建销售订单")
				}
			}

			if err := tx.Create(item).Error; err != nil {
				ctx.Logger.Error("Failed to create sales order item", err)
				return errors.New("failed to create sales order item")
			}
		}

		// 创建步骤链

		stepChain := model.StepChain{
			Uuid:        uuid.New().String(),
			RefId:       orderNo,
			RefType:     model.StepChainRefTypeSalesOrder,
			ChainName:   "销售订单",
			ChainStatus: model.StepChainStatusWait,
			ChainType:   model.StepChainRefTypeSalesOrder,
			CreatedAt:   nowStr,
			UpdatedAt:   nowStr,
		}

		if err := tx.Create(&stepChain).Error; err != nil {
			ctx.Logger.Error("Failed to create step chain", err)
			return errors.New("failed to create step chain")
		}

		steps := make([]model.Step, 0)
		for _, itemStep := range model.SalesSteps {
			itemStep.Uuid = uuid.New().String()
			itemStep.ChainId = stepChain.Uuid
			itemStep.CreatedAt = nowStr
			itemStep.UpdatedAt = nowStr
			itemStep.Status = model.StepStatusWait
			if itemStep.Title == "创建订单" {
				itemStep.Status = model.StepStatusFinish
				itemStep.RefId = orderNo
				itemStep.RefType = model.StepRefTypeSalesOrder
				itemStep.StepType = model.StepTypeDetail
			}
			if itemStep.Title == "订单确认" {
				itemStep.Status = model.StepStatusProcess
			}
			steps = append(steps, itemStep)
		}

		if err := tx.Create(&steps).Error; err != nil {
			ctx.Logger.Error("Failed to create steps", err)
			return errors.New("failed to create steps")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *SalesOrderService) GetSalesOrder(ctx *app.Context, orderNo string) (*model.SalesOrderRes, error) {
	salesOrder := &model.SalesOrder{}
	err := ctx.DB.Where("order_no = ?", orderNo).First(salesOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("sales order not found")
		}
		ctx.Logger.Error("Failed to get sales order by order number", err)
		return nil, errors.New("failed to get sales order by order number")
	}

	customerInfo, err := NewCustomerService().GetCustomerByUUID(ctx, salesOrder.CustomerUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get customer info", err)
		return nil, err
	}

	user, err := NewUserService().GetUserByUUID(ctx, salesOrder.Salesman)
	if err != nil && err.Error() != "user not found" {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, err
	}

	salesOrderRes := &model.SalesOrderRes{
		SalesOrder:   *salesOrder,
		CustomerInfo: customerInfo,
		SalesmanInfo: user,
	}

	return salesOrderRes, nil
}

// 获取订单商品
func (s *SalesOrderService) GetSalesOrderItems(ctx *app.Context, orderNo string) (r []*model.SalesOrderItemRes, err error) {
	var (
		orderItems []*model.SalesOrderItem
	)

	db := ctx.DB.Model(&model.SalesOrderItem{})
	if err = db.Where("order_no = ?", orderNo).Find(&orderItems).Error; err != nil {
		return
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)

	for _, item := range orderItems {
		productUuids = append(productUuids, item.ProductUuid)
		skuUuids = append(skuUuids, item.SkuUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return
	}

	res := make([]*model.SalesOrderItemRes, 0)
	for _, item := range orderItems {
		itemRes := &model.SalesOrderItemRes{
			SalesOrderItem: *item,
		}
		if product, ok := productMap[item.ProductUuid]; ok {
			itemRes.ProductInfo = product
		}
		if sku, ok := skuMap[item.SkuUuid]; ok {
			itemRes.SkuInfo = sku
		}
		res = append(res, itemRes)
	}

	return res, nil

}

func (s *SalesOrderService) UpdateSalesOrder(ctx *app.Context, userId string, req *model.SalesOrderReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	orderNo := req.OrderNo
	salesOrder := &model.SalesOrder{
		OrderNo:            orderNo,
		OrderType:          req.OrderType,
		Title:              req.Title,
		OrderDate:          req.OrderDate,
		DepositAmount:      float64(req.Deposit),
		OrderAmount:        float64(req.OrderAmount),
		Updater:            userId,
		CustomerUuid:       req.CustomerUuid,
		SettlementCurrency: req.SettlementCurrency,
		DepositRatio:       req.DepositRatio,
		FinalPaymentAmount: req.FinalPaymentAmount,
		Remarks:            req.Remarks,
		OrderStatus:        model.OrderStatusPending,
		PurchaseOrderNo:    req.PurchaseOrderNo,
		CreatedAt:          nowStr,
		UpdatedAt:          nowStr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("order_no = ?", orderNo).Updates(salesOrder).Error; err != nil {
			ctx.Logger.Error("Failed to update sales order", err)
			return errors.New("failed to update sales order")
		}

		if err := tx.Where("order_no = ?", orderNo).Delete(&model.SalesOrderItem{}).Error; err != nil {
			ctx.Logger.Error("Failed to delete sales order items", err)
			return errors.New("failed to delete sales order items")
		}

		for _, itemReq := range req.ProductList {
			item := &model.SalesOrderItem{
				Uuid:                   uuid.New().String(),
				OrderNo:                orderNo,
				ProductUuid:            itemReq.ProductUuid,
				PurchaseOrderProductNo: itemReq.PurchaseOrderProductNo,
				SkuUuid:                itemReq.SkuUuid,
				ProductQuantity:        itemReq.ProductQuantity,
				ProductPrice:           itemReq.ProductPrice,
				ProductAmount:          itemReq.ProductAmount,
				BoxNum:                 itemReq.BoxNum,
				CreatedAt:              nowStr,
				UpdatedAt:              nowStr,
			}
			if err := tx.Create(item).Error; err != nil {
				ctx.Logger.Error("Failed to create sales order item", err)
				return errors.New("failed to create sales order item")
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// 更新订单状态
func (s *SalesOrderService) UpdateSalesOrderStatus(ctx *app.Context, orderNo string, status string) error {
	err := ctx.DB.Model(&model.SalesOrder{}).Where("order_no = ?", orderNo).Update("order_status", status).Error
	if err != nil {
		ctx.Logger.Error("Failed to update sales order status", err)
		return errors.New("failed to update sales order status")
	}
	return nil
}

func (s *SalesOrderService) DeleteSalesOrder(ctx *app.Context, orderNo string) error {
	err := ctx.DB.Where("order_no = ?", orderNo).Delete(&model.SalesOrder{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete sales order", err)
		return errors.New("failed to delete sales order")
	}
	return nil
}

func (s *SalesOrderService) ListSalesOrders(ctx *app.Context, param *model.ReqSalesOrderQueryParam) (r *model.PagedResponse, err error) {
	var (
		orderList []*model.SalesOrder
		total     int64
	)

	db := ctx.DB.Model(&model.SalesOrder{})

	if param.CustomerUuid != "" {
		db = db.Where("customer_uuid = ?", param.CustomerUuid)
	}

	if err = db.Order("id DESC").Offset(param.GetOffset()).Limit(param.PageSize).Find(&orderList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	customerUuids := make([]string, 0)
	userUuids := make([]string, 0)

	for _, order := range orderList {
		customerUuids = append(customerUuids, order.CustomerUuid)
		userUuids = append(userUuids, order.Salesman)
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

	res := make([]*model.SalesOrderRes, 0)
	for _, order := range orderList {
		orderRes := &model.SalesOrderRes{
			SalesOrder: *order,
		}
		if customer, ok := customerMap[order.CustomerUuid]; ok {
			orderRes.CustomerInfo = customer
		}
		if user, ok := userMap[order.Salesman]; ok {
			orderRes.SalesmanInfo = user
		}
		res = append(res, orderRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}

// 获取所有可用订单
func (s *SalesOrderService) ListAllSalesOrders(ctx *app.Context) (r []*model.SalesOrder, err error) {
	var (
		orderList []*model.SalesOrder
	)

	db := ctx.DB.Model(&model.SalesOrder{})

	if err = db.Where("order_status = ?", model.OrderStatusPendingShipped).Find(&orderList).Error; err != nil {
		return
	}

	return orderList, nil
}

// 确认订单
func (s *SalesOrderService) ConfirmSalesOrder(ctx *app.Context, params *model.ReqSalesOrderConfirmParam) error {

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		// 更新订单状态
		err := tx.Model(&model.SalesOrder{}).Where("order_no IN ?", params.OrderNoList).Updates(map[string]interface{}{
			"order_status": model.OrderStatusConfirmed,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		}).Error
		if err != nil {
			ctx.Logger.Error("Failed to update sales order status", err)
			tx.Rollback()
			return errors.New("failed to update sales order status")
		}

		for _, orderNo := range params.OrderNoList {

			// 获取步骤链信息
			stepChain := &model.StepChain{}
			err = tx.Where("ref_id = ?", orderNo).First(stepChain).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step chain", err)
				tx.Rollback()
				return errors.New("failed to get step chain")
			}

			if params.Op == "confirm" {
				// 更新步骤状态
				err = tx.Model(&model.Step{}).Where("chain_id = ? AND title = ?", stepChain.Uuid, "订单确认").Updates(map[string]interface{}{
					"status":     model.StepStatusFinish,
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				}).Error
				if err != nil {
					ctx.Logger.Error("Failed to update step status", err)
					tx.Rollback()
					return errors.New("failed to update step status")
				}

				err = tx.Model(&model.Step{}).Where("chain_id = ? AND title = ?", stepChain.Uuid, "创建合同").Updates(map[string]interface{}{
					"status":     model.StepStatusProcess,
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				}).Error
				if err != nil {
					ctx.Logger.Error("Failed to update step status", err)
					tx.Rollback()
					return errors.New("failed to update step status")
				}
			}

			if params.Op == "cancel" {
				// 更新步骤状态
				err = tx.Model(&model.Step{}).Where("chain_id = ? AND title = ?", stepChain.Uuid, "订单确认").Updates(map[string]interface{}{
					"status":     model.StepStatusError,
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				}).Error
				if err != nil {
					ctx.Logger.Error("Failed to update step status", err)
					tx.Rollback()
					return errors.New("failed to update step status")
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
