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

	productUuids := make([]string, 0)
	for _, sku := range skuList {
		productUuids = append(productUuids, sku.ProductUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return nil, errors.New("failed to get product list by UUIDs")
	}

	res := make([]*model.SkuRes, 0)
	for _, sku := range skuList {
		skuRes := &model.SkuRes{
			Sku: *sku,
		}
		if product, ok := productMap[sku.ProductUuid]; ok {
			skuRes.Product = *product
		}
		res = append(res, skuRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}

	return
}

// 根据SKU UUID列表获取SKU列表
func (s *SkuService) GetSkuListByUUIDs(ctx *app.Context, uuids []string) (map[string]*model.Sku, error) {
	skuList := make([]*model.Sku, 0)
	err := ctx.DB.Where("uuid in (?)", uuids).Find(&skuList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get SKU list by UUIDs", err)
		return nil, errors.New("failed to get SKU list by UUIDs")
	}

	skuMap := make(map[string]*model.Sku)
	for _, sku := range skuList {
		skuMap[sku.UUID] = sku
	}

	return skuMap, nil
}

// 根据产品uuid获取SKU列表
func (s *SkuService) GetSkuListByProductUUID(ctx *app.Context, productUUID string) (r []*model.Sku, err error) {
	err = ctx.DB.Where("product_uuid = ?", productUUID).Find(&r).Error
	if err != nil {
		ctx.Logger.Error("Failed to get SKU list by product UUID", err)
		return nil, errors.New("failed to get SKU list by product UUID")
	}
	return
}
