package service

import (
	"errors"
	"sgin/model"
	"sgin/pkg/app"
)

type StepService struct {
}

func NewStepService() *StepService {
	return &StepService{}
}

// 根据reftype 和 refid 获取步骤列表
func (s *StepService) GetStepListByRefTypeAndRefID(ctx *app.Context, reftype string, refid string) ([]*model.Step, error) {
	// 先获取步骤链
	stepChain := &model.StepChain{}
	err := ctx.DB.Where("ref_type = ? AND ref_id = ?", reftype, refid).First(stepChain).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		ctx.Logger.Error("Failed to get step chain", err)
		return nil, errors.New("failed to get step chain")
	}

	// 获取步骤列表
	stepList := []*model.Step{}
	err = ctx.DB.Where("chain_id = ?", stepChain.Uuid).Find(&stepList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get step list", err)
		return nil, errors.New("failed to get step list")
	}
	return stepList, nil
}
