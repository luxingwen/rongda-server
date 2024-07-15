package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type PurchaseBillController struct {
	PurchaseBillService *service.PurchaseBillService
}

func NewPurchaseBillController() *PurchaseBillController {
	return &PurchaseBillController{
		PurchaseBillService: service.NewPurchaseBillService(),
	}
}

// @Summary 创建采购结算
// @Description 创建采购结算
// @Tags 采购结算
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseBillReq true "采购结算参数"
// @Success 200 {object} model.PurchaseBill
// @Router /api/v1/purchase_bill/create [post]
func (t *PurchaseBillController) CreatePurchaseBill(ctx *app.Context) {
	var param model.PurchaseBillReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseBillService.CreatePurchaseBill(ctx, ctx.GetString("userId"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取采购结算
// @Description 获取采购结算
// @Tags 采购结算
// @Accept  json
// @Produce  json
// @Param uuid path string true "采购结算UUID"
// @Success 200 {object} model.PurchaseBill
// @Router /api/v1/purchase_bill/{uuid} [get]
func (t *PurchaseBillController) GetPurchaseBill(ctx *app.Context) {
	uuid := ctx.Param("uuid")
	bill, err := t.PurchaseBillService.GetPurchaseBill(ctx, uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(bill)
}

// @Summary 更新采购结算
// @Description 更新采购结算
// @Tags 采购结算
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseBill true "采购结算参数"
// @Success 200 {object} model.PurchaseBill
// @Router /api/v1/purchase_bill/update [post]
func (t *PurchaseBillController) UpdatePurchaseBill(ctx *app.Context) {
	var param model.PurchaseBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseBillService.UpdatePurchaseBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除采购结算
// @Description 删除采购结算
// @Tags 采购结算
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "采购结算UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/purchase_bill/delete [post]
func (t *PurchaseBillController) DeletePurchaseBill(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseBillService.DeletePurchaseBill(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取采购结算列表
// @Description 获取采购结算列表
// @Tags 采购结算
// @Accept  json
// @Produce  json
// @Param param body model.ReqPurchaseBillQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/purchase_bill/list [post]
func (t *PurchaseBillController) GetPurchaseBillList(ctx *app.Context) {
	param := &model.ReqPurchaseBillQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	bills, err := t.PurchaseBillService.ListPurchaseBills(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(bills)
}
