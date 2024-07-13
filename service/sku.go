package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SkuService struct {
}

func NewSkuService() *SkuService {
	return &SkuService{}
}

func (s *SkuService) CreateSku(ctx *app.Context, sku *model.Sku) error {
	sku.UUID = uuid.New().String()

	sku.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	sku.UpdatedAt = sku.CreatedAt

	err := ctx.DB.Create(sku).Error
	if err != nil {
		ctx.Logger.Error("Failed to create SKU", err)
		return errors.New("failed to create SKU")
	}
	return nil
}

func (s *SkuService) GetSkuByUUID(ctx *app.Context, uuid string) (*model.Sku, error) {
	sku := &model.Sku{}
	err := ctx.DB.Where("uuid = ?", uuid).First(sku).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("SKU not found")
		}
		ctx.Logger.Error("Failed to get SKU by UUID", err)
		return nil, errors.New("failed to get SKU by UUID")
	}
	return sku, nil
}

func (s *SkuService) UpdateSku(ctx *app.Context, sku *model.Sku) error {
	sku.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", sku.UUID).Updates(sku).Error
	if err != nil {
		ctx.Logger.Error("Failed to update SKU", err)
		return errors.New("failed to update SKU")
	}

	return nil
}

func (s *SkuService) DeleteSku(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Sku{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete SKU", err)
		return errors.New("failed to delete SKU")
	}

	return nil
}

// 获取SKU列表
func (s *SkuService) GetSkuList(ctx *app.Context, param *model.ReqSkuQueryParam) (r *model.PagedResponse, err error) {
	var (
		skuList []*model.Sku
		total   int64
	)

	db := ctx.DB.Model(&model.Sku{})

	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&skuList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     skuList,
	}

	return
}
