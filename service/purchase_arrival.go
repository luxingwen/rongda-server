package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

	"gorm.io/gorm"
)

type PurchaseArrivalService struct {
}

func NewPurchaseArrivalService() *PurchaseArrivalService {
	return &PurchaseArrivalService{}
}

func (s *PurchaseArrivalService) CreatePurchaseArrival(ctx *app.Context, userId string, req *model.PurchaseArrivalReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	arrival := &model.PurchaseArrival{
		Uuid:             utils.GenerateOrderID(),
		PurchaseOrderNo:  req.PurchaseOrderNo,
		SupplierUuid:     req.SupplierUuid,
		Batch:            req.Batch,
		ArrivalDate:      req.ArrivalDate,
		Acceptor:         req.Acceptor,
		AcceptanceResult: req.AcceptanceResult,
		Remark:           req.Remark,
		Status:           1, // assuming 1 is the default status
		TotalAmount:      0, // this will be calculated
		CreatedAt:        nowStr,
		UpdatedAt:        nowStr,
	}

	// Calculate total amount
	for _, item := range req.Items {
		arrival.TotalAmount += item.TotalAmount
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(arrival).Error; err != nil {
			ctx.Logger.Error("Failed to create purchase arrival", err)
			return errors.New("failed to create purchase arrival")
		}

		for _, itemReq := range req.Items {
			item := &model.PurchaseArrivalItem{
				PurchaseArrivalNo: arrival.Uuid,
				ProductUuid:       itemReq.ProductUuid,
				ProductName:       itemReq.ProductName,
				SkuUuid:           itemReq.SkuUuid,
				SkuName:           itemReq.SkuName,
				Quantity:          itemReq.Quantity,
				Price:             itemReq.Price,
				TotalAmount:       itemReq.TotalAmount,
				CreatedAt:         nowStr,
				UpdatedAt:         nowStr,
			}

			if err := tx.Create(item).Error; err != nil {
				ctx.Logger.Error("Failed to create purchase arrival item", err)
				return errors.New("failed to create purchase arrival item")
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *PurchaseArrivalService) GetPurchaseArrival(ctx *app.Context, requuid string) (*model.PurchaseArrivalRes, error) {
	arrival := &model.PurchaseArrival{}
	err := ctx.DB.Where("uuid = ?", requuid).First(arrival).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("arrival not found")
		}
		ctx.Logger.Error("Failed to get arrival by ID", err)
		return nil, errors.New("failed to get arrival by ID")
	}

	user, err := NewUserService().GetUserByUUID(ctx, arrival.Acceptor)
	if err != nil && err.Error() != "user not found" {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, err
	}

	supplier, err := NewSupplierService().GetSupplierByUUID(ctx, arrival.SupplierUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get supplier by UUID", err)
		return nil, err
	}

	storage, err := NewStorehouseService().GetStorehouseByUUID(ctx, arrival.StorageUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storage by UUID", err)
		return nil, err
	}

	arrivalRes := &model.PurchaseArrivalRes{
		PurchaseArrival: *arrival,
		AcceptorInfo:    user,
		SupplierInfo:    supplier,
		StorageInfo:     storage,
	}

	return arrivalRes, nil
}

// 获取到货明细
func (s *PurchaseArrivalService) GetPurchaseArrivalItems(ctx *app.Context, requuid string) ([]*model.PurchaseArrivalItemRes, error) {
	var items []*model.PurchaseArrivalItem
	err := ctx.DB.Where("purchase_arrival_no = ?", requuid).Find(&items).Error
	if err != nil {
		ctx.Logger.Error("Failed to get arrival items", err)
		return nil, errors.New("failed to get arrival items")
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	for _, item := range items {
		productUuids = append(productUuids, item.ProductUuid)
		skuUuids = append(skuUuids, item.SkuUuid)
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return nil, err

	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return nil, err
	}

	res := make([]*model.PurchaseArrivalItemRes, 0)
	for _, item := range items {
		purchaseArrivalItem := &model.PurchaseArrivalItemRes{
			PurchaseArrivalItem: *item,
		}
		if product, ok := productMap[item.ProductUuid]; ok {
			purchaseArrivalItem.ProductInfo = product
		}
		if sku, ok := skuMap[item.SkuUuid]; ok {
			purchaseArrivalItem.SkuInfo = sku
		}
		res = append(res, purchaseArrivalItem)
	}

	return res, nil
}

func (s *PurchaseArrivalService) UpdatePurchaseArrival(ctx *app.Context, arrival *model.PurchaseArrival) error {
	arrival.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(arrival).Error
	if err != nil {
		ctx.Logger.Error("Failed to update arrival", err)
		return errors.New("failed to update arrival")
	}
	return nil
}

func (s *PurchaseArrivalService) DeletePurchaseArrival(ctx *app.Context, requuid string) error {
	err := ctx.DB.Where("uuid = ?", requuid).Delete(&model.PurchaseArrival{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete arrival", err)
		return errors.New("failed to delete arrival")
	}
	return nil
}

func (s *PurchaseArrivalService) ListPurchaseArrivals(ctx *app.Context, param *model.ReqPurchaseArrivalQueryParam) (r *model.PagedResponse, err error) {
	var (
		arrivalList []*model.PurchaseArrival
		total       int64
	)

	db := ctx.DB.Model(&model.PurchaseArrival{})

	if param.PurchaseOrderNo != "" {
		db = db.Where("purchase_order_no = ?", param.PurchaseOrderNo)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&arrivalList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     arrivalList,
	}
	return
}
