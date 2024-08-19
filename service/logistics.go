package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogisticsService struct {
}

func NewLogisticsService() *LogisticsService {
	return &LogisticsService{}
}

// CreateLogistics creates a new logistics record in the database
func (s *LogisticsService) CreateLogistics(ctx *app.Context, logistics *model.Logistics) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	logistics.CreatedAt = now
	logistics.UpdatedAt = now
	logistics.Uuid = uuid.New().String()

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(logistics).Error; err != nil {
			ctx.Logger.Error("Failed to create logistics", err)
			return errors.New("failed to create logistics")
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to create logistics", err)
		return errors.New("failed to create logistics")
	}
	return nil
}

// GetLogisticsByUUID retrieves a logistics record by its UUID
func (s *LogisticsService) GetLogisticsByUUID(ctx *app.Context, uuid string) (*model.Logistics, error) {
	logistics := &model.Logistics{}
	if err := ctx.DB.Where("uuid = ?", uuid).First(logistics).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("logistics not found")
		}
		ctx.Logger.Error("Failed to get logistics by UUID", err)
		return nil, errors.New("failed to get logistics by UUID")
	}
	return logistics, nil
}

// UpdateLogistics updates an existing logistics record
func (s *LogisticsService) UpdateLogistics(ctx *app.Context, logistics *model.Logistics) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	logistics.UpdatedAt = now

	if err := ctx.DB.Where("uuid = ?", logistics.Uuid).Updates(logistics).Error; err != nil {
		ctx.Logger.Error("Failed to update logistics", err)
		return errors.New("failed to update logistics")
	}

	return nil
}

// DeleteLogistics marks a logistics record as deleted
func (s *LogisticsService) DeleteLogistics(ctx *app.Context, uuid string) error {
	if err := ctx.DB.Model(&model.Logistics{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error; err != nil {
		ctx.Logger.Error("Failed to delete logistics", err)
		return errors.New("failed to delete logistics")
	}

	return nil
}

// GetLogisticsList retrieves a list of logistics records based on query parameters
func (s *LogisticsService) GetLogisticsList(ctx *app.Context, params *model.ReqLogisticsQueryParam) (*model.PagedResponse, error) {
	var (
		logisticsList []*model.Logistics
		total         int64
	)

	db := ctx.DB.Model(&model.Logistics{})

	if params.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+params.OrderNo+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	if err := db.Count(&total).Error; err != nil {
		ctx.Logger.Error("Failed to get logistics count", err)
		return nil, errors.New("failed to get logistics count")
	}

	if err := db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&logisticsList).Error; err != nil {
		ctx.Logger.Error("Failed to get logistics list", err)
		return nil, errors.New("failed to get logistics list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  logisticsList,
	}, nil
}
