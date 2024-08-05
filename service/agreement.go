package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

	"gorm.io/gorm"
)

type AgreementService struct{}

func NewAgreementService() *AgreementService {
	return &AgreementService{}
}

func (s *AgreementService) CreateAgreement(ctx *app.Context, userid string, agreement *model.Agreement) error {

	agreement.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	agreement.UpdatedAt = agreement.CreatedAt
	agreement.Creater = userid
	agreement.Uuid = utils.GenerateOrderID()

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(agreement).Error
		if err != nil {
			ctx.Logger.Error("Failed to create agreement", err)
			return errors.New("failed to create agreement")
		}

		if agreement.Type == model.AgreementTypeSales && agreement.OrderNo != "" {
			// 更新销售订单的合同状态
			err = tx.Model(&model.SalesOrder{}).Where("order_no = ?", agreement.OrderNo).Updates(map[string]interface{}{
				"agreement_uuid": agreement.Uuid,
				"updated_at":     time.Now().Format("2006-01-02 15:04:05"),
			}).Error

			if err != nil {
				ctx.Logger.Error("Failed to update sales order agreement status", err)
				return errors.New("failed to update sales order agreement status")
			}
		}
		return nil
	})
	if err != nil {
		ctx.Logger.Error("Failed to create agreement", err)
		return err
	}

	return nil
}

func (s *AgreementService) GetAgreement(ctx *app.Context, uuid string) (*model.Agreement, error) {
	agreement := &model.Agreement{}
	err := ctx.DB.Where("uuid = ?", uuid).First(agreement).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return agreement, nil
		}
		ctx.Logger.Error("Failed to get agreement by UUID", err)
		return nil, errors.New("failed to get agreement by UUID")
	}
	return agreement, nil
}

func (s *AgreementService) UpdateAgreement(ctx *app.Context, agreement *model.Agreement) error {
	agreement.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", agreement.Uuid).Updates(agreement).Error
	if err != nil {
		ctx.Logger.Error("Failed to update agreement", err)
		return errors.New("failed to update agreement")
	}
	return nil
}

func (s *AgreementService) DeleteAgreement(ctx *app.Context, uuid string) error {

	ctx.DB.Transaction(func(tx *gorm.DB) error {

		// 先查询合同
		agreement := &model.Agreement{}
		err := tx.Where("uuid = ?", uuid).First(agreement).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil
			}
			ctx.Logger.Error("Failed to get agreement by UUID", err)
			return errors.New("failed to get agreement by UUID")
		}

		err = tx.Where("uuid = ?", uuid).Delete(&model.Agreement{}).Error
		if err != nil {
			ctx.Logger.Error("Failed to delete agreement", err)
			return errors.New("failed to delete agreement")
		}

		if agreement.Type == model.AgreementTypeSales && agreement.OrderNo != "" {
			// 更新销售订单的合同状态
			err = tx.Model(&model.SalesOrder{}).Where("order_no = ?", agreement.OrderNo).Updates(map[string]interface{}{
				"agreement_uuid": "",
				"updated_at":     time.Now().Format("2006-01-02 15:04:05"),
			}).Error

			if err != nil {
				ctx.Logger.Error("Failed to update sales order agreement status", err)
				return errors.New("failed to update sales order agreement status")
			}
		}

		return nil
	})
	return nil
}

func (s *AgreementService) ListAgreements(ctx *app.Context, param *model.ReqAgreementQueryParam) (*model.PagedResponse, error) {
	var (
		agreements []*model.Agreement
		total      int64
	)

	db := ctx.DB.Model(&model.Agreement{})

	if param.Type != "" {
		db = db.Where("type = ?", param.Type)
	}

	if err := db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&agreements).Error; err != nil {
		return nil, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	return &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     agreements,
	}, nil
}
