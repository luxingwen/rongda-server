package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentBillService struct {
}

func NewPaymentBillService() *PaymentBillService {
	return &PaymentBillService{}
}

func (s *PaymentBillService) CreatePaymentBill(ctx *app.Context, paymentBill *model.PaymentBill) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	paymentBill.Uuid = uuid.New().String()
	paymentBill.CreatedAt = now
	paymentBill.UpdatedAt = now

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(paymentBill).Error
		if err != nil {
			ctx.Logger.Error("Failed to create payment bill", err)
			tx.Rollback()
			return errors.New("failed to create payment bill")
		}
		return nil
	})

	if err != nil {
		ctx.Logger.Error("Failed to create payment bill", err)
		return errors.New("failed to create payment bill")
	}
	return nil
}

func (s *PaymentBillService) GetPaymentBillByUUID(ctx *app.Context, uuid string) (*model.PaymentBill, error) {
	paymentBill := &model.PaymentBill{}
	err := ctx.DB.Where("uuid = ?", uuid).First(paymentBill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("payment bill not found")
		}
		ctx.Logger.Error("Failed to get payment bill by UUID", err)
		return nil, errors.New("failed to get payment bill by UUID")
	}
	return paymentBill, nil
}

func (s *PaymentBillService) UpdatePaymentBill(ctx *app.Context, paymentBill *model.PaymentBill) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	paymentBill.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", paymentBill.Uuid).Updates(paymentBill).Error
	if err != nil {
		ctx.Logger.Error("Failed to update payment bill", err)
		return errors.New("failed to update payment bill")
	}

	return nil
}

func (s *PaymentBillService) DeletePaymentBill(ctx *app.Context, uuid string) error {

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.PaymentBill{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete payment bill", err)
			tx.Rollback()
			return errors.New("failed to delete payment bill")
		}
		// 更新步骤ref_id
		err = tx.Model(&model.Step{}).Where("ref_id = ?", uuid).Update("ref_id", "").Error
		if err != nil {
			ctx.Logger.Error("Failed to update payment step ref_id", err)
			tx.Rollback()
			return errors.New("failed to update payment step ref_id")
		}
		return nil
	})

	return err
}

// GetPaymentBillList retrieves a list of payment bills based on query parameters
func (s *PaymentBillService) GetPaymentBillList(ctx *app.Context, params *model.ReqPaymentBillQueryParam) (*model.PagedResponse, error) {
	var (
		paymentBills []*model.PaymentBill
		total        int64
	)

	db := ctx.DB.Model(&model.PaymentBill{})

	if params.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+params.OrderNo+"%")
	}
	if params.AgreementNo != "" {
		db = db.Where("agreement_no LIKE ?", "%"+params.AgreementNo+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get payment bill count", err)
		return nil, errors.New("failed to get payment bill count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&paymentBills).Error
	if err != nil {
		ctx.Logger.Error("Failed to get payment bill list", err)
		return nil, errors.New("failed to get payment bill list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  paymentBills,
	}, nil
}

// UpdatePaymentBillStatus
func (s *PaymentBillService) UpdatePaymentBillStatus(ctx *app.Context, param *model.ReqUpdatePaymentBillStatusParam) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Model(&model.PaymentBill{}).Where("uuid = ?", param.Uuid).Updates(map[string]interface{}{
		"status":     param.Status,
		"updated_at": now,
	}).Error
	if err != nil {
		ctx.Logger.Error("Failed to update payment bill status", err)
		return errors.New("failed to update payment bill status")
	}

	return nil
}

// 根据uuid 列表修改订单状态为已支付待确认
func (s *PaymentBillService) UpdatePaymentBillStatusPaidPendingConfirm(ctx *app.Context, params *model.ReqPaymentBillOrderStatusPaidComfirm) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	// 先获取账单
	var paybillMent model.PaymentBill
	err := ctx.DB.Where("uuid = ?", params.Uuid).First(&paybillMent).Error
	if err != nil {
		ctx.Logger.Error("Failed to get payment bill by UUID", err)
		return errors.New("failed to get payment bill by UUID")
	}

	// 如果应付金额和实付金额不一致
	if params.PaymentAmount < paybillMent.Amount {

		// 未付款
		paybillMent.UnpaidAmount = paybillMent.Amount - params.PaymentAmount
	}

	// 更新账单状态
	err = ctx.DB.Model(&model.PaymentBill{}).Where("uuid = ?", params.Uuid).Updates(map[string]interface{}{
		"status":         model.PaymentBillStatusPaidPendingConfirm,
		"updated_at":     now,
		"payment_amount": params.PaymentAmount,
		"unpaid_amount":  paybillMent.UnpaidAmount,
	}).Error
	if err != nil {
		ctx.Logger.Error("Failed to update payment bill status", err)
		return errors.New("failed to update payment bill status")
	}

	return nil
}