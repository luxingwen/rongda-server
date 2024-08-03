package service

import (
	"errors"

	"sgin/model"
	"sgin/pkg/app"
	"sgin/pkg/utils"

	"gorm.io/gorm"
)

type TeamInviteService struct {
}

func NewTeamInviteService() *TeamInviteService {
	return &TeamInviteService{}
}

// 创建邀请
func (s *TeamInviteService) GetOrCreateInvite(ctx *app.Context, invite *model.TeamInvite) (r *model.TeamInvite, err error) {

	// 先查询是否已经存在邀请
	var oldInvite model.TeamInvite
	err = ctx.DB.Where("team_uuid = ? AND inviter = ?", invite.TeamUuid, invite.Inviter).First(&oldInvite).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger.Error("Failed to get invite", err)
		return nil, errors.New("failed to get invite")
	}

	if err == nil && oldInvite.Id > 0 {
		return &oldInvite, nil
	}

	for {
		inviteCode := utils.GenerateInviteCode(6)
		err := ctx.DB.Where("invite_code = ?", inviteCode).First(&model.TeamInvite{}).Error
		if err == gorm.ErrRecordNotFound {
			invite.InviteCode = inviteCode
			break
		}
	}

	err = ctx.DB.Create(invite).Error
	if err != nil {
		ctx.Logger.Error("Failed to create invite", err)
		return nil, errors.New("failed to create invite")
	}
	return invite, nil
}

// 获取邀请
func (s *TeamInviteService) GetInviteByCode(ctx *app.Context, inviteCode string) (*model.TeamInvite, error) {
	invite := &model.TeamInvite{}
	err := ctx.DB.Where("invite_code = ?", inviteCode).First(invite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invite not found")
		}
		ctx.Logger.Error("Failed to get invite by code", err)
		return nil, errors.New("failed to get invite by code")
	}
	return invite, nil
}

// 更新邀请状态
func (s *TeamInviteService) UpdateInviteStatus(ctx *app.Context, inviteCode string, status int) error {
	invite := &model.TeamInvite{}

	err := ctx.DB.Where("invite_code = ?", inviteCode).First(invite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("invite not found")
		}
		ctx.Logger.Error("Failed to get invite by code", err)
		return errors.New("failed to get invite by code")
	}

	err = ctx.DB.Save(invite).Error
	if err != nil {
		ctx.Logger.Error("Failed to update invite status", err)
		return errors.New("failed to update invite status")
	}

	return nil
}

// 删除邀请
func (s *TeamInviteService) DeleteInvite(ctx *app.Context, inviteCode string) error {
	err := ctx.DB.Model(&model.TeamInvite{}).Where("invite_code = ?", inviteCode).Update("status", 2).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete invite", err)
		return errors.New("failed to delete invite")
	}

	return nil
}

// 获取邀请列表
func (s *TeamInviteService) GetInviteList(ctx *app.Context, params *model.ReqInviteQueryParam) (*model.PagedResponse, error) {
	var (
		invites []*model.TeamInvite
		total   int64
	)

	db := ctx.DB.Model(&model.TeamInvite{})

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get invite count", err)
		return nil, errors.New("failed to get invite count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&invites).Error
	if err != nil {
		ctx.Logger.Error("Failed to get invite list", err)
		return nil, errors.New("failed to get invite list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  invites,
	}, nil
}
