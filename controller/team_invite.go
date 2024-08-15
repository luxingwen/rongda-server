package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type TeamInviteController struct {
	TeamInviteService *service.TeamInviteService
}

// @Summary 创建邀请
// @Description 创建邀请
// @Tags 邀请
// @Accept  json
// @Produce  json
// @Param param body model.TeamInvite true "邀请参数"
// @Success 200 {object} model.TeamInvite
// @Router /api/v1/invite/create [post]
func (t *TeamInviteController) CreateInvite(ctx *app.Context) {
	var param model.TeamInvite
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	wxUserId := ctx.GetString("wx_user_id")

	if wxUserId == "" {
		ctx.JSONError(http.StatusBadRequest, "请先登录")
		return
	}

	param.Inviter = wxUserId
	rinvate, err := t.TeamInviteService.GetOrCreateInvite(ctx, &param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(rinvate)
}

// @Summary 更新邀请状态
// @Description 更新邀请状态
// @Tags 邀请
// @Accept  json
// @Produce  json
// @Param param body model.ReqUpdateInviteStatus true "更新邀请状态参数"
// @Success 200 {object} model.TeamInvite
// @Router /api/v1/invite/update [post]
func (t *TeamInviteController) UpdateInviteStatus(ctx *app.Context) {
	var param model.ReqUpdateInviteStatus
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.TeamInviteService.UpdateInviteStatus(ctx, param.InviteCode, param.Status); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 删除邀请
// @Description 删除邀请
// @Tags 邀请
// @Accept  json
// @Produce  json
// @Param param body model.ReqInviteCodeParam true "邀请代码"
// @Success 200 {string} string "ok"
// @Router /api/v1/invite/delete [post]
func (t *TeamInviteController) DeleteInvite(ctx *app.Context) {
	var param model.ReqInviteCodeParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.TeamInviteService.DeleteInvite(ctx, param.InviteCode); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取邀请信息
// @Description 获取邀请信息
// @Tags 邀请
// @Accept  json
// @Produce  json
// @Param param body model.ReqInviteCodeParam true "邀请代码"
// @Success 200 {object} model.TeamInvite
// @Router /api/v1/invite/info [post]
func (t *TeamInviteController) GetInviteInfo(ctx *app.Context) {
	var param model.ReqInviteCodeParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	invite, err := t.TeamInviteService.GetInviteByCode(ctx, param.InviteCode)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(invite)
}

// @Summary 获取邀请列表
// @Description 获取邀请列表
// @Tags 邀请
// @Accept  json
// @Produce  json
// @Param param body model.ReqInviteQueryParam true "查询参数"
// @Success 200 {object} model.InviteQueryResponse
// @Router /api/v1/invite/list [post]
func (t *TeamInviteController) GetInviteList(ctx *app.Context) {
	param := &model.ReqInviteQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	invites, err := t.TeamInviteService.GetInviteList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(invites)
}

// GetTeamByInviteCode
func (t *TeamInviteController) GetTeamByInviteCode(ctx *app.Context) {
	param := &model.ReqInviteCodeParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	team, err := t.TeamInviteService.GetTeamByInviteCode(ctx, param.InviteCode)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(team)
}

// JoinTeamByInviteCode
func (t *TeamInviteController) JoinTeamByInviteCode(ctx *app.Context) {
	param := &model.ReqInviteCodeParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	wxUserId := ctx.GetString("wx_user_id")

	if wxUserId == "" {
		ctx.JSONError(http.StatusBadRequest, "请先登录")
		return
	}

	if param.TeamUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "团队UUID不能为空")
		return
	}

	err := t.TeamInviteService.JoinTeamByInviteCode(ctx, param.TeamUuid, wxUserId, param.Code)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("ok")
}
