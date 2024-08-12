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

type StorehouseInventoryCheckService struct {
}

func NewStorehouseInventoryCheckService() *StorehouseInventoryCheckService {
	return &StorehouseInventoryCheckService{}
}

func (s *StorehouseInventoryCheckService) CreateInventoryCheck(ctx *app.Context, userId string, req *model.StorehouseInventoryCheckReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	check := &model.StorehouseInventoryCheck{
		StorehouseUuid: req.StorehouseUuid,
		CheckOrderNo:   utils.GenerateOrderID(),
		CheckDate:      req.CheckDate,
		Status:         req.Status,
		CheckBy:        userId,
		CreatedAt:      nowStr,
		UpdatedAt:      nowStr,
	}

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(check).Error; err != nil {
			ctx.Logger.Error("Failed to create inventory check", err)
			return errors.New("failed to create inventory check")
		}

		for _, detailReq := range req.Detail {

			// 获取库存
			stock := &model.StorehouseProduct{}
			err := tx.Where("uuid = ?", detailReq.StorehouseProductUuid).First(stock).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return errors.New("仓库中没有该商品")
				}
				ctx.Logger.Error("Failed to get stock", err)
				return errors.New("failed to get stock")
			}

			var differenceQuantity float64 = 0
			differenceBoxNumOp := "0"
			detailReq.DifferenceOp = "0" // 正常
			beforQuantity := stock.Quantity
			beforeBoxNum := stock.BoxNum
			var differenceBoxNum float64 = 0
			opDesc := "盘点，库存调整"

			if detailReq.Quantity > stock.Quantity {
				detailReq.DifferenceOp = "1"
				differenceQuantity = detailReq.Quantity - stock.Quantity //盘盈
				opDesc = "盘点，库存调整，盘盈, 增加库存"
			}

			if detailReq.Quantity < stock.Quantity {
				detailReq.DifferenceOp = "2"
				differenceQuantity = stock.Quantity - detailReq.Quantity //盘亏
				opDesc = "盘点，库存调整，盘亏, 减少库存"
			}

			if detailReq.BoxNum < stock.BoxNum {
				differenceBoxNum = stock.BoxNum - detailReq.BoxNum
				differenceBoxNumOp = "2"
				if detailReq.DifferenceOp == "0" {
					detailReq.DifferenceOp = "2"
				}
				opDesc += "，减少箱数"
			}

			if detailReq.BoxNum > stock.BoxNum {
				differenceBoxNumOp = "1"
				differenceBoxNum = detailReq.BoxNum - stock.BoxNum
				if detailReq.DifferenceOp == "0" {
					detailReq.DifferenceOp = "1"
				}
				opDesc += "，增加箱数"
			}

			detail := &model.StorehouseInventoryCheckDetail{
				Uuid:                  uuid.New().String(),
				StorehouseUuid:        req.StorehouseUuid,
				StorehouseProductUuid: detailReq.StorehouseProductUuid,
				CheckOrderNo:          check.CheckOrderNo,
				ProductUuid:           detailReq.ProductUuid,
				SkuUuid:               detailReq.SkuUuid,
				Quantity:              detailReq.Quantity,
				BoxNum:                detailReq.BoxNum,
				DifferenceOp:          detailReq.DifferenceOp,
				DifferenceBoxNumOp:    differenceBoxNumOp,
				CreatedAt:             nowStr,
				UpdatedAt:             nowStr,
				DifferenceQuantity:    differenceQuantity,
				DifferenceBoxNum:      differenceBoxNum,
			}

			if err := tx.Create(detail).Error; err != nil {
				ctx.Logger.Error("Failed to create inventory check detail", err)
				return errors.New("failed to create inventory check detail")
			}

			ctx.Logger.Info("detailReq.DifferenceOp:", detailReq.DifferenceOp)
			ctx.Logger.Info("detailReq.Quantity:", detailReq.Quantity)
			ctx.Logger.Info("detailReq.BoxNum: ", detailReq.BoxNum)
			ctx.Logger.Info("stock.Quantity: ", stock.Quantity)
			ctx.Logger.Info("stock.BoxNum: ", stock.BoxNum)

			if detailReq.DifferenceOp != "0" {
				stock.Quantity = detailReq.Quantity
				stock.BoxNum = detailReq.BoxNum
				stock.UpdatedAt = nowStr
				if err := tx.Where("uuid = ?", stock.Uuid).Updates(stock).Error; err != nil {
					ctx.Logger.Error("Failed to update stock", err)
					return err
				}
				// 创建库存记录
				stockopLog := &model.StorehouseProductOpLog{
					Uuid:                  uuid.New().String(),
					StorehouseProductUuid: stock.Uuid,
					StorehouseUuid:        check.StorehouseUuid,
					BeforeQuantity:        beforQuantity,
					Quantity:              stock.Quantity,
					OpQuantity:            differenceQuantity,
					BeforeBoxNum:          beforeBoxNum,
					BoxNum:                stock.BoxNum,
					OpBoxNum:              differenceBoxNum,
					OpType:                model.StorehouseProductOpLogOpTypeInventoryCheck,
					OpDesc:                opDesc,
					OpBy:                  userId,
					CreatedAt:             nowStr,
				}
				if err := tx.Create(stockopLog).Error; err != nil {
					ctx.Logger.Error("Failed to create stockop log", err)
					return errors.New("failed to create stockop log")
				}
			}

		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *StorehouseInventoryCheckService) GetInventoryCheck(ctx *app.Context, requuid string) (*model.StorehouseInventoryCheckRes, error) {
	check := &model.StorehouseInventoryCheck{}
	err := ctx.DB.Where("check_order_no = ?", requuid).First(check).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("inventory check not found")
		}
		ctx.Logger.Error("Failed to get inventory check by ID", err)
		return nil, errors.New("failed to get inventory check by ID")
	}

	storehouse, err := NewStorehouseService().GetStorehouseByUUID(ctx, check.StorehouseUuid)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, errors.New("failed to get storehouse by UUID")
	}

	user, err := NewUserService().GetUserByUUID(ctx, check.CheckBy)
	if err != nil {
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, errors.New("failed to get user by UUID")
	}

	checkRes := &model.StorehouseInventoryCheckRes{
		StorehouseInventoryCheck: *check,
		Storehouse:               *storehouse,
		CheckByUser:              *user,
	}

	return checkRes, nil
}

// 获取盘点清单
func (s *StorehouseInventoryCheckService) GetInventoryCheckDetail(ctx *app.Context, requuid string) ([]*model.StorehouseInventoryCheckDetailRes, error) {
	details := make([]*model.StorehouseInventoryCheckDetail, 0)
	err := ctx.DB.Where("check_order_no = ?", requuid).Find(&details).Error
	if err != nil {
		ctx.Logger.Error("Failed to get inventory check detail", err)
		return nil, errors.New("failed to get inventory check detail")
	}

	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	for _, detail := range details {
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

	res := make([]*model.StorehouseInventoryCheckDetailRes, 0)
	for _, detail := range details {
		detailRes := &model.StorehouseInventoryCheckDetailRes{
			StorehouseInventoryCheckDetail: *detail,
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

func (s *StorehouseInventoryCheckService) UpdateInventoryCheck(ctx *app.Context, check *model.StorehouseInventoryCheck) error {
	check.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(check).Error
	if err != nil {
		ctx.Logger.Error("Failed to update inventory check", err)
		return errors.New("failed to update inventory check")
	}
	return nil
}

func (s *StorehouseInventoryCheckService) DeleteInventoryCheck(ctx *app.Context, requuid string) error {

	ctx.DB.Transaction(func(tx *gorm.DB) error {

		err := tx.Where("check_order_no = ?", requuid).Delete(&model.StorehouseInventoryCheck{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete inventory check", err)
			return errors.New("failed to delete inventory check")
		}
		// 删除清单
		err = tx.Where("check_order_no = ?", requuid).Delete(&model.StorehouseInventoryCheckDetail{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete inventory check detail", err)
			return errors.New("failed to delete inventory check detail")
		}
		return nil
	})

	return nil
}

func (s *StorehouseInventoryCheckService) DeleteInventoryCheckDetail(ctx *app.Context, userId string, requuid string, checkOrderNo string) error {

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		// 先获取盘点明细详情
		var checkdetail model.StorehouseInventoryCheckDetail
		err := tx.Where("uuid = ?", requuid).First(&checkdetail).Error
		if err != nil {
			ctx.Logger.Error("Failed to get inventory check detail", err)
			return errors.New("failed to get inventory check detail")
		}

		if checkdetail.DifferenceOp == "0" && checkdetail.DifferenceBoxNumOp == "0" {
			// 直接删除
			err := tx.Where("uuid = ?", requuid).Delete(&model.StorehouseInventoryCheckDetail{}).Error
			if err != nil {
				ctx.Logger.Error("Failed to delete inventory check detail", err)
				return errors.New("failed to delete inventory check detail")
			}

			return nil
		}
		// 获取库存
		stock := &model.StorehouseProduct{}
		err = tx.Where("uuid = ?", checkdetail.StorehouseProductUuid).First(stock).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("仓库中没有该商品")
			}
			ctx.Logger.Error("Failed to get stock", err)
			return errors.New("failed to get stock")
		}

		beforeQuantify := stock.Quantity
		beforeBoxNum := stock.BoxNum
		if checkdetail.DifferenceOp == "1" {
			stock.Quantity = stock.Quantity - checkdetail.DifferenceQuantity
		}

		if checkdetail.DifferenceOp == "2" {
			stock.Quantity = stock.Quantity + checkdetail.DifferenceQuantity
		}

		if checkdetail.DifferenceBoxNumOp == "1" {
			stock.BoxNum = stock.BoxNum - checkdetail.DifferenceBoxNum
		}

		if checkdetail.DifferenceBoxNumOp == "2" {
			stock.BoxNum = stock.BoxNum + checkdetail.DifferenceBoxNum
		}

		stock.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		err = tx.Where("uuid = ?", stock.Uuid).Select("quantity", "box_num", "updated_at").Updates(stock).Error
		if err != nil {
			ctx.Logger.Error("Failed to update stock", err)
			return errors.New("failed to update stock")
		}

		// 创建库存记录log
		stockopLog := &model.StorehouseProductOpLog{
			Uuid:                  uuid.New().String(),
			StorehouseProductUuid: stock.Uuid,
			StorehouseUuid:        stock.StorehouseUuid,
			BeforeQuantity:        beforeQuantify,
			Quantity:              stock.Quantity,
			OpQuantity:            checkdetail.DifferenceQuantity,
			BeforeBoxNum:          beforeBoxNum,
			BoxNum:                stock.BoxNum,
			OpBoxNum:              checkdetail.DifferenceBoxNum,
			OpType:                model.StorehouseProductOpLogOpTypeInventoryCheck,
			OpDesc:                "盘点，库存调整，删除盘点记录",
			OpBy:                  userId,
			CreatedAt:             time.Now().Format("2006-01-02 15:04:05"),
		}

		err = tx.Create(stockopLog).Error
		if err != nil {
			ctx.Logger.Error("Failed to create stockop log", err)
			return errors.New("failed to create stockop log")
		}

		// 删除盘点记录
		err = tx.Where("uuid = ?", requuid).Delete(&model.StorehouseInventoryCheckDetail{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete inventory check detail", err)
			return errors.New("failed to delete inventory check detail")
		}

		return nil
	})
	if err != nil {
		ctx.Logger.Error("Failed to delete inventory check detail", err)
		return errors.New("failed to delete inventory check detail")
	}

	if checkOrderNo != "" {
		// 查看盘点记录是否全部删除
		var count int64
		err = ctx.DB.Model(&model.StorehouseInventoryCheckDetail{}).Where("check_order_no = ?", checkOrderNo).Count(&count).Error
		if err != nil {
			ctx.Logger.Error("Failed to get inventory check detail count", err)
			return errors.New("failed to get inventory check detail count")
		}
		if count == 0 {
			err = ctx.DB.Where("check_order_no = ?", checkOrderNo).Delete(&model.StorehouseInventoryCheck{}).Error
			if err != nil {
				ctx.Logger.Error("Failed to delete inventory check", err)
				return errors.New("failed to delete inventory check")
			}
		}
	}

	return nil
}

func (s *StorehouseInventoryCheckService) ListInventoryChecks(ctx *app.Context, param *model.ReqInventoryCheckQueryParam) (r *model.PagedResponse, err error) {
	var (
		checkList []*model.StorehouseInventoryCheck
		total     int64
	)

	db := ctx.DB.Model(&model.StorehouseInventoryCheck{})

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&checkList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	userUuids := make([]string, 0)
	storehouseUuids := make([]string, 0)
	for _, check := range checkList {
		storehouseUuids = append(storehouseUuids, check.StorehouseUuid)
		userUuids = append(userUuids, check.CheckBy)
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseInventoryCheckRes, 0)
	for _, check := range checkList {
		checkitem := &model.StorehouseInventoryCheckRes{
			StorehouseInventoryCheck: *check,
		}
		if storehouse, ok := storehouseMap[check.StorehouseUuid]; ok {
			checkitem.Storehouse = *storehouse
		}
		if user, ok := userMap[check.CheckBy]; ok {
			checkitem.CheckByUser = *user
		}

		res = append(res, checkitem)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}

// 根据盘点单号list获取盘点订单
func (s *StorehouseInventoryCheckService) GetInventoryChecksByOrderNos(ctx *app.Context, checkOrderNos []string) (map[string]*model.StorehouseInventoryCheck, error) {
	checkList := make([]*model.StorehouseInventoryCheck, 0)
	err := ctx.DB.Where("check_order_no IN ?", checkOrderNos).Find(&checkList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get inventory check by order nos", err)
		return nil, errors.New("failed to get inventory check by order nos")
	}

	checkMap := make(map[string]*model.StorehouseInventoryCheck)
	for _, check := range checkList {
		checkMap[check.CheckOrderNo] = check
	}

	return checkMap, nil
}

func (s *StorehouseInventoryCheckService) ListInventoryChecks2(ctx *app.Context, param *model.ReqInventoryCheckQueryParam) (r *model.PagedResponse, err error) {
	var (
		checkListDetail []*model.StorehouseInventoryCheckDetail
		total           int64
	)

	db := ctx.DB.Model(&model.StorehouseInventoryCheckDetail{})

	if param.CheckOrderNo != "" {
		db = db.Where("check_order_no = ?", param.CheckOrderNo)
	}

	if param.StorehouseUuid != "" {
		db = db.Where("storehouse_uuid = ?", param.StorehouseUuid)
	}

	if param.ProductUuid != "" {
		db = db.Where("product_uuid = ?", param.ProductUuid)
	}

	if err = db.Order("id DESC").Offset(param.GetOffset()).Limit(param.PageSize).Find(&checkListDetail).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	checkOrderNos := make([]string, 0)
	productUuids := make([]string, 0)
	skuUuids := make([]string, 0)
	storeProductUuids := make([]string, 0)
	for _, check := range checkListDetail {
		checkOrderNos = append(checkOrderNos, check.CheckOrderNo)
		productUuids = append(productUuids, check.ProductUuid)
		skuUuids = append(skuUuids, check.SkuUuid)
		storeProductUuids = append(storeProductUuids, check.StorehouseProductUuid)
	}

	storeProductMap, err := NewStorehouseProductService().GetProductListByUUIDs(ctx, storeProductUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get store product list by UUIDs", err)
		return
	}

	checkMap, err := s.GetInventoryChecksByOrderNos(ctx, checkOrderNos)
	if err != nil {
		ctx.Logger.Error("Failed to get inventory check by order nos", err)
		return
	}

	userUuids := make([]string, 0)
	storehouseUuids := make([]string, 0)

	for _, check := range checkMap {
		storehouseUuids = append(storehouseUuids, check.StorehouseUuid)
		userUuids = append(userUuids, check.CheckBy)
	}

	storehouseMap, err := NewStorehouseService().GetStorehousesByUUIDs(ctx, storehouseUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get storehouse list by UUIDs", err)
		return
	}

	userMap, err := NewUserService().GetUsersByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get user list by UUIDs", err)
		return
	}

	productMap, err := NewProductService().GetProductListByUUIDs(ctx, productUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get product list by UUIDs", err)
		return
	}

	skuMap, err := NewSkuService().GetSkuListByUUIDs(ctx, skuUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get sku list by UUIDs", err)
		return
	}

	res := make([]*model.StorehouseInventoryCheckRes, 0)
	for _, check := range checkListDetail {

		detailResItem := &model.StorehouseInventoryCheckDetailRes{
			StorehouseInventoryCheckDetail: *check,
		}

		if product, ok := productMap[check.ProductUuid]; ok {
			detailResItem.Product = *product
		}

		if sku, ok := skuMap[check.SkuUuid]; ok {
			detailResItem.Sku = *sku
		}

		checkitem := &model.StorehouseInventoryCheckRes{
			StorehouseInventoryCheckDetailRes: *detailResItem,
		}

		if check, ok := checkMap[check.CheckOrderNo]; ok {
			checkitem.StorehouseInventoryCheck = *check
		}

		if storehouse, ok := storehouseMap[checkitem.StorehouseInventoryCheck.StorehouseUuid]; ok {
			checkitem.Storehouse = *storehouse
		}

		if user, ok := userMap[checkitem.StorehouseInventoryCheck.CheckBy]; ok {
			checkitem.CheckByUser = *user
		}

		if storeProduct, ok := storeProductMap[checkitem.StorehouseInventoryCheckDetail.StorehouseProductUuid]; ok {
			checkitem.StorehouseProduct = *storeProduct
		}

		res = append(res, checkitem)
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     res,
	}
	return
}
