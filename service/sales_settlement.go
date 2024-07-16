package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesSettlementService struct {
}

func NewSalesSettlementService() *SalesSettlementService {
	return &SalesSettlementService{}
}

func (s *SalesSettlementService) CreateSalesSettlement(ctx *app.Context, userId string, req *model.SalesSettlementReq) error {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	settlement := &model.SalesSettlement{
		Uuid:               uuid.New().String(),
		OrderUuid:          req.OrderUuid,
		PaymentMethod:      req.PaymentMethod,
		PaymentDate:        req.PaymentDate,
		Amount:             req.Amount,
		Remark:             req.Remark,
		PaymentVoucher:     req.PaymentVoucher,
		FinanceStaff:       req.FinanceStaff,
		FinanceAuditDate:   req.FinanceAuditDate,
		FinanceAuditStatus: req.FinanceAuditStatus,
		FinanceAuditRemark: req.FinanceAuditRemark,
		Status:             req.Status,
		Creater:            userId,
		CreatedAt:          nowStr,
		UpdatedAt:          nowStr,
	}

	if err := ctx.DB.Create(settlement).Error; err != nil {
		ctx.Logger.Error("Failed to create sales settlement", err)
		return errors.New("failed to create sales settlement")
	}
	return nil
}

func (s *SalesSettlementService) GetSalesSettlement(ctx *app.Context, uuid string) (*model.SalesSettlement, error) {
	settlement := &model.SalesSettlement{}
	err := ctx.DB.Where("uuid = ?", uuid).First(settlement).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("sales settlement not found")
		}
		ctx.Logger.Error("Failed to get sales settlement by UUID", err)
		return nil, errors.New("failed to get sales settlement by UUID")
	}
	return settlement, nil
}

func (s *SalesSettlementService) UpdateSalesSettlement(ctx *app.Context, settlement *model.SalesSettlement) error {
	settlement.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(settlement).Error
	if err != nil {
		ctx.Logger.Error("Failed to update sales settlement", err)
		return errors.New("failed to update sales settlement")
	}
	return nil
}

func (s *SalesSettlementService) DeleteSalesSettlement(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.SalesSettlement{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete sales settlement", err)
		return errors.New("failed to delete sales settlement")
	}
	return nil
}

func (s *SalesSettlementService) ListSalesSettlements(ctx *app.Context, param *model.ReqSalesSettlementQueryParam) (r *model.PagedResponse, err error) {
	var (
		settlementList []*model.SalesSettlement
		total          int64
	)

	db := ctx.DB.Model(&model.SalesSettlement{})

	if param.OrderUuid != "" {
		db = db.Where("order_uuid = ?", param.OrderUuid)
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&settlementList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     settlementList,
	}
	return
}
