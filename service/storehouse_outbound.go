package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseOutboundService struct {
}

func NewStorehouseOutboundService() *StorehouseOutboundService {
	return &StorehouseOutboundService{}
}

func (s *StorehouseOutboundService) CreateOutbound(ctx *app.Context, userId string, req *model.StorehouseOutboundReq) error {

	if req.OutboundDate == "" {
		req.OutboundDate = time.Now().Format("2006-01-02")
	}

	nowstr := time.Now().Format("2006-01-02 15:04:05")
	outbound := &model.StorehouseOutbound{
		StorehouseUuid:        req.StorehouseUuid,
		OutboundType:          req.OutboundType,
		Status:                req.Status,
		OutboundOrderNo:       utils.GenerateOrderID(), // Generating a unique order number
		OutboundDate:          time.Now().Format("2006-01-02"),
		OutboundBy:            userId, // Assuming the user ID is available in the context
		CreatedAt:             time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:             time.Now().Format("2006-01-02 15:04:05"),
		CustomerUuid:          req.CustomerUuid,
		SalesOrderProductType: req.SalesOrderProductType,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(outbound).Error; err != nil {
			ctx.Logger.Error("Failed to create outbound", err)
			return errors.New("failed to create outbound")
		}

		for _, detailReq := range req.Detail {
			detail := &model.StorehouseOutboundDetail{
				Uuid:            uuid.New().String(),
				StorehouseUuid:  req.StorehouseUuid,
				OutboundOrderNo: outbound.OutboundOrderNo,
				ProductUuid:     detailReq.ProductUuid,
				SkuUuid:         detailReq.SkuUuid,
				Quantity:        detailReq.Quantity,
				BoxNum:          detailReq.BoxNum,
				CabinetNo:       detailReq.CabinetNo,
				SalesOrderType:  req.SalesOrderProductType,
				CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt:       time.Now().Format("2006-01-02 15:04:05"),
			}

			// 获取库存

			// 先获取库存
			stock := &model.StorehouseProduct{}
			err := tx.Where("storehouse_uuid = ? AND product_uuid = ? AND sku_uuid = ?", req.StorehouseUuid, detailReq.ProductUuid, detailReq.SkuUuid).First(stock).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.New("仓库中没有该商品")
				}
				ctx.Logger.Error("Failed to get stock", err)
				return err
			}

			// 出库数量大于库存数量
			if stock.Quantity < detailReq.Quantity {
				return errors.New("库存不足")
			}

			if err := tx.Create(detail).Error; err != nil {
				ctx.Logger.Error("Failed to create outbound detail", err)
				return errors.New("failed to create outbound detail")
			}
			beforQuantity := stock.Quantity

			stock.Quantity -= detailReq.Quantity
			stock.UpdatedAt = nowstr
			if err := tx.Where("uuid = ?", stock.Uuid).Updates(stock).Error; err != nil {
				return err
			}

			// 创建库存记录
			stockopLog := &model.StorehouseProductOpLog{
				Uuid:                  uuid.New().String(),
				StorehouseProductUuid: stock.Uuid,
				StorehouseUuid:        outbound.StorehouseUuid,
				BeforeQuantity:        beforQuantity,
				Quantity:              stock.Quantity,
				OpQuantity:            detailReq.Quantity,
				OpType:                model.StorehouseProductOpLogOpTypeOutbound,
				OpDesc:                "仓库出库，减少库存",
				OpBy:                  userId,
				CreatedAt:             nowstr,
			}
			if err := tx.Create(stockopLog).Error; err != nil {
				ctx.Logger.Error("Failed to create stockop log", err)
				return errors.New("failed to create stockop log")
			}

		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *StorehouseOutboundService) GetOutbound(ctx *app.Context, requuid string) (*model.StorehouseOutboundRes, error) {
	outbound := &model.StorehouseOutbound{}
	err := ctx.DB.Where("outbound_order_no = ?", requuid).First(outbound).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("outbound not found")
		}
		ctx.Logger.Error("Failed to get outbound by ID", err)
		return nil, errors.New("failed to get outbound by ID")
	}

	storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, outbound.StorehouseUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, err
	}

	user, err := NewUserService().GetUserByUUID(ctx, outbound.OutboundBy)
	if err != nil {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, err
	}

	outboundRes := &model.StorehouseOutboundRes{
		StorehouseOutbound: *outbound,
		Storehouse:         *storehouse,
		OutboundByUser:     *user,
	}

	return outboundRes, nil
}

// 获取出库明细
func (s *StorehouseOutboundService) GetOutboundDetail(ctx *app.Context, requuid string) ([]*model.StorehouseOutboundDetailRes, error) {
	details := make([]*model.StorehouseOutboundDetail, 0)
	err := ctx.DB.Where("outbound_order_no = ?", requuid).Find(&details).Error
	if err != nil {
		ctx.Logger.Error("Failed to get outbound detail", err)
		return nil, errors.New("failed to get outbound detail")
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	for _, detail := range details {
		productUuids = append(productUuids, detail.ProductUuid)
		skuUuids = append(skuUuids, detail.SkuUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get products by UUIDs", err)
		return nil, err
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get skus by UUIDs", err)
		return nil, err
	}

	res := make([]*model.StorehouseOutboundDetailRes, 0)
	for _, detail := range details {
		detailRes := &model.StorehouseOutboundDetailRes{
			StorehouseOutboundDetail: *detail,
		}
		if product, ok := productMap[detail.ProductUuid]; ok {
			detailRes.Product = *product
		}
		if sku, ok := skuMap[detail.SkuUuid]; ok {
			detailRes.Sku = *sku
		}
		res = append(res, detailRes)
	}

	return res, nil
}

func (s *StorehouseOutboundService) UpdateOutbound(ctx *app.Context, outbound *model.StorehouseOutbound) error {
	outbound.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(outbound).Error
	if err != nil {
		ctx.Logger.Error("Failed to update outbound", err)
		return errors.New("failed to update outbound")
	}
	return nil
}

func (s *StorehouseOutboundService) DeleteOutbound(ctx *app.Context, requuid string) error {

	ctx.DB.Transaction(func(tx *gorm.DB) error {

		// err := tx.Where("outbound_order_no = ?", requuid).Delete(&model.StorehouseOutbound{}).Error
		// if err != nil {
		// 	ctx.Logger.Error("Failed to delete outbound", err)
		// 	return errors.New("failed to delete outbound")
		// }

		// 删除清单
		err := tx.Where("uuid = ?", requuid).Delete(&model.StorehouseOutboundDetail{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete outbound detail", err)
			return errors.New("failed to delete outbound detail")
		}
		return nil
	})

	return nil
}

func (s *StorehouseOutboundService) ListOutbounds(ctx *app.Context, param *model.ReqStorehouseOutboundQueryParam) (r *model.PagedResponse, err error) {
	var (
		outboundList []*model.StorehouseOutbound
		total        int64
	)

	db := ctx.DB.Model(&model.StorehouseOutbound{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&outboundList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	userUuids := make([]string, 0)

	storehouseUuids := make([]string, 0)
	for _, outbound := range outboundList {
		storehouseUuids = append(storehouseUuids, outbound.StorehouseUuid)
		userUuids = append(userUuids, outbound.OutboundBy)
	}

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
		return
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseOutboundRes, 0)
	for _, outbound := range outboundList {
		outboundRes := &model.StorehouseOutboundRes{
			StorehouseOutbound: *outbound,
		}
		if storehouse, ok := storehouseMap[outbound.StorehouseUuid]; ok {
			outboundRes.Storehouse = *storehouse
		}
		if user, ok := userMap[outbound.OutboundBy]; ok {
			outboundRes.OutboundByUser = *user
		}
		res = append(res, outboundRes)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}

// 根据入库单号列表 获取
func (s *StorehouseOutboundService) GetOutboundsByOrderNos(ctx *app.Context, orderNos []string) (map[string]*model.StorehouseOutbound, error) {
	outboundList := make([]*model.StorehouseOutbound, 0)
	err := ctx.DB.Where("outbound_order_no IN ?", orderNos).Find(&outboundList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get inbound list by order nos", err)
		return nil, errors.New("failed to get inbound list by order nos")
	}

	outboundMap := make(map[string]*model.StorehouseOutbound)
	for _, outbound := range outboundList {
		outboundMap[outbound.OutboundOrderNo] = outbound
	}

	return outboundMap, nil
}

func (s *StorehouseOutboundService) ListOutbounds2(ctx *app.Context, param *model.ReqStorehouseOutboundQueryParam) (r *model.PagedResponse, err error) {
	var (
		outboundDetailList []*model.StorehouseOutboundDetail
		total              int64
	)

	db := ctx.DB.Model(&model.StorehouseOutboundDetail{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}
	if param.SalesOrderProductType != "" {
		db = db.Where("sales_order_product_type = ?", param.SalesOrderProductType)
	}

	if param.CustomerUuid != "" {

		var inboundOrders []string
		err = ctx.DB.Model(&model.StorehouseOutbound{}).Where("customer_uuid = ?", param.CustomerUuid).Pluck("outbound_order_no", &inboundOrders).Error
		if err != nil {
			ctx.Logger.Error("Failed to get inbound order no by customer uuid", err)
			return
		}
		db = db.Where("outbound_order_no IN ?", inboundOrders)
	}

	if param.ProductUuid != "" {
		db = db.Where("product_uuid = ?", param.ProductUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&outboundDetailList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	skuuuids := make([]string, 0)
	productUuids := make([]string, 0)
	outboundOrders := make([]string, 0)
	customerUuids := make([]string, 0)
	for _, outboundDetail := range outboundDetailList {
		outboundOrders = append(outboundOrders, outboundDetail.OutboundOrderNo)
		skuuuids = append(skuuuids, outboundDetail.SkuUuid)
		productUuids = append(productUuids, outboundDetail.ProductUuid)

	}

	outboundMap, err := s.GetOutboundsByOrderNos(ctx, outboundOrders)
	if err != nil {
		ctx.Logger.Error("Failed to get inbound list by order nos", err)
		return
	}

	userUuids := make([]string, 0)
	storehouseUuids := make([]string, 0)
	for _, outbound := range outboundMap {
		userUuids = append(userUuids, outbound.OutboundBy)
		storehouseUuids = append(storehouseUuids, outbound.StorehouseUuid)
		customerUuids = append(customerUuids, outbound.CustomerUuid)
	}

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
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

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	curstomeMap, err := NewCustomerService().GetCustomerListByUUIDs(ctx, customerUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get customer list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseOutboundRes, 0)

	for _, outboundDetail := range outboundDetailList {

		outboundDetailRes := &model.StorehouseOutboundDetailRes{
			StorehouseOutboundDetail: *outboundDetail,
		}
		if sku, ok := skuMap[outboundDetail.SkuUuid]; ok {
			outboundDetailRes.Sku = *sku
		}

		if product, ok := productMap[outboundDetail.ProductUuid]; ok {
			outboundDetailRes.Product = *product
		}

		outboundRes := &model.StorehouseOutboundRes{
			StorehouseOutboundDetailRes: *outboundDetailRes,
		}
		if outbound, ok := outboundMap[outboundDetail.OutboundOrderNo]; ok {
			outboundRes.StorehouseOutbound = *outbound
		}

		if storehouse, ok := storehouseMap[outboundRes.StorehouseOutbound.StorehouseUuid]; ok {
			outboundRes.Storehouse = *storehouse
		}

		if user, ok := userMap[outboundRes.StorehouseOutbound.OutboundBy]; ok {
			outboundRes.OutboundByUser = *user
		}

		if customer, ok := curstomeMap[outboundRes.StorehouseOutbound.CustomerUuid]; ok {
			outboundRes.CustomerInfo = *customer
		}

		res = append(res, outboundRes)
	}

	// // for _, inbound := range inboundList {
	// // 	inboundRes := &model.StorehouseInboundRes{
	// // 		StorehouseInbound: *inbound,
	// // 	}
	// // 	if storehouse, ok := storehouseMap[inbound.StorehouseUuid]; ok {
	// // 		inboundRes.Storehouse = *storehouse
	// // 	}

	// // 	if user, ok := userMap[inbound.InboundBy]; ok {
	// // 		inboundRes.InboundByUser = *user
	// // 	}

	// // 	res = append(res, inboundRes)
	// // }

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
