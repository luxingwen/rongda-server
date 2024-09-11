package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamMemberService struct {
}

func NewTeamMemberService() *TeamMemberService {
	return &TeamMemberService{}
}

func (s *TeamMemberService) CreateTeamMember(ctx *app.Context, teamMember *model.TeamMember) error {
	teamMember.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	teamMember.UpdatedAt = teamMember.CreatedAt

	// 检查是否已经存在
	var teamMember1 model.TeamMember
	err := ctx.DB.Where("team_uuid = ? AND user_uuid = ?", teamMember.TeamUUID, teamMember.UserUUID).First(&teamMember1).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger.Error("Failed to get team member", err)
		return errors.New("failed to get team member")
	}

	if err == nil && teamMember1.Id > 0 {
		return errors.New("team member already exists")
	}
	teamMember.UUID = uuid.New().String()

	err = ctx.DB.Create(teamMember).Error
	if err != nil {
		ctx.Logger.Error("Failed to create team member", err)
		return errors.New("failed to create team member")
	}
	return nil
}

func (s *TeamMemberService) GetTeamMemberByUUID(ctx *app.Context, uuid string) (*model.TeamMember, error) {
	teamMember := &model.TeamMember{}
	err := ctx.DB.Where("uuid = ?", uuid).First(teamMember).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("team member not found")
		}
		ctx.Logger.Error("Failed to get team member by UUID", err)
		return nil, errors.New("failed to get team member by UUID")
	}
	return teamMember, nil
}

func (s *TeamMemberService) UpdateTeamMember(ctx *app.Context, teamMember *model.TeamMember) error {
	teamMember.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", teamMember.UUID).Updates(teamMember).Error
	if err != nil {
		ctx.Logger.Error("Failed to update team member", err)
		return errors.New("failed to update team member")
	}

	return nil
}

func (s *TeamMemberService) DeleteTeamMember(ctx *app.Context, userUuid string, teamUuid string) error {
	err := ctx.DB.Where("team_uuid = ? AND user_uuid = ?", teamUuid, userUuid).Delete(&model.TeamMember{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete team member", err)
		return errors.New("failed to delete team member")
	}

	return nil
}

// 获取团队成员用户列表
func (s *TeamMemberService) GetTeamMemberUserList(ctx *app.Context, teamUUID string) ([]*model.User, error) {
	var teamMembers []*model.TeamMember

	err := ctx.DB.Where("team_uuid = ?", teamUUID).Find(&teamMembers).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("team member not found")
		}
		ctx.Logger.Error("Failed to get team member by UUID", err)
		return nil, errors.New("failed to get team member by UUID")
	}

	var users []*model.User
	var userIds []string
	for _, teamMember := range teamMembers {
		userIds = append(userIds, teamMember.UserUUID)
	}

	err = ctx.DB.Where("uuid in ?", userIds).Find(&users).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		ctx.Logger.Error("Failed to get user by UUID", err)
		return nil, errors.New("failed to get user by UUID")
	}

	return users, nil
}

// 获取团队成员列表
func (s *TeamMemberService) GetTeamMemberList(ctx *app.Context, params *model.ReqTeamMemberQueryParam) (*model.PagedResponse, error) {

	var (
		teamMembers []*model.TeamMember
		total       int64
	)

	db := ctx.DB.Model(&model.TeamMember{})

	if params.TeamUUID != "" {
		db = db.Where("team_uuid = ?", params.TeamUUID)
	}

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to count team member", err)
		return nil, errors.New("failed to count team member")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&teamMembers).Error
	if err != nil {
		ctx.Logger.Error("Failed to get team member list", err)
		return nil, errors.New("failed to get team member list")
	}

	userUuids := make([]string, 0)
	for _, teamMember := range teamMembers {
		userUuids = append(userUuids, teamMember.UserUUID)
	}

	reslist, err := NewWxUserService().GetWxUserListByUUIDs(ctx, userUuids)
	if err != nil {
		ctx.Logger.Error("Failed to get wxUser list by uuids", err)
		return nil, errors.New("failed to get wxUser list by uuids")
	}

	mMember := make(map[string]*model.TeamMember)
	for _, teamMember := range teamMembers {
		mMember[teamMember.UserUUID] = teamMember
	}

	res := make([]*model.WxUserRes, 0)
	for _, wxUser := range reslist {
		item := &model.WxUserRes{
			WxUser: *wxUser,
		}
		if v, ok := mMember[wxUser.Uuid]; ok {
			item.Role = v.Role
			item.TeamMemberStatus = v.Status
		}
		res = append(res, item)
	}

	return &model.PagedResponse{
		Total: total,
		Data:  res,
	}, nil
}

// UpdateTeamMemberRole
func (s *TeamMemberService) UpdateTeamMemberRole(ctx *app.Context, param *model.ReqTeamMemberUpdateRoleParam) error {

	err := ctx.DB.Model(&model.TeamMember{}).Where("team_uuid = ? AND user_uuid = ?", param.TeamUUID, param.UserUUID).Update("role", param.Role).Error
	if err != nil {
		ctx.Logger.Error("Failed to update team member role", err)
		return errors.New("failed to update team member role")
	}

	return nil
}
