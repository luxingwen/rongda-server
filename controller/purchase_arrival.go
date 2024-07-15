package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type PurchaseArrivalController struct {
	PurchaseArrivalService *service.PurchaseArrivalService
}

func NewPurchaseArrivalController() *PurchaseArrivalController {
	return &PurchaseArrivalController{
		PurchaseArrivalService: service.NewPurchaseArrivalService(),
	}
}

// @Summary 创建到货
// @Description 创建到货
// @Tags 到货
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseArrivalReq true "到货参数"
// @Success 200 {object} model.PurchaseArrival
// @Router /api/v1/purchase_arrival/create [post]
func (t *PurchaseArrivalController) CreatePurchaseArrival(ctx *app.Context) {
	var param model.PurchaseArrivalReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseArrivalService.CreatePurchaseArrival(ctx, ctx.GetString("user_id"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取到货
// @Description 获取到货
// @Tags 到货
// @Accept  json
// @Produce  json
// @Param uuid path string true "到货UUID"
// @Success 200 {object} model.PurchaseArrival
// @Router /api/v1/purchase_arrival/{uuid} [get]
func (t *PurchaseArrivalController) GetPurchaseArrival(ctx *app.Context) {
	uuid := ctx.Param("uuid")
	arrival, err := t.PurchaseArrivalService.GetPurchaseArrival(ctx, uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(arrival)
}

// @Summary 更新到货
// @Description 更新到货
// @Tags 到货
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseArrival true "到货参数"
// @Success 200 {object} model.PurchaseArrival
// @Router /api/v1/purchase_arrival/update [post]
func (t *PurchaseArrivalController) UpdatePurchaseArrival(ctx *app.Context) {
	var param model.PurchaseArrival
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseArrivalService.UpdatePurchaseArrival(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除到货
// @Description 删除到货
// @Tags 到货
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "到货UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/purchase_arrival/delete [post]
func (t *PurchaseArrivalController) DeletePurchaseArrival(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseArrivalService.DeletePurchaseArrival(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取到货列表
// @Description 获取到货列表
// @Tags 到货
// @Accept  json
// @Produce  json
// @Param param body model.ReqPurchaseArrivalQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/purchase_arrival/list [post]
func (t *PurchaseArrivalController) GetPurchaseArrivalList(ctx *app.Context) {
	param := &model.ReqPurchaseArrivalQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	arrivals, err := t.PurchaseArrivalService.ListPurchaseArrivals(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(arrivals)
}
