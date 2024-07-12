package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type AgentController struct {
	AgentService *service.AgentService
}

// @Summary 创建代理
// @Description 创建代理
// @Tags 代理
// @Accept  json
// @Produce  json
// @Param param body model.Agent true "代理参数"
// @Success 200 {object} model.Agent
// @Router /api/v1/agent/create [post]
func (a *AgentController) CreateAgent(ctx *app.Context) {
	var param model.Agent
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := a.AgentService.CreateAgent(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新代理
// @Description 更新代理
// @Tags 代理
// @Accept  json
// @Produce  json
// @Param param body model.Agent true "代理参数"
// @Success 200 {object} model.Agent
// @Router /api/v1/agent/update [post]
func (a *AgentController) UpdateAgent(ctx *app.Context) {
	var param model.Agent
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := a.AgentService.UpdateAgent(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除代理
// @Description 删除代理
// @Tags 代理
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "代理UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/agent/delete [post]
func (a *AgentController) DeleteAgent(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := a.AgentService.DeleteAgent(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取代理信息
// @Description 获取代理信息
// @Tags 代理
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "代理UUID"
// @Success 200 {object} model.Agent
// @Router /api/v1/agent/info [post]
func (a *AgentController) GetAgentInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	agent, err := a.AgentService.GetAgentByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(agent)
}

// @Summary 获取代理列表
// @Description 获取代理列表
// @Tags 代理
// @Accept  json
// @Produce  json
// @Param param body model.ReqAgentQueryParam true "查询参数"
// @Success 200 {object} model.AgentQueryResponse
// @Router /api/v1/agent/list [post]
func (a *AgentController) GetAgentList(ctx *app.Context) {
	param := &model.ReqAgentQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	agents, err := a.AgentService.GetAgentList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(agents)
}
