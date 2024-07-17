package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type PurchaseOrderController struct {
	PurchaseOrderService *service.PurchaseOrderService
}

func NewPurchaseOrderController() *PurchaseOrderController {
	return &PurchaseOrderController{
		PurchaseOrderService: service.NewPurchaseOrderService(),
	}
}

// @Summary 创建采购单
// @Description 创建采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseOrderReq true "采购单参数"
// @Success 200 {object} model.PurchaseOrder
// @Router /api/v1/purchase_order/create [post]
func (t *PurchaseOrderController) CreatePurchaseOrder(ctx *app.Context) {
	var param model.PurchaseOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseOrderService.CreatePurchaseOrder(ctx, ctx.GetString("userId"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取采购单
// @Description 获取采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param orderNo path string true "采购单号"
// @Success 200 {object} model.PurchaseOrderRes
// @Router /api/v1/purchase_order/{orderNo} [get]
func (t *PurchaseOrderController) GetPurchaseOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	order, err := t.PurchaseOrderService.GetPurchaseOrder(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(order)
}

// @Summary 获取采购单商品列表
// @Description 获取采购单商品列表
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param orderNo path string true "采购单号"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/purchase_order/{orderNo}/products [get]
func (t *PurchaseOrderController) GetPurchaseOrderProducts(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	products, err := t.PurchaseOrderService.GetPurchaseOrderItems(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(products)
}

// @Summary 更新采购单
// @Description 更新采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseOrder true "采购单参数"
// @Success 200 {object} model.PurchaseOrder
// @Router /api/v1/purchase_order/update [post]
func (t *PurchaseOrderController) UpdatePurchaseOrder(ctx *app.Context) {
	var param model.PurchaseOrder
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseOrderService.UpdatePurchaseOrder(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除采购单
// @Description 删除采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "采购单UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/purchase_order/delete [post]
func (t *PurchaseOrderController) DeletePurchaseOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseOrderService.DeletePurchaseOrder(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取采购单列表
// @Description 获取采购单列表
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.ReqPurchaseOrderQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/purchase_order/list [post]
func (t *PurchaseOrderController) GetPurchaseOrderList(ctx *app.Context) {
	param := &model.ReqPurchaseOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orders, err := t.PurchaseOrderService.ListPurchaseOrders(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orders)
}
