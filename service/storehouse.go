package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StorehouseService struct {
}

func NewStorehouseService() *StorehouseService {
	return &StorehouseService{}
}

func (s *StorehouseService) CreateStorehouse(ctx *app.Context, storehouse *model.Storehouse) error {
	storehouse.Uuid = uuid.New().String()
	storehouse.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	storehouse.UpdatedAt = storehouse.CreatedAt

	err := ctx.DB.Create(storehouse).Error
	if err != nil {
		ctx.Logger.Error("Failed to create storehouse", err)
		return errors.New("failed to create storehouse")
	}
	return nil
}

func (s *StorehouseService) GetStorehouseByUUID(ctx *app.Context, uuid string) (*model.Storehouse, error) {
	storehouse := &model.Storehouse{}
	err := ctx.DB.Where("uuid = ?", uuid).First(storehouse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("storehouse not found")
		}
		ctx.Logger.Error("Failed to get storehouse by UUID", err)
		return nil, errors.New("failed to get storehouse by UUID")
	}
	return storehouse, nil
}

func (s *StorehouseService) UpdateStorehouse(ctx *app.Context, storehouse *model.Storehouse) error {
	storehouse.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", storehouse.Uuid).Updates(storehouse).Error
	if err != nil {
		ctx.Logger.Error("Failed to update storehouse", err)
		return errors.New("failed to update storehouse")
	}
	return nil
}

func (s *StorehouseService) DeleteStorehouse(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Storehouse{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete storehouse", err)
		return errors.New("failed to delete storehouse")
	}
	return nil
}

// 获取仓库列表
func (s *StorehouseService) GetStorehouseList(ctx *app.Context, param *model.ReqStorehouseQueryParam) (r *model.PagedResponse, err error) {
	var (
		storehouseList []*model.Storehouse
		total          int64
	)

	db := ctx.DB.Model(&model.Storehouse{})

	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&storehouseList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     storehouseList,
	}
	return
}

// 获取所有可用仓库
func (s *StorehouseService) GetAvailableStorehouses(ctx *app.Context) (storehouseList []*model.Storehouse, err error) {
	err = ctx.DB.Where("status = ?", model.StorehouseStatusEnabled).Find(&storehouseList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get available storehouses", err)
		return nil, errors.New("failed to get available storehouses")
	}
	return
}

// 根据uuids获取仓库列表
func (s *StorehouseService) GetStorehousesByUUIDs(ctx *app.Context, uuids []string) (r map[string]*model.Storehouse, err error) {
	storehouseList := make([]*model.Storehouse, 0)
	r = make(map[string]*model.Storehouse)
	err = ctx.DB.Where("uuid in (?)", uuids).Find(&storehouseList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get storehouses by UUIDs", err)
		return nil, errors.New("failed to get storehouses by UUIDs")
	}
	for _, storehouse := range storehouseList {
		r[storehouse.Uuid] = storehouse
	}
	return
}
