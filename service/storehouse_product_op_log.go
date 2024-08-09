package service

import (
	"sgin/model"
	"sgin/pkg/app"
)

type StorehouseProductOpLogService struct {
}

func NewStorehouseProductOpLogService() *StorehouseProductOpLogService {
	return &StorehouseProductOpLogService{}
}

// 获取库存操作日志列表
func (s *StorehouseProductOpLogService) GetStorehouseProductOpLogList(ctx *app.Context, params *model.ReqStorehouseProductOpLogListParam) (*model.PagedResponse, error) {
	var list []model.StorehouseProductOpLog
	var total int64
	query := ctx.DB.Model(&model.StorehouseProductOpLog{})

	if len(params.OpTypes) > 0 {
		query = query.Where("op_type IN (?)", params.OpTypes)
	}

	if params.TeamUuid != "" {
		query = query.Where("team_uuid = ?", params.TeamUuid)
	}

	err := query.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse product op log count", err)
		return nil, err
	}

	err = query.Offset(params.GetOffset()).Limit(params.PageSize).Find(&list).Error
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse product op log list", err)
		return nil, err
	}

	productUuids := make([]string, 0)
	skuUuid := make([]string, 0)
	storehouseUuids := make([]string, 0)

	for _, v := range list {
		productUuids = append(productUuids, v.ProductUuid)
		skuUuid = append(skuUuid, v.SkuUuid)
		storehouseUuids = append(storehouseUuids, v.StorehouseUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by uuids", err)
		return nil, err
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuid)
	if err != nil {

		ctx.Logger.Error("Failed to get sku list by uuids", err)
		return nil, err
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by uuids", err)
		return nil, err
	}

	res := make([]*model.StorehouseProductOpLogRes, 0)
	for _, v := range list {
		oplogItem := model.StorehouseProductOpLogRes{
			StorehouseProductOpLog: v,
		}
		if product, ok := productMap[v.ProductUuid]; ok {
			oplogItem.ProductInfo = *product
		}
		if sku, ok := skuMap[v.SkuUuid]; ok {
			oplogItem.SkuInfo = *sku
		}
		if storehouse, ok := storehouseMap[v.StorehouseUuid]; ok {
			oplogItem.StorehouseInfo = *storehouse
		}
		res = append(res, &oplogItem)
	}
	return &model.PagedResponse{
		Total: total,
		Data:  res,
	}, nil

}
