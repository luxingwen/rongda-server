package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SettlementCurrencyService struct {
}

func NewSettlementCurrencyService() *SettlementCurrencyService {
	return &SettlementCurrencyService{}
}

func (s *SettlementCurrencyService) CreateSettlementCurrency(ctx *app.Context, currency *model.SettlementCurrency) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	currency.CreatedAt = now
	currency.UpdatedAt = now
	currency.Uuid = uuid.New().String()

	err := ctx.DB.Create(currency).Error
	if err != nil {
		ctx.Logger.Error("Failed to create settlement currency", err)
		return errors.New("failed to create settlement currency")
	}
	return nil
}

func (s *SettlementCurrencyService) GetSettlementCurrencyByUUID(ctx *app.Context, uuid string) (*model.SettlementCurrency, error) {
	currency := &model.SettlementCurrency{}
	err := ctx.DB.Where("uuid = ?", uuid).First(currency).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("settlement currency not found")
		}
		ctx.Logger.Error("Failed to get settlement currency by UUID", err)
		return nil, errors.New("failed to get settlement currency by UUID")
	}
	return currency, nil
}

func (s *SettlementCurrencyService) UpdateSettlementCurrency(ctx *app.Context, currency *model.SettlementCurrency) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	currency.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", currency.Uuid).Updates(currency).Error
	if err != nil {
		ctx.Logger.Error("Failed to update settlement currency", err)
		return errors.New("failed to update settlement currency")
	}

	return nil
}

func (s *SettlementCurrencyService) DeleteSettlementCurrency(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.SettlementCurrency{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete settlement currency", err)
		return errors.New("failed to delete settlement currency")
	}

	return nil
}

// GetSettlementCurrencyList retrieves a list of settlement currencies based on query parameters
func (s *SettlementCurrencyService) GetSettlementCurrencyList(ctx *app.Context, params *model.ReqSettlementCurrencyQueryParam) (*model.PagedResponse, error) {
	var (
		currencies []*model.SettlementCurrency
		total      int64
	)

	db := ctx.DB.Model(&model.SettlementCurrency{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get settlement currency count", err)
		return nil, errors.New("failed to get settlement currency count")
	}

	err = db.Find(&currencies).Error
	if err != nil {
		ctx.Logger.Error("Failed to get settlement currency list", err)
		return nil, errors.New("failed to get settlement currency list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  currencies,
	}, nil
}

// 获取全部可用的结算币种
func (s *SettlementCurrencyService) GetAvailableSettlementCurrencyList(ctx *app.Context) ([]*model.SettlementCurrency, error) {
	var currencies []*model.SettlementCurrency
	err := ctx.DB.Model(&model.SettlementCurrency{}).Where("status = ?", model.SettlementCurrencyStatusEnabled).Find(&currencies).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available settlement currency list", err)
		return nil, errors.New("failed to get available settlement currency list")
	}
	return currencies, nil
}

// 根据uuid获取结算币种
func (s *SettlementCurrencyService) GetSettlementCurrencyByUuids(ctx *app.Context, uuids []string) (map[string]*model.SettlementCurrency, error) {
	currencies := make([]*model.SettlementCurrency, 0)
	err := ctx.DB.Model(&model.SettlementCurrency{}).Where("uuid IN (?)", uuids).Find(&currencies).Error
	if err != nil {
		ctx.Logger.Error("Failed to get settlement currency by uuid", err)
		return nil, errors.New("failed to get settlement currency by uuid")
	}
	result := make(map[string]*model.SettlementCurrency)
	for _, currency := range currencies {
		result[currency.Uuid] = currency
	}
	return result, nil
}
