package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

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

	err := ctx.DB.Create(agreement).Error
	if err != nil {
		ctx.Logger.Error("Failed to create agreement", err)
		return errors.New("failed to create agreement")
	}
	return nil
}

func (s *AgreementService) GetAgreement(ctx *app.Context, uuid string) (*model.Agreement, error) {
	agreement := &model.Agreement{}
	err := ctx.DB.Where("uuid = ?", uuid).First(agreement).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("agreement not found")
		}
		ctx.Logger.Error("Failed to get agreement by UUID", err)
		return nil, errors.New("failed to get agreement by UUID")
	}
	return agreement, nil
}

func (s *AgreementService) UpdateAgreement(ctx *app.Context, agreement *model.Agreement) error {
	agreement.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Save(agreement).Error
	if err != nil {
		ctx.Logger.Error("Failed to update agreement", err)
		return errors.New("failed to update agreement")
	}
	return nil
}

func (s *AgreementService) DeleteAgreement(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Agreement{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete agreement", err)
		return errors.New("failed to delete agreement")
	}
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
