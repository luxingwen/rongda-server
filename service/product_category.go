package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategoryService struct {
}

func NewProductCategoryService() *ProductCategoryService {
	return &ProductCategoryService{}
}

func (s *ProductCategoryService) CreateProductCategory(ctx *app.Context, category *model.ProductCategory) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	category.CreatedAt = now
	category.UpdatedAt = now
	category.Uuid = uuid.New().String()

	err := ctx.DB.Create(category).Error
	if err != nil {
		ctx.Logger.Error("Failed to create product category", err)
		return errors.New("failed to create product category")
	}
	return nil
}

func (s *ProductCategoryService) GetProductCategoryByUUID(ctx *app.Context, uuid string) (*model.ProductCategory, error) {
	category := &model.ProductCategory{}
	err := ctx.DB.Where("uuid = ?", uuid).First(category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("product category not found")
		}
		ctx.Logger.Error("Failed to get product category by UUID", err)
		return nil, errors.New("failed to get product category by UUID")
	}
	return category, nil
}

func (s *ProductCategoryService) UpdateProductCategory(ctx *app.Context, category *model.ProductCategory) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	category.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", category.Uuid).Updates(category).Error
	if err != nil {
		ctx.Logger.Error("Failed to update product category", err)
		return errors.New("failed to update product category")
	}

	return nil
}

func (s *ProductCategoryService) DeleteProductCategory(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.ProductCategory{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete product category", err)
		return errors.New("failed to delete product category")
	}

	return nil
}

// GetProductCategoryList retrieves a list of product categories based on query parameters
func (s *ProductCategoryService) GetProductCategoryList(ctx *app.Context, params *model.ReqProductCategoryQueryParam) (*model.PagedResponse, error) {
	var (
		categories []*model.ProductCategory
		total      int64
	)

	db := ctx.DB.Model(&model.ProductCategory{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get product category count", err)
		return nil, errors.New("failed to get product category count")
	}

	err = db.Find(&categories).Error
	if err != nil {
		ctx.Logger.Error("Failed to get product category list", err)
		return nil, errors.New("failed to get product category list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  categories,
	}, nil
}

// 获取所有可用得产品分类
func (s *ProductCategoryService) GetProductCategoryAll(ctx *app.Context) ([]*model.ProductCategory, error) {
	var categories []*model.ProductCategory
	err := ctx.DB.Where("is_deleted = ?", 0).Find(&categories).Error
	if err != nil {
		ctx.Logger.Error("Failed to get product category list", err)
		return nil, errors.New("failed to get product category list")
	}
	return categories, nil
}

// 根据UUID批量获取产品分类
func (s *ProductCategoryService) GetProductCategoryListByUUIDs(ctx *app.Context, uuids []string) (map[string]*model.ProductCategory, error) {
	categories := make([]*model.ProductCategory, 0)
	err := ctx.DB.Where("uuid in (?)", uuids).Find(&categories).Error
	if err != nil {
		ctx.Logger.Error("Failed to get product category list by UUIDs", err)
		return nil, errors.New("failed to get product category list by UUIDs")
	}

	categoryMap := make(map[string]*model.ProductCategory)
	for _, category := range categories {
		categoryMap[category.Uuid] = category
	}

	return categoryMap, nil
}
