package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type RemittanceBillController struct {
	RemittanceBillService *service.RemittanceBillService
}

// @Summary 创建付汇账单
// @Description 创建付汇账单
// @Tags 付汇账单
// @Accept  json
// @Produce  json
// @Param param body model.RemittanceBill true "付汇账单参数"
// @Success 200 {object} model.RemittanceBill
// @Router /api/v1/remittance-bill/create [post]
func (r *RemittanceBillController) CreateRemittanceBill(ctx *app.Context) {
	var param model.RemittanceBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := r.RemittanceBillService.CreateRemittanceBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新付汇账单
// @Description 更新付汇账单
// @Tags 付汇账单
// @Accept  json
// @Produce  json
// @Param param body model.RemittanceBill true "付汇账单参数"
// @Success 200 {object} model.RemittanceBill
// @Router /api/v1/remittance-bill/update [post]
func (r *RemittanceBillController) UpdateRemittanceBill(ctx *app.Context) {
	var param model.RemittanceBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := r.RemittanceBillService.UpdateRemittanceBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除付汇账单
// @Description 删除付汇账单
// @Tags 付汇账单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "付汇账单UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/remittance-bill/delete [post]
func (r *RemittanceBillController) DeleteRemittanceBill(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := r.RemittanceBillService.DeleteRemittanceBill(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取付汇账单信息
// @Description 获取付汇账单信息
// @Tags 付汇账单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "付汇账单UUID"
// @Success 200 {object} model.RemittanceBill
// @Router /api/v1/remittance-bill/info [post]
func (r *RemittanceBillController) GetRemittanceBillInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	remittanceBill, err := r.RemittanceBillService.GetRemittanceBillByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(remittanceBill)
}

// @Summary 获取付汇账单列表
// @Description 获取付汇账单列表
// @Tags 付汇账单
// @Accept  json
// @Produce  json
// @Param param body model.ReqRemittanceBillQueryParam true "查询参数"
// @Success 200 {object} model.RemittanceBillQueryResponse
// @Router /api/v1/remittance-bill/list [post]
func (r *RemittanceBillController) GetRemittanceBillList(ctx *app.Context) {
	param := &model.ReqRemittanceBillQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	remittanceBills, err := r.RemittanceBillService.GetRemittanceBillList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(remittanceBills)
}

// UpdateRemittanceBillStatus
// func (r *RemittanceBillController) UpdateRemittanceBillStatus(ctx *app.Context) {
// 	var param model.ReqUpdateRemittanceBillStatusParam
// 	if err := ctx.ShouldBindJSON(&param); err != nil {
// 		ctx.JSONError(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	if err := r.RemittanceBillService.UpdateRemittanceBillStatus(ctx, &param); err != nil {
// 		ctx.JSONError(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	ctx.JSONSuccess("ok")
// }
