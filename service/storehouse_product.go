package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseProductService struct {
}

func NewStorehouseProductService() *StorehouseProductService {
	return &StorehouseProductService{}
}

func (s *StorehouseProductService) CreateProduct(ctx *app.Context, userId string, product *model.StorehouseProduct) error {
	product.Uuid = uuid.New().String()
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	product.UpdatedAt = product.CreatedAt

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		// 查询物品是否存在该仓库

		resProduct := &model.StorehouseProduct{}
		err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", product.StorehouseUuid, product.ProductUuid, product.SkuUuid).First(resProduct).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 不存在则创建

				if err := tx.Create(product).Error; err != nil {
					ctx.Logger.Error("Failed to create product", err)
					return errors.New("failed to create product")
				}

				// 创建库存记录
				stockopLog := &model.StorehouseProductOpLog{
					Uuid:                  uuid.New().String(),
					StorehouseUuid:        product.StorehouseUuid,
					StorehouseProductUuid: product.Uuid,
					BeforeQuantity:        0,
					Quantity:              product.Quantity,
					OpQuantity:            product.Quantity,
					OpType:                model.StorehouseProductOpLogOpTypeInbound,
					OpDesc:                "入库",
					OpBy:                  userId,
					CreatedAt:             product.CreatedAt,
				}
				if err := tx.Create(stockopLog).Error; err != nil {
					ctx.Logger.Error("Failed to create stockop log", err)
					return errors.New("failed to create stockop log")
				}
				return nil

			} else {
				return err
			}
		}
		// 存在则更新
		beforQuantity := resProduct.Quantity
		resProduct.Quantity += product.Quantity
		resProduct.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		err = tx.Where("uuid = ?", resProduct.Uuid).Updates(resProduct).Error
		if err != nil {
			ctx.Logger.Error("Failed to update product", err)
			return errors.New("failed to update product")
		}

		// 创建库存记录
		stockopLog := &model.StorehouseProductOpLog{
			Uuid:                  uuid.New().String(),
			StorehouseProductUuid: product.Uuid,
			StorehouseUuid:        product.StorehouseUuid,
			BeforeQuantity:        beforQuantity,
			Quantity:              resProduct.Quantity,
			OpQuantity:            product.Quantity,
			OpType:                model.StorehouseProductOpLogOpTypeInbound,
			OpDesc:                "增加库存",
			OpBy:                  userId,
			CreatedAt:             product.CreatedAt,
		}
		if err := tx.Create(stockopLog).Error; err != nil {
			ctx.Logger.Error("Failed to create stockop log", err)
			return errors.New("failed to create stockop log")
		}

		return nil
	})
	if err != nil {
		ctx.Logger.Error("Failed to create product", err)
		return errors.New("failed to create product")
	}
	return nil
}

func (s *StorehouseProductService) GetProduct(ctx *app.Context, uuid string) (*model.StorehouseProductRes, error) {
	product := &model.StorehouseProduct{}
	err := ctx.DB.Where("uuid = ?", uuid).First(product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("product not found")
		}
		ctx.Logger.Error("Failed to get product by ID", err)
		return nil, errors.New("failed to get product by ID")
	}

	storehouseInfo, err := NewStorehouseService().GetStorehouseByUUID(ctx, product.StorehouseUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, errors.New("failed to get storehouse by UUID")
	}

	productInfo, err := NewProductService().GetProductByUUID(ctx, product.ProductUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get product by UUID", err)
		return nil, errors.New("failed to get product by UUID")
	}
	skuInfo, err := NewSkuService().GetSkuByUUID(ctx, product.SkuUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get sku by UUID", err)
		return nil, errors.New("failed to get sku by UUID")
	}

	res := &model.StorehouseProductRes{
		StorehouseProduct: *product,
		Storehouse:        *storehouseInfo,
		Product:           *productInfo,
		Sku:               *skuInfo,
	}

	return res, nil
}

// 分页获取获取物品操作记录
func (s *StorehouseProductService) GetProductOpLog(ctx *app.Context, param *model.ReqStorehouseProductOpLogQueryParam) (r *model.PagedResponse, err error) {
	var (
		opLogList []*model.StorehouseProductOpLog
		total     int64
	)

	db := ctx.DB.Model(&model.StorehouseProductOpLog{})
	if param.Uuid != "" {
		db = db.Where("storehouse_product_uuid = ?", param.Uuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Order("id DESC").Find(&opLogList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	userUuids := make([]string, 0)
	for _, v := range opLogList {
		userUuids = append(userUuids, v.OpBy)
	}
	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseProductOpLogRes, 0)
	for _, v := range opLogList {
		opLogItem := &model.StorehouseProductOpLogRes{
			StorehouseProductOpLog: *v,
		}
		if user, ok := userMap[v.OpBy]; ok {
			opLogItem.OpByUser = *user
		}
		res = append(res, opLogItem)
	}

	return &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}, nil
}

func (s *StorehouseProductService) UpdateProduct(ctx *app.Context, userId string, product *model.StorehouseProduct) error {
	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	ctx.DB.Transaction(func(tx *gorm.DB) error {

		// 先获取原来的库存
		resProduct := &model.StorehouseProduct{}
		err := tx.Where("uuid = ?", product.Uuid).First(resProduct).Error
		if err != nil {
			ctx.Logger.Error("Failed to get product by ID", err)
			return errors.New("failed to get product by ID")
		}

		err = tx.Where("uuid = ?", product.Uuid).Updates(product).Error
		if err != nil {
			ctx.Logger.Error("Failed to update product", err)
			return errors.New("failed to update product")
		}

		if resProduct.Quantity == product.Quantity {
			return nil
		}

		// 创建库存记录
		stockopLog := &model.StorehouseProductOpLog{
			Uuid:                  uuid.New().String(),
			StorehouseProductUuid: product.Uuid,
			StorehouseUuid:        product.StorehouseUuid,
			BeforeQuantity:        resProduct.Quantity,
			Quantity:              product.Quantity,
			OpQuantity:            product.Quantity,
			OpType:                model.StorehouseProductOpLogOpTypeUpdate,
			OpDesc:                "更新库存",
			OpBy:                  userId,
			CreatedAt:             product.CreatedAt,
		}
		if err := tx.Create(stockopLog).Error; err != nil {
			ctx.Logger.Error("Failed to create stockop log", err)
			return errors.New("failed to create stockop log")
		}
		return nil
	})

	return nil
}

func (s *StorehouseProductService) DeleteProduct(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.StorehouseProduct{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete product", err)
		return errors.New("failed to delete product")
	}
	return nil
}

func (s *StorehouseProductService) ListProducts(ctx *app.Context, param *model.ReqStorehouseProductQueryParam) (r *model.PagedResponse, err error) {
	var (
		productList []*model.StorehouseProduct
		total       int64
	)

	db := ctx.DB.Model(&model.StorehouseProduct{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&productList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	productUuids := make([]string, 0)
	for _, v := range productList {
		productUuids = append(productUuids, v.ProductUuid)
	}
	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return
	}

	skuUuids := make([]string, 0)
	for _, v := range productList {
		skuUuids = append(skuUuids, v.SkuUuid)
	}
	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return
	}

	storehouseUuids := make([]string, 0)
	for _, v := range productList {
		storehouseUuids = append(storehouseUuids, v.StorehouseUuid)
	}
	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseProductRes, 0)
	for _, v := range productList {
		productItem := &model.StorehouseProductRes{
			StorehouseProduct: *v,
		}
		if product, ok := productMap[v.ProductUuid]; ok {
			productItem.Product = *product
		}
		if sku, ok := skuMap[v.SkuUuid]; ok {
			productItem.Sku = *sku
		}
		if storehouse, ok := storehouseMap[v.StorehouseUuid]; ok {
			productItem.Storehouse = *storehouse
		}
		res = append(res, productItem)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
