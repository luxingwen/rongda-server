package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

	"gorm.io/gorm"
)

type EntrustOrderService struct{}

func NewEntrustOrderService() *EntrustOrderService {
	return &EntrustOrderService{}
}

// CreateEntrustOrder 创建新的委托订单
func (s *EntrustOrderService) CreateEntrustOrder(ctx *app.Context, userUuid string, order *model.EntrustOrder) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	order.CreatedAt = now
	order.UpdatedAt = now
	order.OrderId = utils.GenerateOrderID()
	order.UserUuid = userUuid

	err := ctx.DB.Create(order).Error
	if err != nil {
		ctx.Logger.Error("Failed to create entrust order", err)
		return errors.New("failed to create entrust order")
	}
	return nil
}

// GetEntrustOrderByUUID 根据UUID获取委托订单
func (s *EntrustOrderService) GetEntrustOrderByUUID(ctx *app.Context, uuid string) (*model.EntrustOrder, error) {
	order := &model.EntrustOrder{}
	err := ctx.DB.Where("order_id = ?", uuid).First(order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("entrust order not found")
		}
		ctx.Logger.Error("Failed to get entrust order by UUID", err)
		return nil, errors.New("failed to get entrust order by UUID")
	}
	return order, nil
}

// UpdateEntrustOrder 更新委托订单
func (s *EntrustOrderService) UpdateEntrustOrder(ctx *app.Context, order *model.EntrustOrder) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	order.UpdatedAt = now
	err := ctx.DB.Where("order_id = ?", order.OrderId).Updates(order).Error
	if err != nil {
		ctx.Logger.Error("Failed to update entrust order", err)
		return errors.New("failed to update entrust order")
	}
	return nil
}

// DeleteEntrustOrder 删除委托订单
func (s *EntrustOrderService) DeleteEntrustOrder(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.EntrustOrder{}).Where("order_id = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete entrust order", err)
		return errors.New("failed to delete entrust order")
	}
	return nil
}

// GetEntrustOrderList 获取委托订单列表
func (s *EntrustOrderService) GetEntrustOrderList(ctx *app.Context, params *model.ReqEntrustOrderQueryParam) (*model.PagedResponse, error) {
	var (
		orders []*model.EntrustOrder
		total  int64
	)

	db := ctx.DB.Model(&model.EntrustOrder{})

	if params.UserUuid != "" {
		db = db.Where("user_uuid = ?", params.UserUuid)
	}

	if params.TeamUuid != "" {
		db = db.Where("team_uuid = ?", params.TeamUuid)
	}

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	if params.StartDate != "" {
		db = db.Where("created_at >= ?", params.StartDate)
	}

	if params.EndDate != "" {
		db = db.Where("created_at <= ?", params.EndDate)
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get entrust order count", err)
		return nil, errors.New("failed to get entrust order count")
	}

	err = db.Order("id DESC").Offset(params.GetOffset()).Limit(params.PageSize).Find(&orders).Error
	if err != nil {
		ctx.Logger.Error("Failed to get entrust order list", err)
		return nil, errors.New("failed to get entrust order list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  orders,
	}, nil
}
