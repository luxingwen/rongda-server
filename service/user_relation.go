package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRelationService struct {
}

func NewUserRelationService() *UserRelationService {
	return &UserRelationService{}
}

func (s *UserRelationService) CreateUserRelation(ctx *app.Context, relation *model.UserRelation) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	relation.CreatedAt = now
	relation.UpdatedAt = now
	relation.Uuid = uuid.New().String()

	// 先查询是否已经存在关系
	relationExist := &model.UserRelation{}
	err := ctx.DB.Where("from_module = ? AND   user_uuid = ? AND  ref_module = ? AND ref_uuid = ?", relation.FromModule, relation.UserUuid, relation.RefModule, relation.RefUuid).First(relationExist).Error

	if err == nil && relationExist.ID > 0 {
		return errors.New("user relation already exists")
	}
	err = ctx.DB.Create(relation).Error
	if err != nil {
		ctx.Logger.Error("Failed to create user relation", err)
		return errors.New("failed to create user relation")
	}
	return nil
}

func (s *UserRelationService) GetUserRelationByUUID(ctx *app.Context, uuid string) (*model.UserRelation, error) {
	relation := &model.UserRelation{}
	err := ctx.DB.Where("uuid = ?", uuid).First(relation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user relation not found")
		}
		ctx.Logger.Error("Failed to get user relation by UUID", err)
		return nil, errors.New("failed to get user relation by UUID")
	}
	return relation, nil
}

// 根据用户id获取用户关系列表
func (s *UserRelationService) GetUserRelationByUserUUID(ctx *app.Context, userUuid string) ([]model.UserRelation, error) {
	var relations []model.UserRelation
	err := ctx.DB.Where("user_uuid = ?", userUuid).Find(&relations).Error
	if err != nil {
		ctx.Logger.Error("Failed to get user relation by user uuid", err)
		return nil, errors.New("failed to get user relation by user uuid")
	}
	return relations, nil
}

func (s *UserRelationService) UpdateUserRelation(ctx *app.Context, relation *model.UserRelation) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	relation.UpdatedAt = now

	err := ctx.DB.Where("uuid = ?", relation.UserUuid).Updates(relation).Error
	if err != nil {
		ctx.Logger.Error("Failed to update user relation", err)
		return errors.New("failed to update user relation")
	}

	return nil
}

func (s *UserRelationService) DeleteUserRelation(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.UserRelation{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete user relation", err)
		return errors.New("failed to delete user relation")
	}

	return nil
}
