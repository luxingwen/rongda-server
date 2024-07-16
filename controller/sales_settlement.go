package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SalesSettlementController struct {
	SalesSettlementService *service.SalesSettlementService
}

func NewSalesSettlementController() *SalesSettlementController {
	return &SalesSettlementController{
		SalesSettlementService: service.NewSalesSettlementService(),
	}
}

// @Summary 创建销售订单结算
// @Description 创建销售订单结算
// @Tags 销售订单结算
// @Accept  json
// @Produce  json
// @Param param body model.SalesSettlementReq true "销售订单结算参数"
// @Success 200 {object} model.SalesSettlement
// @Router /api/v1/sales_settlement/create [post]
func (t *SalesSettlementController) CreateSalesSettlement(ctx *app.Context) {
	var param model.SalesSettlementReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesSettlementService.CreateSalesSettlement(ctx, ctx.GetString("user_id"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取销售订单结算
// @Description 获取销售订单结算
// @Tags 销售订单结算
// @Accept  json
// @Produce  json
// @Param uuid path string true "结算UUID"
// @Success 200 {object} model.SalesSettlement
// @Router /api/v1/sales_settlement/{uuid} [get]
func (t *SalesSettlementController) GetSalesSettlement(ctx *app.Context) {
	uuid := ctx.Param("uuid")
	settlement, err := t.SalesSettlementService.GetSalesSettlement(ctx, uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(settlement)
}

// @Summary 更新销售订单结算
// @Description 更新销售订单结算
// @Tags 销售订单结算
// @Accept  json
// @Produce  json
// @Param param body model.SalesSettlement true "销售订单结算参数"
// @Success 200 {object} model.SalesSettlement
// @Router /api/v1/sales_settlement/update [post]
func (t *SalesSettlementController) UpdateSalesSettlement(ctx *app.Context) {
	var param model.SalesSettlement
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesSettlementService.UpdateSalesSettlement(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除销售订单结算
// @Description 删除销售订单结算
// @Tags 销售订单结算
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "结算UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/sales_settlement/delete [post]
func (t *SalesSettlementController) DeleteSalesSettlement(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesSettlementService.DeleteSalesSettlement(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取销售订单结算列表
// @Description 获取销售订单结算列表
// @Tags 销售订单结算
// @Accept  json
// @Produce  json
// @Param param body model.ReqSalesSettlementQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/sales_settlement/list [post]
func (t *SalesSettlementController) GetSalesSettlementList(ctx *app.Context) {
	param := &model.ReqSalesSettlementQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	settlements, err := t.SalesSettlementService.ListSalesSettlements(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(settlements)
}
