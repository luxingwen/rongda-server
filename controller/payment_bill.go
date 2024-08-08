package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type PaymentBillController struct {
	PaymentBillService *service.PaymentBillService
}

// @Summary 创建支付账单
// @Description 创建支付账单
// @Tags 支付账单
// @Accept  json
// @Produce  json
// @Param param body model.PaymentBill true "支付账单参数"
// @Success 200 {object} model.PaymentBill
// @Router /api/v1/payment-bill/create [post]
func (p *PaymentBillController) CreatePaymentBill(ctx *app.Context) {
	var param model.PaymentBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PaymentBillService.CreatePaymentBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新支付账单
// @Description 更新支付账单
// @Tags 支付账单
// @Accept  json
// @Produce  json
// @Param param body model.PaymentBill true "支付账单参数"
// @Success 200 {object} model.PaymentBill
// @Router /api/v1/payment-bill/update [post]
func (p *PaymentBillController) UpdatePaymentBill(ctx *app.Context) {
	var param model.PaymentBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PaymentBillService.UpdatePaymentBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除支付账单
// @Description 删除支付账单
// @Tags 支付账单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "支付账单UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/payment-bill/delete [post]
func (p *PaymentBillController) DeletePaymentBill(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PaymentBillService.DeletePaymentBill(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取支付账单信息
// @Description 获取支付账单信息
// @Tags 支付账单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "支付账单UUID"
// @Success 200 {object} model.PaymentBill
// @Router /api/v1/payment-bill/info [post]
func (p *PaymentBillController) GetPaymentBillInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	paymentBill, err := p.PaymentBillService.GetPaymentBillByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(paymentBill)
}

// @Summary 获取支付账单列表
// @Description 获取支付账单列表
// @Tags 支付账单
// @Accept  json
// @Produce  json
// @Param param body model.ReqPaymentBillQueryParam true "查询参数"
// @Success 200 {object} model.PaymentBillQueryResponse
// @Router /api/v1/payment-bill/list [post]
func (p *PaymentBillController) GetPaymentBillList(ctx *app.Context) {
	param := &model.ReqPaymentBillQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	paymentBills, err := p.PaymentBillService.GetPaymentBillList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(paymentBills)
}

// UpdatePaymentBillStatus
func (p *PaymentBillController) UpdatePaymentBillStatus(ctx *app.Context) {
	var param model.ReqUpdatePaymentBillStatusParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PaymentBillService.UpdatePaymentBillStatus(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}
