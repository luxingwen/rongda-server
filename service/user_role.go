package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"gorm.io/gorm"
)

type UserRoleService struct {
}

func NewUserRoleService() *UserRoleService {
	return &UserRoleService{}
}

func (s *UserRoleService) CreateUserRole(ctx *app.Context, userRole *model.ReqUserRole) error {

	for _, roleuuid := range userRole.RoleUUIDs {
		ruserRole, err := s.GetUserRoleByUUID(ctx, userRole.UserUUID, roleuuid)
		if ruserRole != nil {
			continue
		}
		if err != nil && err.Error() != "user role not found" {
			ctx.Logger.Error("Failed to get user role by UUID", err)
			return errors.New("failed to get user role by UUID")
		}
		userRole := &model.UserRole{
			UserUUID: userRole.UserUUID,
			RoleUUID: roleuuid,
		}
		err = ctx.DB.Create(userRole).Error
		if err != nil {
			ctx.Logger.Error("Failed to create user role", err)
			return errors.New("failed to create user role")
		}
	}
	return nil
}

func (s *UserRoleService) GetUserRoleByUUID(ctx *app.Context, userUuid, roleUuid string) (*model.UserRole, error) {
	userRole := &model.UserRole{}
	err := ctx.DB.Where("user_uuid = ? and role_uuid = ?", userUuid, roleUuid).First(userRole).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user role not found")
		}
		ctx.Logger.Error("Failed to get user role by UUID", err)
		return nil, errors.New("failed to get user role by UUID")
	}
	return userRole, nil
}

func (s *UserRoleService) UpdateUserRole(ctx *app.Context, userRole *model.UserRole) error {
	userRole.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", userRole.UUID).Updates(userRole).Error
	if err != nil {
		ctx.Logger.Error("Failed to update user role", err)
		return errors.New("failed to update user role")
	}

	return nil
}

func (s *UserRoleService) DeleteUserRole(ctx *app.Context, userRole *model.ReqUserRole) error {
	err := ctx.DB.Where("user_uuid = ?", userRole.UserUUID).Where("role_uuid IN ?", userRole.RoleUUIDs).Delete(&model.UserRole{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete user role", err)
		return errors.New("failed to delete user role")
	}

	return nil
}

// 获取用户的角色信息
func (s *UserRoleService) GetUserRoleByUserID(ctx *app.Context, userID string) ([]*model.Role, error) {
	var userRole []*model.UserRole
	err := ctx.DB.Where("user_uuid = ?", userID).Find(&userRole).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user role not found")
		}
		ctx.Logger.Error("Failed to get user role by user ID", err)
		return nil, errors.New("failed to get user role by user ID")
	}

	var roleIDs []string
	for _, v := range userRole {
		roleIDs = append(roleIDs, v.RoleUUID)
	}

	var roles []*model.Role
	err = ctx.DB.Where("uuid in ?", roleIDs).Find(&roles).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("role not found")
		}
		ctx.Logger.Error("Failed to get role by UUID", err)
		return nil, errors.New("failed to get role by UUID")
	}

	return roles, nil
}
