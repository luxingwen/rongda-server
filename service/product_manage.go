package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductManageService struct {
}

func NewProductManageService() *ProductManageService {
	return &ProductManageService{}
}

func (s *ProductManageService) CreateProduct(ctx *app.Context, product *model.ProductManage) error {
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

func (s *ProductManageService) GetProductByUUID(ctx *app.Context, uuid string) (*model.ProductManage, error) {
	product := &model.ProductManage{}
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

func (s *ProductManageService) UpdateProduct(ctx *app.Context, product *model.ProductManage) error {
	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", product.Uuid).Updates(product).Error
	if err != nil {
		ctx.Logger.Error("Failed to update product", err)
		return errors.New("failed to update product")
	}
	return nil
}

func (s *ProductManageService) DeleteProduct(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.ProductManage{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete product", err)
		return errors.New("failed to delete product")
	}
	return nil
}

func (s *ProductManageService) GetProductList(ctx *app.Context, param *model.ReqProductQueryParam) (r *model.PagedResponse, err error) {
	var (
		productList []*model.ProductManage
		total       int64
	)

	db := ctx.DB.Model(&model.ProductManage{})

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

	productUuids := make([]string, 0)
	skuuuids := make([]string, 0)
	supplierUuids := make([]string, 0)
	for _, product := range productList {
		supplierUuids = append(supplierUuids, product.Supplier)
		skuuuids = append(skuuuids, product.Sku)
		productUuids = append(productUuids, product.Product)
	}

	supplierMap, err := NewSupplierService().GetSupplierListByUUIDs(ctx, supplierUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get supplier list by UUIDs", err)
		return
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuuuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return
	}

	res := make([]*model.ProductManageRes, 0)
	for _, product := range productList {
		productRes := &model.ProductManageRes{
			ProductManage: *product,
		}
		if supplier, ok := supplierMap[product.Supplier]; ok {
			productRes.SupplierInfo = supplier
		}
		if sku, ok := skuMap[product.Sku]; ok {
			productRes.SkuInfo = sku
		}
		if product, ok := productMap[product.Product]; ok {
			productRes.ProductInfo = product
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

func (s *ProductManageService) GetAvailableProductList(ctx *app.Context) (r []*model.ProductManage, err error) {
	err = ctx.DB.Find(&r).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available product list", err)
		return nil, errors.New("failed to get available product list")
	}
	return
}

func (s *ProductManageService) GetProductManageListByUUIDs(ctx *app.Context, uuids []string) (r map[string]*model.ProductManage, err error) {
	var productList []*model.ProductManage
	r = make(map[string]*model.ProductManage)

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
