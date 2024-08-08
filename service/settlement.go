package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SettlementService struct {
}

func NewSettlementService() *SettlementService {
	return &SettlementService{}
}

func (s *SettlementService) CreateSettlement(ctx *app.Context, settlement *model.Settlement) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	settlement.Uuid = uuid.New().String()
	settlement.CreatedAt = now
	settlement.UpdatedAt = now

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(settlement).Error
		if err != nil {
			ctx.Logger.Error("Failed to create settlement", err)
			return errors.New("failed to create settlement")
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to create settlement", err)
		return errors.New("failed to create settlement")
	}
	return nil
}

func (s *SettlementService) GetSettlementByUUID(ctx *app.Context, uuid string) (*model.Settlement, error) {
	settlement := &model.Settlement{}
	err := ctx.DB.Where("uuid = ?", uuid).First(settlement).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("settlement not found")
		}
		ctx.Logger.Error("Failed to get settlement by UUID", err)
		return nil, errors.New("failed to get settlement by UUID")
	}
	return settlement, nil
}

func (s *SettlementService) UpdateSettlement(ctx *app.Context, settlement *model.Settlement) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	settlement.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", settlement.Uuid).Updates(settlement).Error
	if err != nil {
		ctx.Logger.Error("Failed to update settlement", err)
		return errors.New("failed to update settlement")
	}

	return nil
}

func (s *SettlementService) DeleteSettlement(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.Settlement{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete settlement", err)
		return errors.New("failed to delete settlement")
	}

	return nil
}

// GetSettlementList retrieves a list of settlements based on query parameters
func (s *SettlementService) GetSettlementList(ctx *app.Context, params *model.ReqSettlementQueryParam) (*model.PagedResponse, error) {
	var (
		settlements []*model.Settlement
		total       int64
	)

	db := ctx.DB.Model(&model.Settlement{})

	if params.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+params.OrderNo+"%")
	}
	if params.PurchaseOrderNo != "" {
		db = db.Where("purchase_order_no LIKE ?", "%"+params.PurchaseOrderNo+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get settlement count", err)
		return nil, errors.New("failed to get settlement count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&settlements).Error
	if err != nil {
		ctx.Logger.Error("Failed to get settlement list", err)
		return nil, errors.New("failed to get settlement list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  settlements,
	}, nil
}
