package service

import (
	"errors"
	"sgin/model"
	"sgin/pkg/app"
	"time"

	"github.com/google/uuid"
)

type OrderFileService struct {
}

func NewOrderFileService() *OrderFileService {
	return &OrderFileService{}
}

func (s *OrderFileService) CreateOrderFile(ctx *app.Context, orderFile *model.OrderFile) error {
	orderFile.Uuid = uuid.New().String()
	orderFile.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	orderFile.UpdatedAt = orderFile.CreatedAt

	err := ctx.DB.Create(orderFile).Error
	if err != nil {
		ctx.Logger.Error("Failed to create orderFile", err)
		return errors.New("failed to create orderFile")
	}
	return nil
}

func (s *OrderFileService) GetOrderFileListByOrderNo(ctx *app.Context, orderNo string) ([]*model.OrderFile, error) {
	orderFiles := make([]*model.OrderFile, 0)
	err := ctx.DB.Where("order_no = ?", orderNo).Find(&orderFiles).Error
	if err != nil {
		ctx.Logger.Error("Failed to get orderFile list by orderNo", err)
		return nil, errors.New("failed to get orderFile list by orderNo")
	}
	return orderFiles, nil
}

// DeleteOrderFile 删除订单文件
func (s *OrderFileService) DeleteOrderFile(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.OrderFile{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete orderFile", err)
		return errors.New("failed to delete orderFile")
	}
	return nil
}

// GetOrderFileByUuid
func (s *OrderFileService) GetOrderFileByUuid(ctx *app.Context, uuid string) (*model.OrderFile, error) {
	orderFile := new(model.OrderFile)
	err := ctx.DB.Where("uuid = ?", uuid).First(orderFile).Error
	if err != nil {
		ctx.Logger.Error("Failed to get orderFile by uuid", err)
		return nil, errors.New("failed to get orderFile by uuid")
	}
	return orderFile, nil
}
