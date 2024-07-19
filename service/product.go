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
	err := ctx.DB.Where("uuid = ?", uuid).Update("is_deleted", 1).Error
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

	db = db.Where("is_deleted = ?", 0)

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&productList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	supplierUuids := make([]string, 0)
	for _, product := range productList {
		supplierUuids = append(supplierUuids, product.Supplier)
	}

	supplierMap, err := NewSupplierService().GetSupplierListByUUIDs(ctx, supplierUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get supplier list by UUIDs", err)
		return
	}

	res := make([]*model.ProductRes, 0)
	for _, product := range productList {
		productRes := &model.ProductRes{
			Product: *product,
		}
		if supplier, ok := supplierMap[product.Supplier]; ok {
			productRes.SupplierInfo = supplier
		}
		res = append(res, productRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}

// 获取所有可用的产品
func (s *ProductService) GetAvailableProductList(ctx *app.Context) (r []*model.Product, err error) {
	err = ctx.DB.Find(&r).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available product list", err)
		return nil, errors.New("failed to get available product list")
	}
	return
}

// 根据uuid列表获取产品列表
func (s *ProductService) GetProductListByUUIDs(ctx *app.Context, uuids []string) (r map[string]*model.Product, err error) {
	var productList []*model.Product
	r = make(map[string]*model.Product)

	err = ctx.DB.Where("uuid in ?", uuids).Find(&productList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return nil, errors.New("failed to get product list by UUIDs")
	}

	for _, product := range productList {
		r[product.Uuid] = product
	}

	return
}
