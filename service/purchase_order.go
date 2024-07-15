package service

import (
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

func (s *PurchaseOrderService) CreatePurchaseOrder(ctx *app.Context, userId string, req *model.PurchaseOrderReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.PurchaseOrder{
		Title:        req.Title,
		OrderNo:      uuid.New().String(), // Generating a unique order number
		SupplierUuid: req.SupplierUuid,
		Date:         req.Date,
		Deposit:      req.Deposit,
		Tax:          req.Tax,
		TotalAmount:  req.TotalAmount,
		Purchaser:    userId,
		Status:       1, // Assuming 1 is the initial status
		CreatedAt:    nowStr,
		UpdatedAt:    nowStr,
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
				CreatedAt:       nowStr,
				UpdatedAt:       nowStr,
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

	return order, nil
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
	for _, order := range orderList {
		supplierUuids = append(supplierUuids, order.SupplierUuid)
	}

	supplierMap, err := NewSupplierService().GetSupplierListByUUIDs(ctx, supplierUuids)
	if err != nil {
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
