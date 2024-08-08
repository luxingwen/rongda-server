package controller

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SalesOrderController struct {
	SalesOrderService  *service.SalesOrderService
	AgreementService   *service.AgreementService
	StepService        *service.StepService
	PaymentBillService *service.PaymentBillService
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
	var param model.SalesOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}
	if err := t.SalesOrderService.UpdateSalesOrder(ctx, userId, &param); err != nil {
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

// CreateSalesAgreement
func (t *SalesOrderController) CreateSalesAgreement(ctx *app.Context) {
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}

	title := ctx.PostForm("title")

	signature_position_list := ctx.PostForm("signature_position_list")

	if title == "" {
		ctx.JSONError(http.StatusBadRequest, "title is required")
		return
	}

	orderNo := ctx.PostForm("order_no")
	if orderNo == "" {
		ctx.JSONError(http.StatusBadRequest, "order_no is required")
		return
	}

	if signature_position_list == "" {
		ctx.JSONError(http.StatusBadRequest, "signature_position_list is required")
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Logger.Info("file", file.Filename)
	// 保存合同

	extfile := filepath.Ext(file.Filename)

	filename := "/sales_order/agreement/" + "sales_" + orderNo + extfile

	err = ctx.SaveUploadedFile(file, ctx.Config.Upload.Dir+filename)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	fileAttachment := model.FileAttachment{
		Name: file.Filename,
		Url:  filename,
	}

	ctx.Logger.Info("title", title)
	ctx.Logger.Info("signature_position_list", signature_position_list)

	srcfilebyte, _ := json.Marshal(fileAttachment)

	agreement := model.Agreement{
		Title:             title,
		OrderNo:           orderNo,
		SourceFile:        string(srcfilebyte),
		SignaturePosition: signature_position_list,
		Creater:           userId,
		Type:              model.AgreementTypeSales,
	}

	err = t.AgreementService.CreateAgreement(ctx, userId, &agreement)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(nil)
}

// CreateDepositAgreement
func (t *SalesOrderController) CreateDepositAgreement(ctx *app.Context) {
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}

	title := ctx.PostForm("title")

	signature_position_list := ctx.PostForm("signature_position_list")

	if title == "" {
		ctx.JSONError(http.StatusBadRequest, "title is required")
		return
	}

	orderNo := ctx.PostForm("order_no")
	if orderNo == "" {
		ctx.JSONError(http.StatusBadRequest, "order_no is required")
		return
	}

	if signature_position_list == "" {
		ctx.JSONError(http.StatusBadRequest, "signature_position_list is required")
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Logger.Info("file", file.Filename)
	// 保存合同

	extfile := filepath.Ext(file.Filename)

	filename := "/sales_order/agreement/" + "deposit_" + orderNo + extfile

	err = ctx.SaveUploadedFile(file, ctx.Config.Upload.Dir+filename)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	fileAttachment := model.FileAttachment{
		Name: file.Filename,
		Url:  filename,
	}

	ctx.Logger.Info("title", title)
	ctx.Logger.Info("signature_position_list", signature_position_list)

	srcfilebyte, _ := json.Marshal(fileAttachment)

	agreement := model.Agreement{
		Title:             title,
		OrderNo:           orderNo,
		SourceFile:        string(srcfilebyte),
		SignaturePosition: signature_position_list,
		Creater:           userId,
		Type:              model.AgreementTypeSalesDeposit,
	}

	err = t.AgreementService.CreateAgreement(ctx, userId, &agreement)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(nil)
}

// CreateFinalAgreement
func (t *SalesOrderController) CreateFinalAgreement(ctx *app.Context) {
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}

	title := ctx.PostForm("title")

	signature_position_list := ctx.PostForm("signature_position_list")

	if title == "" {
		ctx.JSONError(http.StatusBadRequest, "title is required")
		return
	}

	orderNo := ctx.PostForm("order_no")
	if orderNo == "" {
		ctx.JSONError(http.StatusBadRequest, "order_no is required")
		return
	}

	if signature_position_list == "" {
		ctx.JSONError(http.StatusBadRequest, "signature_position_list is required")
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Logger.Info("file", file.Filename)
	// 保存合同

	extfile := filepath.Ext(file.Filename)

	filename := "/sales_order/agreement/" + "final_" + orderNo + extfile

	err = ctx.SaveUploadedFile(file, ctx.Config.Upload.Dir+filename)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	fileAttachment := model.FileAttachment{
		Name: file.Filename,
		Url:  filename,
	}

	ctx.Logger.Info("title", title)
	ctx.Logger.Info("signature_position_list", signature_position_list)

	srcfilebyte, _ := json.Marshal(fileAttachment)

	agreement := model.Agreement{
		Title:             title,
		OrderNo:           orderNo,
		SourceFile:        string(srcfilebyte),
		SignaturePosition: signature_position_list,
		Creater:           userId,
		Type:              model.AgreementTypeSalesFinalPayment,
	}

	err = t.AgreementService.CreateAgreement(ctx, userId, &agreement)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(nil)
}

// GetSalesOrderStepList
func (t *SalesOrderController) GetSalesOrderStepList(ctx *app.Context) {
	params := &model.ReqUuidParam{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	steps, err := t.StepService.GetStepListByRefTypeAndRefID(ctx, model.StepChainRefTypeSalesOrder, params.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(steps)
}

// ConfirmSalesOrder
func (t *SalesOrderController) ConfirmSalesOrder(ctx *app.Context) {
	var param model.ReqSalesOrderConfirmParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOrderService.ConfirmSalesOrder(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// CreatePaymentBill
func (t *SalesOrderController) CreateDepositPaymentBill(ctx *app.Context) {
	var param model.PaymentBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	param.Type = model.PaymentBillTypeDeposit
	param.Status = model.PaymentBillStatusPendingPayment
	if err := t.SalesOrderService.CreateDepositPaymentBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// CreateFinalPaymentBill
func (t *SalesOrderController) CreateFinalPaymentBill(ctx *app.Context) {
	var param model.PaymentBill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	param.Type = model.PaymentBillTypeFinal
	param.Status = model.PaymentBillStatusPendingPayment
	if err := t.SalesOrderService.CreateFinalPaymentBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}
