package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PurchaseBillService struct {
}

func NewPurchaseBillService() *PurchaseBillService {
	return &PurchaseBillService{}
}

func (s *PurchaseBillService) CreatePurchaseBill(ctx *app.Context, userId string, req *model.PurchaseBillReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	purchaseBill := &model.PurchaseBill{
		Uuid:               uuid.New().String(),
		PurchaseOrderNo:    req.PurchaseOrderNo,
		StockInOrderNo:     req.StockInOrderNo,
		SupplierUuid:       req.SupplierUuid,
		BankAccount:        req.BankAccount,
		BankName:           req.BankName,
		BankAccountName:    req.BankAccountName,
		Amount:             req.Amount,
		PaymentDate:        req.PaymentDate,
		PaymentMethod:      req.PaymentMethod,
		Remark:             req.Remark,
		FinanceStaff:       userId,
		FinanceAuditDate:   nowStr,
		FinanceAuditStatus: 1, // Default to '待审核'
		Status:             1, // Default to '待付款'
		CreatedAt:          nowStr,
		UpdatedAt:          nowStr,
	}

	err := ctx.DB.Create(purchaseBill).Error
	if err != nil {
		ctx.Logger.Error("Failed to create purchase bill", err)
		return errors.New("failed to create purchase bill")
	}

	return nil
}

func (s *PurchaseBillService) GetPurchaseBill(ctx *app.Context, uuid string) (*model.PurchaseBill, error) {
	purchaseBill := &model.PurchaseBill{}
	err := ctx.DB.Where("uuid = ?", uuid).First(purchaseBill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("purchase bill not found")
		}
		ctx.Logger.Error("Failed to get purchase bill by ID", err)
		return nil, errors.New("failed to get purchase bill by ID")
	}
	return purchaseBill, nil
}

func (s *PurchaseBillService) UpdatePurchaseBill(ctx *app.Context, purchaseBill *model.PurchaseBill) error {
	purchaseBill.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(purchaseBill).Error
	if err != nil {
		ctx.Logger.Error("Failed to update purchase bill", err)
		return errors.New("failed to update purchase bill")
	}
	return nil
}

func (s *PurchaseBillService) DeletePurchaseBill(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.PurchaseBill{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete purchase bill", err)
		return errors.New("failed to delete purchase bill")
	}
	return nil
}

func (s *PurchaseBillService) ListPurchaseBills(ctx *app.Context, param *model.ReqPurchaseBillQueryParam) (r *model.PagedResponse, err error) {
	var (
		billList []*model.PurchaseBill
		total    int64
	)

	db := ctx.DB.Model(&model.PurchaseBill{})

	if param.SupplierUuid != "" {
		db = db.Where("supplier_uuid = ?", param.SupplierUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&billList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     billList,
	}
	return
}
