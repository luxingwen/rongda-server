package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplierService struct {
}

func NewSupplierService() *SupplierService {
	return &SupplierService{}
}

func (s *SupplierService) CreateSupplier(ctx *app.Context, supplier *model.Supplier) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	supplier.CreatedAt = now
	supplier.UpdatedAt = now
	supplier.Uuid = uuid.New().String()

	err := ctx.DB.Create(supplier).Error
	if err != nil {
		ctx.Logger.Error("Failed to create supplier", err)
		return errors.New("failed to create supplier")
	}
	return nil
}

func (s *SupplierService) GetSupplierByUUID(ctx *app.Context, uuid string) (*model.Supplier, error) {
	supplier := &model.Supplier{}
	err := ctx.DB.Where("uuid = ?", uuid).First(supplier).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("supplier not found")
		}
		ctx.Logger.Error("Failed to get supplier by UUID", err)
		return nil, errors.New("failed to get supplier by UUID")
	}
	return supplier, nil
}

func (s *SupplierService) UpdateSupplier(ctx *app.Context, supplier *model.Supplier) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	supplier.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", supplier.Uuid).Updates(supplier).Error
	if err != nil {
		ctx.Logger.Error("Failed to update supplier", err)
		return errors.New("failed to update supplier")
	}

	return nil
}

func (s *SupplierService) DeleteSupplier(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete supplier", err)
		return errors.New("failed to delete supplier")
	}

	return nil
}

// GetSupplierList retrieves a list of suppliers based on query parameters
func (s *SupplierService) GetSupplierList(ctx *app.Context, params *model.ReqSupplierQueryParam) (*model.PagedResponse, error) {
	var (
		suppliers []*model.Supplier
		total     int64
	)

	db := ctx.DB.Model(&model.Supplier{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get supplier count", err)
		return nil, errors.New("failed to get supplier count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&suppliers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get supplier list", err)
		return nil, errors.New("failed to get supplier list")
	}

	settlementCurrencyUuids := make([]string, 0)
	for _, supplier := range suppliers {
		settlementCurrencyUuids = append(settlementCurrencyUuids, supplier.SettlementCurrency)
	}

	rSettlementCurrency, err := NewSettlementCurrencyService().GetSettlementCurrencyByUuids(ctx, settlementCurrencyUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get settlement currency by UUID", err)
		return nil, errors.New("failed to get settlement currency by UUID")
	}

	res := make([]*model.SupplierRes, 0)

	for _, supplier := range suppliers {
		supplierRes := &model.SupplierRes{
			Supplier: *supplier,
		}
		if currency, ok := rSettlementCurrency[supplier.SettlementCurrency]; ok {
			supplierRes.SettlementCurrencyInfo = currency
		}
		res = append(res, supplierRes)
	}

	return &model.PagedResponse{
		Total: total,
		Data:  res,
	}, nil
}

// 获取所有可用的供应商
func (s *SupplierService) GetAvailableSupplierList(ctx *app.Context) (r []*model.Supplier, err error) {
	err = ctx.DB.Where("status = ?", 1).Find(&r).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available supplier list", err)
		return nil, errors.New("failed to get available supplier list")
	}
	return
}

// 根据uuid列表获取供应商列表
func (s *SupplierService) GetSupplierListByUUIDs(ctx *app.Context, uuids []string) (r map[string]*model.Supplier, err error) {
	var supplierList []*model.Supplier
	r = make(map[string]*model.Supplier)

	err = ctx.DB.Where("uuid in ?", uuids).Find(&supplierList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get supplier list by UUIDs", err)
		return nil, errors.New("failed to get supplier list by UUIDs")
	}

	for _, supplier := range supplierList {
		r[supplier.Uuid] = supplier
	}

	return
}
