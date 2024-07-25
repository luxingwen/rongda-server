package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SalesOrderController struct {
	SalesOrderService *service.SalesOrderService
}

func NewSalesOrderController() *SalesOrderController {
	return &SalesOrderController{
		SalesOrderService: service.NewSalesOrderService(),
	}
}

// @Summary 创建销售订单
// @Description 创建销售订单
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Param param body model.SalesOrderReq true "销售订单参数"
// @Success 200 {object} model.SalesOrder
// @Router /api/v1/sales_order/create [post]
func (t *SalesOrderController) CreateSalesOrder(ctx *app.Context) {
	var param model.SalesOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOrderService.CreateSalesOrder(ctx, ctx.GetString("user_id"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取销售订单
// @Description 获取销售订单
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Param orderNo path string true "订单号"
// @Success 200 {object} model.SalesOrder
// @Router /api/v1/sales_order/{orderNo} [get]
func (t *SalesOrderController) GetSalesOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	order, err := t.SalesOrderService.GetSalesOrder(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(order)
}

// @Summary 获取订单商品列表
// @Description 获取订单商品列表
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Param orderNo path string true "订单号"
// @Success 200 {object} model.SalesOrder
// @Router /api/v1/sales_order/{orderNo}/products [get]
func (t *SalesOrderController) GetSalesOrderProducts(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	products, err := t.SalesOrderService.GetSalesOrderItems(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(products)
}

// @Summary 更新销售订单
// @Description 更新销售订单
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Param param body model.SalesOrder true "销售订单参数"
// @Success 200 {object} model.SalesOrder
// @Router /api/v1/sales_order/update [post]
func (t *SalesOrderController) UpdateSalesOrder(ctx *app.Context) {
	var param model.SalesOrder
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOrderService.UpdateSalesOrder(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除销售订单
// @Description 删除销售订单
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "销售订单UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/sales_order/delete [post]
func (t *SalesOrderController) DeleteSalesOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOrderService.DeleteSalesOrder(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取销售订单列表
// @Description 获取销售订单列表
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Param param body model.ReqSalesOrderQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/sales_order/list [post]
func (t *SalesOrderController) GetSalesOrderList(ctx *app.Context) {
	param := &model.ReqSalesOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orders, err := t.SalesOrderService.ListSalesOrders(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orders)
}

// @Summary 获取所有可用销售订单
// @Description 获取所有可用销售订单
// @Tags 销售订单
// @Accept  json
// @Produce  json
// @Success 200 {object} model.SalesOrder
// @Router /api/v1/sales_order/all [get]
func (t *SalesOrderController) GetAllSalesOrder(ctx *app.Context) {
	orders, err := t.SalesOrderService.ListAllSalesOrders(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(orders)
}

// 更新订单状态
func (t *SalesOrderController) UpdateSalesOrderStatus(ctx *app.Context) {
	var param model.ReqUpdateOrderStatus
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOrderService.UpdateSalesOrderStatus(ctx, param.OrderNo, param.Status); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}
