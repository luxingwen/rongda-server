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
	agreement.Status = model.AgreementStatusUnSigned

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {

		// 获取销售订单
		salesOrder := &model.SalesOrder{}
		err := tx.Where("order_no = ?", agreement.OrderNo).First(salesOrder).Error
		if err != nil {
			ctx.Logger.Error("Failed to get sales order", err)
			tx.Rollback()
			return errors.New("failed to get sales order")
		}

		agreement.TeamUuid = salesOrder.CustomerUuid

		err = tx.Create(agreement).Error
		if err != nil {
			ctx.Logger.Error("Failed to create agreement", err)
			tx.Rollback()
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
				tx.Rollback()
				return errors.New("failed to update sales order agreement status")
			}

			// 获取步骤链
			stepChain := &model.StepChain{}
			err = tx.Where("ref_id = ? AND ref_type = ?", agreement.OrderNo, model.StepChainRefTypeSalesOrder).First(stepChain).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step chain", err)
				tx.Rollback()
				return errors.New("failed to get step chain")
			}

			// 获取步骤
			step := &model.Step{}
			err = tx.Where("chain_id = ? and title = ?", stepChain.Uuid, "创建合同").First(step).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step", err)
				tx.Rollback()
				return errors.New("failed to get step")
			}

			agreement.RefId = step.Uuid
			agreement.RefType = model.AgreementTypeSales

			// 更新步骤

			err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "创建合同").Updates(map[string]interface{}{
				"status":     model.StepStatusFinish,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				"ref_id":     agreement.Uuid,
				"ref_type":   model.AgreementTypeSales,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update step", err)
				tx.Rollback()
				return errors.New("failed to update step")
			}

			err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "签署合同").Updates(map[string]interface{}{
				"status":     model.StepStatusWait,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				"ref_type":   model.AgreementTypeSales,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update step", err)
				tx.Rollback()
				return errors.New("failed to update step")
			}

		}

		if agreement.Type == model.AgreementTypeSalesDeposit && agreement.OrderNo != "" {
			// 获取步骤链
			stepChain := &model.StepChain{}
			err = tx.Where("ref_id = ? AND ref_type = ?", agreement.OrderNo, model.StepChainRefTypeSalesOrder).First(stepChain).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step chain", err)
				tx.Rollback()
				return errors.New("failed to get step chain")
			}

			// 获取步骤
			step := &model.Step{}
			err = tx.Where("chain_id = ? and title = ?", stepChain.Uuid, "创建定金合同").First(step).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step", err)
				tx.Rollback()
				return errors.New("failed to get step")
			}

			agreement.RefId = step.Uuid
			agreement.RefType = model.AgreementTypeSalesDeposit

			// 更新步骤

			err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "创建定金合同").Updates(map[string]interface{}{
				"status":     model.StepStatusFinish,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				"ref_id":     agreement.Uuid,
				"ref_type":   model.AgreementTypeSalesDeposit,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update step", err)
				tx.Rollback()
				return errors.New("failed to update step")
			}

			err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "签署定金合同").Updates(map[string]interface{}{
				"status":     model.StepStatusWait,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				"ref_type":   model.AgreementTypeSalesDeposit,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update step", err)
				tx.Rollback()
				return errors.New("failed to update step")
			}
		}

		if agreement.Type == model.AgreementTypeSalesFinalPayment && agreement.OrderNo != "" {
			// 获取步骤链
			stepChain := &model.StepChain{}
			err = tx.Where("ref_id = ? AND ref_type = ?", agreement.OrderNo, model.StepChainRefTypeSalesOrder).First(stepChain).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step chain", err)
				tx.Rollback()
				return errors.New("failed to get step chain")
			}

			// 获取步骤
			step := &model.Step{}
			err = tx.Where("chain_id = ? and title = ?", stepChain.Uuid, "创建尾款合同").First(step).Error
			if err != nil {
				ctx.Logger.Error("Failed to get step", err)
				tx.Rollback()
				return errors.New("failed to get step")
			}

			agreement.RefId = step.Uuid
			agreement.RefType = model.AgreementTypeSalesFinalPayment

			// 更新步骤

			err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "创建尾款合同").Updates(map[string]interface{}{
				"status":     model.StepStatusFinish,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				"ref_id":     agreement.Uuid,
				"ref_type":   model.AgreementTypeSalesFinalPayment,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update step", err)
				tx.Rollback()
				return errors.New("failed to update step")
			}

			err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "签署尾款合同").Updates(map[string]interface{}{
				"status":     model.StepStatusWait,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				"ref_type":   model.AgreementTypeSalesFinalPayment,
			}).Error
			if err != nil {
				ctx.Logger.Error("Failed to update step", err)
				tx.Rollback()
				return errors.New("failed to update step")
			}
		}

		err = tx.Model(&model.Agreement{}).Where("uuid = ?", agreement.Uuid).Updates(map[string]interface{}{
			"ref_id":   agreement.RefId,
			"ref_type": agreement.RefType,
		}).Error

		if err != nil {
			ctx.Logger.Error("Failed to update agreement ref_id", err)
			tx.Rollback()
			return errors.New("failed to update agreement ref_id")
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
			tx.Rollback()
			ctx.Logger.Error("Failed to get agreement by UUID", err)
			return errors.New("failed to get agreement by UUID")
		}

		err = tx.Where("uuid = ?", uuid).Delete(&model.Agreement{}).Error
		if err != nil {
			tx.Rollback()
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
				tx.Rollback()
				ctx.Logger.Error("Failed to update sales order agreement status", err)
				return errors.New("failed to update sales order agreement status")
			}

			// // 获取步骤链
			// stepChain := &model.StepChain{}
			// err = tx.Where("ref_id = ? AND ref_type = ?", agreement.OrderNo, model.StepChainRefTypeSalesOrder).First(stepChain).Error
			// if err != nil {
			// 	ctx.Logger.Error("Failed to get step chain", err)
			// 	return errors.New("failed to get step chain")
			// }

			// // 更新步骤

			// err = tx.Model(&model.Step{}).Where("chain_id = ? and title = ?", stepChain.Uuid, "创建合同").Updates(map[string]interface{}{
			// 	"status":     model.StepStatusWait,
			// 	"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			// 	"ref_id":     "",
			// 	"ref_type":   model.AgreementTypeSalesDeposit,
			// }).Error
			// if err != nil {
			// 	ctx.Logger.Error("Failed to update step", err)
			// 	return errors.New("failed to update step")
			// }

		}

		// 更新步骤
		err = tx.Model(&model.Step{}).Where("ref_id = ?", agreement.Uuid).Updates(map[string]interface{}{
			"status":     model.StepStatusWait,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			"ref_id":     "",
		}).Error
		if err != nil {
			tx.Rollback()
			ctx.Logger.Error("Failed to update step", err)
			return errors.New("failed to update step")
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

	if param.TeamUuid != "" {
		db = db.Where("team_uuid = ?", param.TeamUuid)
	}

	if param.Status != "" {
		db = db.Where("status = ?", param.Status)
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

// GetAgreementByOrder
func (s *AgreementService) GetAgreementByOrder(ctx *app.Context, params *model.ReqOrderAgreementQueryParam) (*model.Agreement, error) {
	agreement := &model.Agreement{}
	err := ctx.DB.Where("order_no = ? AND type = ?", params.OrderNo, params.Type).First(agreement).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return agreement, nil
		}
		ctx.Logger.Error("Failed to get agreement by order no", err)
		return nil, errors.New("failed to get agreement by order no")
	}
	return agreement, nil
}

// UpdateAgreementSign
func (s *AgreementService) UpdateAgreementSign(ctx *app.Context, agreement *model.Agreement) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Model(&model.Agreement{}).Where("uuid = ?", agreement.Uuid).Updates(map[string]interface{}{
		"status":          model.AgreementStatusSigned,
		"signature_image": agreement.SignatureImage,
		"signature_user":  agreement.SignatureUser,
		"signature_time":  now,
		"updated_at":      now,
	}).Error
	if err != nil {
		ctx.Logger.Error("Failed to update agreement sign", err)
		return errors.New("failed to update agreement sign")
	}

	return nil
}
