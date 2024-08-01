package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type EntrustOrderController struct {
	EntrustOrderService *service.EntrustOrderService
}

// @Summary 创建委托订单
// @Description 创建新的委托订单
// @Tags 委托订单
// @Accept  json
// @Produce  json
// @Param param body model.EntrustOrder true "委托订单参数"
// @Success 200 {object} model.EntrustOrder
// @Router /api/v1/entrust/create [post]
func (c *EntrustOrderController) CreateEntrustOrder(ctx *app.Context) {
	var param model.EntrustOrder
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.EntrustOrderService.CreateEntrustOrder(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新委托订单
// @Description 更新委托订单信息
// @Tags 委托订单
// @Accept  json
// @Produce  json
// @Param param body model.EntrustOrder true "委托订单参数"
// @Success 200 {object} model.EntrustOrder
// @Router /api/v1/entrust/update [post]
func (c *EntrustOrderController) UpdateEntrustOrder(ctx *app.Context) {
	var param model.EntrustOrder
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.EntrustOrderService.UpdateEntrustOrder(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除委托订单
// @Description 删除指定的委托订单
// @Tags 委托订单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "委托订单UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/entrust/delete [post]
func (c *EntrustOrderController) DeleteEntrustOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.EntrustOrderService.DeleteEntrustOrder(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取委托订单信息
// @Description 获取指定UUID的委托订单信息
// @Tags 委托订单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "委托订单UUID"
// @Success 200 {object} model.EntrustOrder
// @Router /api/v1/entrust/info [post]
func (c *EntrustOrderController) GetEntrustOrderInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	order, err := c.EntrustOrderService.GetEntrustOrderByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(order)
}

// @Summary 获取委托订单列表
// @Description 获取符合条件的委托订单列表
// @Tags 委托订单
// @Accept  json
// @Produce  json
// @Param param body model.ReqEntrustOrderQueryParam true "查询参数"
// @Success 200 {object} model.EntrustOrderQueryResponse
// @Router /api/v1/entrust/list [post]
func (c *EntrustOrderController) GetEntrustOrderList(ctx *app.Context) {
	param := &model.ReqEntrustOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orders, err := c.EntrustOrderService.GetEntrustOrderList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orders)
}

func (c *EntrustOrderController) GetEntrustOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	order, err := c.EntrustOrderService.GetEntrustOrderByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(order)
}
