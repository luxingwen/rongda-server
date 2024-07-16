package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

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
	orderNo := uuid.New().String()
	salesOrder := &model.SalesOrder{
		OrderNo:       orderNo,
		OrderType:     req.OrderType,
		OrderDate:     req.OrderDate,
		DepositAmount: float64(req.Deposit),
		OrderAmount:   float64(req.OrderAmount),
		Salesman:      userId,
		CustomerUuid:  req.CustomerUuid,
		TaxAmount:     req.TaxAmount,
		Remarks:       req.Remarks,
		OrderStatus:   "待支付",
		CreatedAt:     nowStr,
		UpdatedAt:     nowStr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(salesOrder).Error; err != nil {
			ctx.Logger.Error("Failed to create sales order", err)
			return errors.New("failed to create sales order")
		}

		for _, itemReq := range req.ProductList {
			item := &model.SalesOrderItem{
				Uuid:            uuid.New().String(),
				OrderNo:         orderNo,
				ProductUuid:     itemReq.ProductUuid,
				SkuUuid:         itemReq.SkuUuid,
				ProductQuantity: float64(itemReq.ProductQuantity),
				ProductPrice:    float64(itemReq.ProductPrice),
				ProductAmount:   float64(itemReq.ProductAmount),
				CreatedAt:       nowStr,
				UpdatedAt:       nowStr,
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

func (s *SalesOrderService) GetSalesOrder(ctx *app.Context, orderNo string) (*model.SalesOrder, error) {
	salesOrder := &model.SalesOrder{}
	err := ctx.DB.Where("order_no = ?", orderNo).First(salesOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("sales order not found")
		}
		ctx.Logger.Error("Failed to get sales order by order number", err)
		return nil, errors.New("failed to get sales order by order number")
	}
	return salesOrder, nil
}

func (s *SalesOrderService) UpdateSalesOrder(ctx *app.Context, salesOrder *model.SalesOrder) error {
	salesOrder.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(salesOrder).Error
	if err != nil {
		ctx.Logger.Error("Failed to update sales order", err)
		return errors.New("failed to update sales order")
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

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&orderList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     orderList,
	}
	return
}

// 获取所有可用订单
func (s *SalesOrderService) ListAllSalesOrders(ctx *app.Context) (r []*model.SalesOrder, err error) {
	var (
		orderList []*model.SalesOrder
	)

	db := ctx.DB.Model(&model.SalesOrder{})

	if err = db.Find(&orderList).Error; err != nil {
		return
	}

	return orderList, nil
}
