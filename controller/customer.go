package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type CustomerController struct {
	CustomerService *service.CustomerService
}

// @Summary 创建客户
// @Description 创建客户
// @Tags 客户
// @Accept  json
// @Produce  json
// @Param param body model.Customer true "客户参数"
// @Success 200 {object} model.Customer
// @Router /api/v1/customer/create [post]
func (c *CustomerController) CreateCustomer(ctx *app.Context) {
	var param model.Customer
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.CustomerService.CreateCustomer(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新客户
// @Description 更新客户
// @Tags 客户
// @Accept  json
// @Produce  json
// @Param param body model.Customer true "客户参数"
// @Success 200 {object} model.Customer
// @Router /api/v1/customer/update [post]
func (c *CustomerController) UpdateCustomer(ctx *app.Context) {
	var param model.Customer
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.CustomerService.UpdateCustomer(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除客户
// @Description 删除客户
// @Tags 客户
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "客户UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/customer/delete [post]
func (c *CustomerController) DeleteCustomer(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.CustomerService.DeleteCustomer(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取客户信息
// @Description 获取客户信息
// @Tags 客户
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "客户UUID"
// @Success 200 {object} model.Customer
// @Router /api/v1/customer/info [post]
func (c *CustomerController) GetCustomerInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	customer, err := c.CustomerService.GetCustomerByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(customer)
}

// @Summary 获取客户列表
// @Description 获取客户列表
// @Tags 客户
// @Accept  json
// @Produce  json
// @Param param body model.ReqCustomerQueryParam true "查询参数"
// @Success 200 {object} model.CustomerQueryResponse
// @Router /api/v1/customer/list [post]
func (c *CustomerController) GetCustomerList(ctx *app.Context) {
	param := &model.ReqCustomerQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	customers, err := c.CustomerService.GetCustomerList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(customers)
}

// @Summary 获取所有客户列表
// @Description 获取所有客户列表
// @Tags 客户
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CustomerQueryResponse
// @Router /api/v1/customer/all [post]
func (c *CustomerController) GetAllCustomerList(ctx *app.Context) {
	customers, err := c.CustomerService.GetAllCustomers(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(customers)
}

// 获取订单列表
func (c *CustomerController) GetOrderList(ctx *app.Context) {
	param := &model.ReqSalesOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orders, err := c.CustomerService.GetCustomerOrders(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orders)
}

// UpdateOrderStatus
func (c *CustomerController) UpdateOrderStatus(ctx *app.Context) {
	param := &model.ReqSalesOrderUpdateStatusParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	err := c.CustomerService.UpdateOrderStatus(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("更新订单状态成功")
}
