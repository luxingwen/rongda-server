package service

import (
	"errors"
	"time"

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

// 根据邀请码获取企业信息
func (s *TeamInviteService) GetTeamByInviteCode(ctx *app.Context, inviteCode string) (*model.Team, error) {
	invite, err := s.GetInviteByCode(ctx, inviteCode)
	if err != nil {
		return nil, err
	}

	team := &model.TeamRef{}
	err = ctx.DB.Where("team_uuid = ?", invite.TeamUuid).First(team).Error
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("team not found")
		}
		ctx.Logger.Error("Failed to get team by uuid", err)
		return nil, errors.New("failed to get team by uuid")
	}

	if team.Category == model.TeamCategoryCustomer {
		customeInfo, err := NewCustomerService().GetCustomerByUUID(ctx, team.TeamUuid)
		if err != nil {
			return nil, err
		}
		return &model.Team{
			UUID: customeInfo.Uuid,
			Name: customeInfo.Name,
		}, nil
	}

	return nil, errors.New("team not found")
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

// JoinTeamByInviteCode 通过邀请码加入团队
func (s *TeamInviteService) JoinTeamByInviteCode(ctx *app.Context, teamuuid string, userUUID string, inviteCode string) error {

	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		// 从teamRef表中获取team信息
		teamRef := &model.TeamRef{}
		err := tx.Where("team_uuid = ?", teamuuid).First(teamRef).Error
		if err != nil {
			ctx.Logger.Error("Failed to get team by uuid", err)
			return errors.New("failed to get team by uuid")
		}

		if teamRef.TeamUuid == "" {
			return errors.New("team not found")
		}

		// 查询是否已经加入团队
		teamMember := &model.TeamMember{}
		err = tx.Where("team_uuid = ? AND user_uuid = ?", teamRef.TeamUuid, userUUID).First(teamMember).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			ctx.Logger.Error("Failed to get team member", err)
			return errors.New("failed to get team member")
		}

		if teamMember.Id > 0 {
			return errors.New("你已经加入过该团队")
		}

		// 创建teammember
		teamMember = &model.TeamMember{
			TeamUUID:   teamRef.TeamUuid,
			UserUUID:   userUUID,
			Role:       "成员",
			Status:     "未审核",
			InviteCode: inviteCode,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		}

		err = tx.Create(teamMember).Error
		if err != nil {
			ctx.Logger.Error("Failed to create team member", err)
			return errors.New("failed to create team member")
		}

		return nil

	})

	if err != nil {
		return err
	}

	return nil
}
