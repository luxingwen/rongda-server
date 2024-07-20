package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SysBankInfoService struct {
}

func NewSysBankInfoService() *SysBankInfoService {
	return &SysBankInfoService{}
}

func (s *SysBankInfoService) CreateSysBankInfo(ctx *app.Context, bankInfo *model.SysBankInfo) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	bankInfo.CreatedAt = now
	bankInfo.UpdatedAt = now
	bankInfo.Uuid = uuid.New().String()

	err := ctx.DB.Create(bankInfo).Error
	if err != nil {
		ctx.Logger.Error("Failed to create bank info", err)
		return errors.New("failed to create bank info")
	}
	return nil
}

func (s *SysBankInfoService) GetSysBankInfoByUUID(ctx *app.Context, uuid string) (*model.SysBankInfo, error) {
	bankInfo := &model.SysBankInfo{}
	err := ctx.DB.Where("uuid = ?", uuid).First(bankInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("bank info not found")
		}
		ctx.Logger.Error("Failed to get bank info by UUID", err)
		return nil, errors.New("failed to get bank info by UUID")
	}
	return bankInfo, nil
}

func (s *SysBankInfoService) UpdateSysBankInfo(ctx *app.Context, bankInfo *model.SysBankInfo) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	bankInfo.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", bankInfo.Uuid).Updates(bankInfo).Error
	if err != nil {
		ctx.Logger.Error("Failed to update bank info", err)
		return errors.New("failed to update bank info")
	}

	return nil
}

func (s *SysBankInfoService) DeleteSysBankInfo(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.SysBankInfo{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete bank info", err)
		return errors.New("failed to delete bank info")
	}

	return nil
}

// GetSysBankInfoList retrieves a list of bank infos based on query parameters
func (s *SysBankInfoService) GetSysBankInfoList(ctx *app.Context, params *model.ReqSysBankInfoQueryParam) (*model.PagedResponse, error) {
	var (
		bankInfos []*model.SysBankInfo
		total     int64
	)

	db := ctx.DB.Model(&model.SysBankInfo{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get bank info count", err)
		return nil, errors.New("failed to get bank info count")
	}

	err = db.Find(&bankInfos).Error
	if err != nil {
		ctx.Logger.Error("Failed to get bank info list", err)
		return nil, errors.New("failed to get bank info list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  bankInfos,
	}, nil
}

// GetAvailableSysBankInfoList retrieves a list of available bank infos
func (s *SysBankInfoService) GetAvailableSysBankInfoList(ctx *app.Context) ([]*model.SysBankInfo, error) {
	var bankInfos []*model.SysBankInfo
	err := ctx.DB.Find(&bankInfos).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available bank info list", err)
		return nil, errors.New("failed to get available bank info list")
	}
	return bankInfos, nil
}
