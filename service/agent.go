package service

import (
	"errors"
	"time"

	"sgin/model"
	"sgin/pkg/app"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AgentService struct {
}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (s *AgentService) CreateAgent(ctx *app.Context, agent *model.Agent) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	agent.CreatedAt = now
	agent.UpdatedAt = now
	agent.Uuid = uuid.New().String()

	err := ctx.DB.Create(agent).Error
	if err != nil {
		ctx.Logger.Error("Failed to create agent", err)
		return errors.New("failed to create agent")
	}
	return nil
}

func (s *AgentService) GetAgentByUUID(ctx *app.Context, uuid string) (*model.Agent, error) {
	agent := &model.Agent{}
	err := ctx.DB.Where("uuid = ?", uuid).First(agent).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("agent not found")
		}
		ctx.Logger.Error("Failed to get agent by UUID", err)
		return nil, errors.New("failed to get agent by UUID")
	}
	return agent, nil
}

func (s *AgentService) UpdateAgent(ctx *app.Context, agent *model.Agent) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	agent.UpdatedAt = now
	err := ctx.DB.Where("uuid = ?", agent.Uuid).Updates(agent).Error
	if err != nil {
		ctx.Logger.Error("Failed to update agent", err)
		return errors.New("failed to update agent")
	}

	return nil
}

func (s *AgentService) DeleteAgent(ctx *app.Context, uuid string) error {
	err := ctx.DB.Model(&model.Agent{}).Where("uuid = ?", uuid).Update("is_deleted", 1).Error
	if err != nil {
		ctx.Logger.Error("Failed to delete agent", err)
		return errors.New("failed to delete agent")
	}

	return nil
}

// GetAgentList retrieves a list of agents based on query parameters
func (s *AgentService) GetAgentList(ctx *app.Context, params *model.ReqAgentQueryParam) (*model.PagedResponse, error) {
	var (
		agents []*model.Agent
		total  int64
	)

	db := ctx.DB.Model(&model.Agent{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	db = db.Where("is_deleted = ?", 0)

	err := db.Count(&total).Error
	if err != nil {
		ctx.Logger.Error("Failed to get agent count", err)
		return nil, errors.New("failed to get agent count")
	}

	err = db.Offset(params.GetOffset()).Limit(params.PageSize).Find(&agents).Error
	if err != nil {
		ctx.Logger.Error("Failed to get agent list", err)
		return nil, errors.New("failed to get agent list")
	}

	return &model.PagedResponse{
		Total: total,
		Data:  agents,
	}, nil
}
