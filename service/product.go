package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) CreateProduct(ctx *app.Context, product *model.Product) error {
	product.Uuid = uuid.New().String()
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	product.UpdatedAt = product.CreatedAt

	err := ctx.DB.Create(product).Error
	if err != nil {
		ctx.Logger.Error("Failed to create product", err)
		return errors.New("failed to create product")
	}
	return nil
}

func (s *ProductService) GetProductByUUID(ctx *app.Context, uuid string) (*model.Product, error) {
	product := &model.Product{}
	err := ctx.DB.Where("uuid = ?", uuid).First(product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("product not found")
		}
		ctx.Logger.Error("Failed to get product by UUID", err)
		return nil, errors.New("failed to get product by UUID")
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx *app.Context, product *model.Product) error {
	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", product.Uuid).Updates(product).Error
	if err != nil {
		ctx.Logger.Error("Failed to update product", err)
		return errors.New("failed to update product")
	}
	return nil
}

func (s *ProductService) DeleteProduct(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Product{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete product", err)
		return errors.New("failed to delete product")
	}
	return nil
}

// 获取产品列表
func (s *ProductService) GetProductList(ctx *app.Context, param *model.ReqProductQueryParam) (r *model.PagedResponse, err error) {
	var (
		productList []*model.Product
		total       int64
	)

	db := ctx.DB.Model(&model.Product{})

	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&productList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     productList,
	}
	return
}
