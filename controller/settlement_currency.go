package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SettlementCurrencyController struct {
	SettlementCurrencyService *service.SettlementCurrencyService
}

// @Summary 创建结算币种
// @Description 创建结算币种
// @Tags 结算币种
// @Accept  json
// @Produce  json
// @Param param body model.SettlementCurrency true "结算币种参数"
// @Success 200 {object} model.SettlementCurrency
// @Router /api/v1/settlement_currency/create [post]
func (s *SettlementCurrencyController) CreateSettlementCurrency(ctx *app.Context) {
	var param model.SettlementCurrency
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SettlementCurrencyService.CreateSettlementCurrency(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新结算币种
// @Description 更新结算币种
// @Tags 结算币种
// @Accept  json
// @Produce  json
// @Param param body model.SettlementCurrency true "结算币种参数"
// @Success 200 {object} model.SettlementCurrency
// @Router /api/v1/settlement_currency/update [post]
func (s *SettlementCurrencyController) UpdateSettlementCurrency(ctx *app.Context) {
	var param model.SettlementCurrency
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SettlementCurrencyService.UpdateSettlementCurrency(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除结算币种
// @Description 删除结算币种
// @Tags 结算币种
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "结算币种UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/settlement_currency/delete [post]
func (s *SettlementCurrencyController) DeleteSettlementCurrency(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SettlementCurrencyService.DeleteSettlementCurrency(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取结算币种信息
// @Description 获取结算币种信息
// @Tags 结算币种
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "结算币种UUID"
// @Success 200 {object} model.SettlementCurrency
// @Router /api/v1/settlement_currency/info [post]
func (s *SettlementCurrencyController) GetSettlementCurrencyInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	currency, err := s.SettlementCurrencyService.GetSettlementCurrencyByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(currency)
}

// @Summary 获取结算币种列表
// @Description 获取结算币种列表
// @Tags 结算币种
// @Accept  json
// @Produce  json
// @Param param body model.ReqSettlementCurrencyQueryParam true "查询参数"
// @Success 200 {object} model.SettlementCurrencyQueryResponse
// @Router /api/v1/settlement_currency/list [post]
func (s *SettlementCurrencyController) GetSettlementCurrencyList(ctx *app.Context) {
	param := &model.ReqSettlementCurrencyQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	currencies, err := s.SettlementCurrencyService.GetSettlementCurrencyList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(currencies)
}
