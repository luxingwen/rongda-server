package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SettlementController struct {
	SettlementService *service.SettlementService
}

// @Summary 创建结算单
// @Description 创建结算单
// @Tags 结算单
// @Accept  json
// @Produce  json
// @Param param body model.Settlement true "结算单参数"
// @Success 200 {object} model.Settlement
// @Router /api/v1/settlement/create [post]
func (s *SettlementController) CreateSettlement(ctx *app.Context) {
	var param model.Settlement
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	param.Status = model.SettlementStatusPending

	if err := s.SettlementService.CreateSettlement(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新结算单
// @Description 更新结算单
// @Tags 结算单
// @Accept  json
// @Produce  json
// @Param param body model.Settlement true "结算单参数"
// @Success 200 {object} model.Settlement
// @Router /api/v1/settlement/update [post]
func (s *SettlementController) UpdateSettlement(ctx *app.Context) {
	var param model.Settlement
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SettlementService.UpdateSettlement(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除结算单
// @Description 删除结算单
// @Tags 结算单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "结算单UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/settlement/delete [post]
func (s *SettlementController) DeleteSettlement(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SettlementService.DeleteSettlement(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取结算单信息
// @Description 获取结算单信息
// @Tags 结算单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "结算单UUID"
// @Success 200 {object} model.Settlement
// @Router /api/v1/settlement/info [post]
func (s *SettlementController) GetSettlementInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	settlement, err := s.SettlementService.GetSettlementByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(settlement)
}

// @Summary 获取结算单列表
// @Description 获取结算单列表
// @Tags 结算单
// @Accept  json
// @Produce  json
// @Param param body model.ReqSettlementQueryParam true "查询参数"
// @Success 200 {object} model.SettlementQueryResponse
// @Router /api/v1/settlement/list [post]
func (s *SettlementController) GetSettlementList(ctx *app.Context) {
	param := &model.ReqSettlementQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	settlements, err := s.SettlementService.GetSettlementList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(settlements)
}
