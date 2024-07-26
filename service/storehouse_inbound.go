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

type StorehouseInboundService struct {
}

func NewStorehouseInboundService() *StorehouseInboundService {
	return &StorehouseInboundService{}
}

func (s *StorehouseInboundService) CreateInbound(ctx *app.Context, userId string, req *model.StorehouseInboundReq) error {
	nowstr := time.Now().Format("2006-01-02 15:04:05")

	if req.InDate == "" {
		req.InDate = time.Now().Format("2006-01-02")
	}

	inbound := &model.StorehouseInbound{
		StorehouseUuid:  req.StorehouseUuid,
		Title:           req.Title,
		InboundType:     req.InboundType,
		Status:          req.Status,
		InboundOrderNo:  utils.GenerateOrderID(),
		InboundDate:     req.InDate,
		InboundBy:       userId,
		PurchaseOrderNo: req.PurchaseOrderNo,
		CustomerUuid:    req.CustomerUuid,
		CreatedAt:       nowstr,
		UpdatedAt:       nowstr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(inbound).Error; err != nil {
			return err
		}

		for _, detailReq := range req.Detail {
			detail := &model.StorehouseInboundDetail{
				Uuid:                     uuid.New().String(),
				PurchaseOrderProductNo:   detailReq.PurchaseOrderProductNo,
				StorehouseUuid:           req.StorehouseUuid,
				InboundOrderNo:           inbound.InboundOrderNo,
				ProductUuid:              detailReq.ProductUuid,
				SkuUuid:                  detailReq.SkuUuid,
				Quantity:                 detailReq.Quantity,
				BoxNum:                   detailReq.BoxNum,
				CabinetNo:                detailReq.CabinetNo,
				PurchaseOrderProductType: req.PurchaseOrderProductType,
				CreatedAt:                nowstr,
				UpdatedAt:                nowstr,
			}
			if err := tx.Create(detail).Error; err != nil {
				return err
			}

			// 创建库存
			stock := &model.StorehouseProduct{}

			stock.Uuid = uuid.New().String()
			stock.StorehouseUuid = req.StorehouseUuid
			stock.ProductUuid = detailReq.ProductUuid
			stock.PurchaseOrderNo = req.PurchaseOrderNo
			stock.SkuUuid = detailReq.SkuUuid
			stock.Quantity = detailReq.Quantity
			stock.BoxNum = detailReq.BoxNum
			stock.CabinetNo = detailReq.CabinetNo
			stock.InDate = req.InDate
			stock.CustomerUuid = req.CustomerUuid
			stock.PurchaseOrderProductType = req.PurchaseOrderProductType
			stock.PurchaseOrderProductNo = detailReq.PurchaseOrderProductNo
			stock.CreatedAt = nowstr
			stock.UpdatedAt = nowstr
			if err := tx.Create(stock).Error; err != nil {
				ctx.Logger.Error("Failed to create stock", err)
				return err
			}

			// 创建库存记录
			stockopLog := &model.StorehouseProductOpLog{
				Uuid:                  uuid.New().String(),
				InboundOrderNo:        inbound.InboundOrderNo,
				InboudItemDdetailUuid: detail.Uuid,
				StorehouseProductUuid: stock.Uuid,
				StorehouseUuid:        inbound.StorehouseUuid,
				BeforeQuantity:        0,
				Quantity:              detail.Quantity,
				OpQuantity:            detail.Quantity,
				OpType:                1,
				BoxNum:                detail.BoxNum,
				OpDesc:                "采购入库",
				OpBy:                  userId,
				CreatedAt:             nowstr,
			}

			if err := tx.Create(stockopLog).Error; err != nil {
				ctx.Logger.Error("Failed to create stockopLog", err)
				return err
			}

		}

		// 更新采购单状态

		err := tx.Model(&model.PurchaseOrder{}).Where("order_no = ?", req.PurchaseOrderNo).Update("status", model.PurchaseOrderStatusInStore).Error
		if err != nil {
			ctx.Logger.Error("Failed to update purchase order status", err)
			return err
		}

		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to create inbound", err)
		return errors.New("failed to create inbound")
	}

	return nil
}

func (s *StorehouseInboundService) GetInbound(ctx *app.Context, uuid string) (*model.StorehouseInboundRes, error) {
	inbound := &model.StorehouseInbound{}
	err := ctx.DB.Where("inbound_order_no = ?", uuid).First(inbound).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("inbound not found")
		}
		ctx.Logger.Error("Failed to get inbound by UUID", err)
		return nil, errors.New("failed to get inbound by UUID")
	}

	storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, inbound.StorehouseUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, errors.New("failed to get storehouse by UUID")
	}

	user, err := NewUserService().GetUserByUUID(ctx, inbound.InboundBy)
	if err != nil {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, errors.New("failed to get user by UUID")
	}

	inboundRes := &model.StorehouseInboundRes{
		StorehouseInbound: *inbound,
		Storehouse:        *storehouse,
		InboundByUser:     *user,
	}

	return inboundRes, nil
}

// 获取入库明细
func (s *StorehouseInboundService) GetInboundDetailByInboundOrderNo(ctx *app.Context, uuid string) ([]*model.StorehouseInboundDetailRes, error) {
	var inboundDetail []*model.StorehouseInboundDetail
	err := ctx.DB.Where("inbound_order_no = ?", uuid).Find(&inboundDetail).Error
	if err != nil {
		ctx.Logger.Error("Failed to get inbound detail", err)
		return nil, errors.New("failed to get inbound detail")
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	for _, detail := range inboundDetail {
		productUuids = append(productUuids, detail.ProductUuid)
		skuUuids = append(skuUuids, detail.SkuUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return nil, errors.New("failed to get product list by UUIDs")
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return nil, errors.New("failed to get sku list by UUIDs")
	}

	res := make([]*model.StorehouseInboundDetailRes, 0)
	for _, detail := range inboundDetail {
		detailRes := &model.StorehouseInboundDetailRes{
			StorehouseInboundDetail: *detail,
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

// 获取入库明细
func (s *StorehouseInboundService) GetInboundDetailInfo(ctx *app.Context, uuid string) (*model.StorehouseInboundDetailInfoRes, error) {
	var inboundDetail model.StorehouseInboundDetail
	err := ctx.DB.Where("uuid = ?", uuid).First(&inboundDetail).Error
	if err != nil {
		ctx.Logger.Error("Failed to get inbound detail", err)
		return nil, errors.New("failed to get inbound detail")
	}

	product, err := NewProductService().GetProductByUUID(ctx, inboundDetail.ProductUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get product by UUID", err)
		return nil, errors.New("failed to get product by UUID")
	}

	sku, err := NewSkuService().GetSkuByUUID(ctx, inboundDetail.SkuUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get sku by UUID", err)
		return nil, errors.New("failed to get sku by UUID")
	}

	storehouseInboud, err := s.GetInbound(ctx, inboundDetail.InboundOrderNo)
	if err != nil {
		ctx.Logger.Error("Failed to get inbound by UUID", err)
		return nil, errors.New("failed to get inbound by UUID")
	}

	purchaseOrder, err := NewPurchaseOrderService().GetPurchaseOrderRecord(ctx, storehouseInboud.StorehouseInbound.PurchaseOrderNo)
	if err != nil {
		ctx.Logger.Error("Failed to get purchase order by UUID", err)
		return nil, errors.New("failed to get purchase order by UUID")
	}

	customer, err := NewCustomerService().GetCustomerByUUID(ctx, storehouseInboud.StorehouseInbound.CustomerUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get customer by UUID", err)
		return nil, errors.New("failed to get customer by UUID")
	}

	res := &model.StorehouseInboundDetailInfoRes{
		StorehouseInboundDetail: inboundDetail,
		Product:                 *product,
		Sku:                     *sku,
		StorehouseInbound:       *storehouseInboud,
		PurchaseOrderInfo:       *purchaseOrder,
		CustomerInfo:            *customer,
	}

	return res, nil
}

func (s *StorehouseInboundService) UpdateInbound(ctx *app.Context, req *model.StorehouseInboundUpdateReq) error {
	inbound := &model.StorehouseInbound{
		StorehouseUuid: req.StorehouseUuid,
		Title:          req.Title,
		InboundType:    req.InboundType,
		Status:         req.Status,
		InboundDate:    time.Now().Format("2006-01-02"),
		UpdatedAt:      time.Now().Format("2006-01-02 15:04:05"),
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("inbound_order_no = ?", req.InboundOrderNo).Updates(inbound).Error; err != nil {
			return err
		}

		if err := tx.Where("inbound_order_no = ?", req.InboundOrderNo).Delete(&model.StorehouseInboundDetail{}).Error; err != nil {
			return err
		}

		for _, detailReq := range req.Detail {
			detail := &model.StorehouseInboundDetail{
				InboundOrderNo: req.InboundOrderNo,
				ProductUuid:    detailReq.ProductUuid,
				SkuUuid:        detailReq.SkuUuid,
				Quantity:       detailReq.Quantity,
			}
			if err := tx.Create(detail).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to update inbound", err)
		return errors.New("failed to update inbound")
	}

	return nil
}

func (s *StorehouseInboundService) DeleteInbound(ctx *app.Context, uuid string) error {
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		// if err := tx.Where("inbound_order_no = ?", uuid).Delete(&model.StorehouseInbound{}).Error; err != nil {
		// 	return err
		// }
		if err := tx.Where("uuid = ?", uuid).Delete(&model.StorehouseInboundDetail{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to delete inbound", err)
		return errors.New("failed to delete inbound")
	}

	return nil
}

func (s *StorehouseInboundService) ListInbounds(ctx *app.Context, param *model.ReqStorehouseInboundQueryParam) (r *model.PagedResponse, err error) {
	var (
		inboundList []*model.StorehouseInbound
		total       int64
	)

	db := ctx.DB.Model(&model.StorehouseInbound{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&inboundList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	userUuids := make([]string, 0)
	for _, inbound := range inboundList {
		userUuids = append(userUuids, inbound.InboundBy)
	}

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
		return
	}

	storehouseUuids := make([]string, 0)
	for _, inbound := range inboundList {
		storehouseUuids = append(storehouseUuids, inbound.StorehouseUuid)
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}
	res := make([]*model.StorehouseInboundRes, 0)
	for _, inbound := range inboundList {
		inboundRes := &model.StorehouseInboundRes{
			StorehouseInbound: *inbound,
		}
		if storehouse, ok := storehouseMap[inbound.StorehouseUuid]; ok {
			inboundRes.Storehouse = *storehouse
		}

		if user, ok := userMap[inbound.InboundBy]; ok {
			inboundRes.InboundByUser = *user
		}

		res = append(res, inboundRes)
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
func (s *StorehouseInboundService) GetInboundsByOrderNos(ctx *app.Context, orderNos []string) (map[string]*model.StorehouseInbound, error) {
	inboundList := make([]*model.StorehouseInbound, 0)
	err := ctx.DB.Where("inbound_order_no IN ?", orderNos).Find(&inboundList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get inbound list by order nos", err)
		return nil, errors.New("failed to get inbound list by order nos")
	}

	inboundMap := make(map[string]*model.StorehouseInbound)
	for _, inbound := range inboundList {
		inboundMap[inbound.InboundOrderNo] = inbound
	}

	return inboundMap, nil
}

func (s *StorehouseInboundService) ListInbounds2(ctx *app.Context, param *model.ReqStorehouseInboundQueryParam) (r *model.PagedResponse, err error) {
	var (
		inboundDetailList []*model.StorehouseInboundDetail
		total             int64
	)

	db := ctx.DB.Model(&model.StorehouseInboundDetail{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}
	if param.PurchaseOrderProductType != "" {
		db = db.Where("purchase_order_product_type = ?", param.PurchaseOrderProductType)
	}

	if param.CustomerUuid != "" {

		var inboundOrders []string
		err = ctx.DB.Model(&model.StorehouseInbound{}).Where("customer_uuid = ?", param.CustomerUuid).Pluck("inbound_order_no", &inboundOrders).Error
		if err != nil {
			ctx.Logger.Error("Failed to get inbound order no by customer uuid", err)
			return
		}
		db = db.Where("inbound_order_no IN ?", inboundOrders)
	}

	if param.ProductUuid != "" {
		db = db.Where("product_uuid = ?", param.ProductUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&inboundDetailList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	skuuuids := make([]string, 0)
	productUuids := make([]string, 0)
	inboundOrders := make([]string, 0)
	customerUuids := make([]string, 0)
	for _, inboundDetail := range inboundDetailList {
		inboundOrders = append(inboundOrders, inboundDetail.InboundOrderNo)
		skuuuids = append(skuuuids, inboundDetail.SkuUuid)
		productUuids = append(productUuids, inboundDetail.ProductUuid)

	}

	inboundMap, err := s.GetInboundsByOrderNos(ctx, inboundOrders)
	if err != nil {
		ctx.Logger.Error("Failed to get inbound list by order nos", err)
		return
	}

	userUuids := make([]string, 0)
	storehouseUuids := make([]string, 0)
	for _, inbound := range inboundMap {
		userUuids = append(userUuids, inbound.InboundBy)
		storehouseUuids = append(storehouseUuids, inbound.StorehouseUuid)
		customerUuids = append(customerUuids, inbound.CustomerUuid)
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

	res := make([]*model.StorehouseInboundItemRes, 0)

	for _, inboundDetail := range inboundDetailList {

		inboundDetailRes := &model.StorehouseInboundDetailRes{
			StorehouseInboundDetail: *inboundDetail,
		}
		if sku, ok := skuMap[inboundDetail.SkuUuid]; ok {
			inboundDetailRes.Sku = *sku
		}

		if product, ok := productMap[inboundDetail.ProductUuid]; ok {
			inboundDetailRes.Product = *product
		}

		inboundRes := &model.StorehouseInboundItemRes{
			StorehouseInboundDetailRes: *inboundDetailRes,
		}
		if inbound, ok := inboundMap[inboundDetail.InboundOrderNo]; ok {
			inboundRes.StorehouseInbound = *inbound
		}

		if storehouse, ok := storehouseMap[inboundRes.StorehouseInbound.StorehouseUuid]; ok {
			inboundRes.Storehouse = *storehouse
		}

		if user, ok := userMap[inboundRes.StorehouseInbound.InboundBy]; ok {
			inboundRes.InboundByUser = *user
		}

		if customer, ok := curstomeMap[inboundRes.StorehouseInbound.CustomerUuid]; ok {
			inboundRes.CustomerInfo = *customer
		}

		res = append(res, inboundRes)
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
