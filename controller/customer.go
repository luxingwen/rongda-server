package controller

import (
	"net/http"
	"path/filepath"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
	"time"
)

type CustomerController struct {
	CustomerService               *service.CustomerService
	PaymentBillService            *service.PaymentBillService
	SettlementService             *service.SettlementService
	StorehouseProductOpLogService *service.StorehouseProductOpLogService
	StorehouseProductService      *service.StorehouseProductService
	StorehouseService             *service.StorehouseService
	AgreementService              *service.AgreementService
	SalesOrderService             *service.SalesOrderService
	PurchaseOrderService          *service.PurchaseOrderService
	RemittanceBillService         *service.RemittanceBillService
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

// GetOrderItemList
func (c *CustomerController) GetOrderItemList(ctx *app.Context) {
	param := &model.ReqSalesOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	items, err := c.SalesOrderService.GetSalesOrderItems(ctx, param.OrderNo)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(items)
}

// GetOrderInfo
func (c *CustomerController) GetOrderInfo(ctx *app.Context) {
	param := &model.ReqSalesOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	order, err := c.SalesOrderService.GetSalesOrder(ctx, param.OrderNo)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(order)
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

// UpdateOrderStatusPaidComfirm
func (c *CustomerController) UpdateOrderStatusPaidComfirm(ctx *app.Context) {
	param := &model.ReqPaymentBillOrderStatusPaidComfirm{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	err := c.PaymentBillService.UpdatePaymentBillStatusPaidPendingConfirm(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("确认支付成功")
}

// GetSettlementList
func (c *CustomerController) GetSettlementList(ctx *app.Context) {
	param := &model.ReqSettlementQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.TeamUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "团队UUID不能为空")
		return
	}

	settlements, err := c.SettlementService.GetSettlementList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(settlements)
}

// GetStorehouseProductInOutList
func (c *CustomerController) GetStorehouseProductInOutList(ctx *app.Context) {
	param := &model.ReqStorehouseProductOpLogListParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.TeamUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "团队UUID不能为空")
		return
	}
	param.OpTypes = []int{model.StorehouseProductOpLogOpTypeInbound, model.StorehouseProductOpLogOpTypeOutbound}

	list, err := c.StorehouseProductOpLogService.GetStorehouseProductOpLogList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(list)
}

// GetStorehouseProductList
func (c *CustomerController) GetStorehouseProductList(ctx *app.Context) {
	param := &model.ReqStorehouseProductQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.TeamUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "团队UUID不能为空")
		return
	}

	list, err := c.StorehouseProductService.ListProducts(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(list)
}

// PaymentBillUpdateIsAdvance
func (c *CustomerController) PaymentBillUpdateIsAdvance(ctx *app.Context) {
	param := &model.ReqUpdatePaymentBillIsAdvanceParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	err := c.PaymentBillService.UpdatePaymentBillIsAdvance(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("更新是否预付款成功")
}

// GetPaymentBillList
func (c *CustomerController) GetPaymentBillList(ctx *app.Context) {
	param := &model.ReqPaymentBillQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.TeamUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "团队UUID不能为空")
		return
	}

	list, err := c.PaymentBillService.GetPaymentBillList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(list)
}

// UpdateAgreementSign
func (c *CustomerController) UpdateAgreementSign(ctx *app.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orderNo := ctx.PostForm("order_no")

	if orderNo == "" {
		ctx.JSONError(http.StatusBadRequest, "order_no is required")
		return
	}

	uuidStr := ctx.PostForm("uuid")
	if uuidStr == "" {
		ctx.JSONError(http.StatusBadRequest, "uuid is required")
		return
	}

	// 保存头像

	userid := ctx.GetString("wx_user_id")
	if userid == "" {
		ctx.JSONError(http.StatusBadRequest, "user_id is required")
		return
	}

	// 获取文件后缀

	extfile := filepath.Ext(file.Filename)

	filename := "/sales_order/agreement/sign_" + orderNo + extfile

	err = ctx.SaveUploadedFile(file, ctx.Config.Upload.Dir+filename)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	param := &model.Agreement{
		Uuid:           uuidStr,
		SignatureImage: filename,
		SignatureTime:  time.Now().Format("2006-01-02 15:04:05"),
		SignatureUser:  userid,
	}

	err = c.AgreementService.UpdateAgreementSign(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("更新合同签署成功")
}

// UpdatePurchaseOrderStorehouse
func (c *CustomerController) UpdatePurchaseOrderStorehouse(ctx *app.Context) {
	param := &model.ReqUpdatePurchaseOrderStorehouseParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	err := c.PurchaseOrderService.UpdatePurchaseOrderStorehouse(ctx, param.PurchaseOrderNo, param.StorehouseUuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("更新采购订单仓库成功")
}

// CreateStorehouseOutboundOrder
func (c *CustomerController) CreateStorehouseOutboundOrder(ctx *app.Context) {
	param := &model.ReqStorehouseOutboundOrder{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	err := c.StorehouseService.CreateStorehouseOutboundOrder(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("创建出库单成功")
}

// GetStorehouseOutboundOrderList
func (c *CustomerController) GetStorehouseOutboundOrderList(ctx *app.Context) {
	param := &model.ReqStorehouseOutboundOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	list, err := c.StorehouseService.GetStorehouseOutboundOrderList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(list)
}

// GetStorehouseOutboundOrderInfo
func (c *CustomerController) GetStorehouseOutboundOrderInfo(ctx *app.Context) {
	param := &model.ReqUuidParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	info, err := c.StorehouseService.GetStorehouseOutboundOrderInfo(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(info)
}

// GetRemittanceBillList
func (c *CustomerController) GetRemittanceBillList(ctx *app.Context) {

	param := &model.ReqRemittanceBillQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if param.TeamUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "团队UUID不能为空")
		return
	}

	list, err := c.RemittanceBillService.GetRemittanceBillList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(list)
}
