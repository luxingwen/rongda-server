package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamService struct {
}

func NewTeamService() *TeamService {
	return &TeamService{}
}

func (s *TeamService) CreateTeam(ctx *app.Context, team *model.Team) error {
	team.UUID = uuid.New().String()

	team.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	team.UpdatedAt = team.CreatedAt

	err := ctx.DB.Create(team).Error
	if err != nil {
		ctx.Logger.Error("Failed to create team", err)
		return errors.New("failed to create team")
	}
	return nil
}

func (s *TeamService) GetTeamByUUID(ctx *app.Context, uuid string) (*model.ResTeam, error) {
	team := &model.TeamRef{}
	err := ctx.DB.Where("team_uuid = ?", uuid).First(team).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("team not found")
		}
		ctx.Logger.Error("Failed to get team by UUID", err)
		return nil, errors.New("failed to get team by UUID")
	}

	return s.GetTeamInfoByTeamRef(ctx, team)
}

func (s *TeamService) GetTeamInfoByTeamRef(ctx *app.Context, teamRef *model.TeamRef) (*model.ResTeam, error) {
	if teamRef.Category == model.TeamCategoryCustomer {
		curtomeInfo, err := NewCustomerService().GetCustomerByUUID(ctx, teamRef.TeamUuid)
		if err != nil {
			ctx.Logger.Error("Failed to get customer by uuid", err)
			return nil, errors.New("failed to get customer by uuid")
		}
		return &model.ResTeam{
			TeamUuid: teamRef.TeamUuid,
			Name:     curtomeInfo.Name,
			Category: teamRef.Category,
		}, nil
	}

	if teamRef.Category == model.TeamCategorySupplier {
		supplierInfo, err := NewSupplierService().GetSupplierByUUID(ctx, teamRef.TeamUuid)
		if err != nil {
			ctx.Logger.Error("Failed to get supplier by uuid", err)
			return nil, errors.New("failed to get supplier by uuid")
		}
		return &model.ResTeam{
			TeamUuid: teamRef.TeamUuid,
			Name:     supplierInfo.Name,
			Category: teamRef.Category,
		}, nil
	}

	if teamRef.Category == model.TeamCategoryAgent {
		agentInfo, err := NewAgentService().GetAgentByUUID(ctx, teamRef.TeamUuid)
		if err != nil {
			ctx.Logger.Error("Failed to get agent by uuid", err)
			return nil, errors.New("failed to get agent by uuid")
		}
		return &model.ResTeam{
			TeamUuid: teamRef.TeamUuid,
			Name:     agentInfo.Name,
			Category: teamRef.Category,
		}, nil
	}
	return nil, errors.New("failed to get team info by team ref")
}

func (s *TeamService) UpdateTeam(ctx *app.Context, team *model.Team) error {
	team.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.DB.Where("uuid = ?", team.UUID).Updates(team).Error
	if err != nil {
		ctx.Logger.Error("Failed to update team", err)
		return errors.New("failed to update team")
	}

	return nil
}

func (s *TeamService) DeleteTeam(ctx *app.Context, uuid string) error {
	err := ctx.DB.Where("uuid = ?", uuid).Delete(&model.Team{}).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete team", err)
		return errors.New("failed to delete team")
	}

	return nil
}

// 获取团队列表
func (s *TeamService) GetTeamList(ctx *app.Context, param *model.ReqTeamQueryParam) (r *model.PagedResponse, err error) {
	var (
		teamList []*model.Team
		total    int64
	)

	db := ctx.DB.Model(&model.Team{})

	if param.Name != "" {
		db = db.Where("name like ?", "%"+param.Name+"%")
	}

	if err = db.Offset(param.GetOffset()).Limit(param.PageSize).Find(&teamList).Error; err != nil {
		return
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	r = &model.PagedResponse{
		Total:    total,
		Current:  param.Current,
		PageSize: param.PageSize,
		Data:     teamList,
	}

	return
}

// 根据团队uuid列表获取团队Ref列表
func (s *TeamService) GetTeamRefListByUUIDs(ctx *app.Context, teamUUIDs []string) (r map[string]*model.TeamRef, err error) {
	var teamList []*model.TeamRef
	err = ctx.DB.Where("team_uuid in (?)", teamUUIDs).Find(&teamList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get team list by uuids", err)
		return nil, errors.New("failed to get team list by uuids")
	}

	r = make(map[string]*model.TeamRef)
	for _, team := range teamList {
		r[team.TeamUuid] = team
	}

	return
}

// 根据用户ID获取团队列表
func (s *TeamService) GetTeamListByUserID(ctx *app.Context, userID string) (teamList []*model.ResTeam, err error) {
	var teamMemberList []*model.TeamMember
	err = ctx.DB.Where("user_uuid = ?", userID).Find(&teamMemberList).Error
	if err != nil {
		ctx.Logger.Error("Failed to get team member list by user ID", err)
		return nil, errors.New("failed to get team member list by user ID")
	}

	var teamUUIDs []string
	for _, teamMember := range teamMemberList {
		teamUUIDs = append(teamUUIDs, teamMember.TeamUUID)
	}

	if len(teamUUIDs) == 0 {
		return
	}

	teamRefMap, err := s.GetTeamRefListByUUIDs(ctx, teamUUIDs)
	if err != nil {
		ctx.Logger.Error("Failed to get team ref list by uuids", err)

		return
	}

	teamList = make([]*model.ResTeam, 0)

	for _, teamRef := range teamRefMap {
		if teamRef.Category == model.TeamCategoryCustomer {
			curtomeInfo, err := NewCustomerService().GetCustomerByUUID(ctx, teamRef.TeamUuid)
			if err != nil {
				ctx.Logger.Error("Failed to get customer by uuid", err)
				return nil, errors.New("failed to get customer by uuid")
			}
			teamList = append(teamList, &model.ResTeam{
				TeamUuid: teamRef.TeamUuid,
				Name:     curtomeInfo.Name,
				Category: teamRef.Category,
			})
		}

		if teamRef.Category == model.TeamCategorySupplier {
			supplierInfo, err := NewSupplierService().GetSupplierByUUID(ctx, teamRef.TeamUuid)
			if err != nil {
				ctx.Logger.Error("Failed to get supplier by uuid", err)
				return nil, errors.New("failed to get supplier by uuid")
			}
			teamList = append(teamList, &model.ResTeam{
				TeamUuid: teamRef.TeamUuid,
				Name:     supplierInfo.Name,
				Category: teamRef.Category,
			})
		}

		if teamRef.Category == model.TeamCategoryAgent {
			agentInfo, err := NewAgentService().GetAgentByUUID(ctx, teamRef.TeamUuid)
			if err != nil {
				ctx.Logger.Error("Failed to get agent by uuid", err)
				return nil, errors.New("failed to get agent by uuid")
			}
			teamList = append(teamList, &model.ResTeam{
				TeamUuid: teamRef.TeamUuid,
				Name:     agentInfo.Name,
				Category: teamRef.Category,
			})
		}

	}

	return
}
