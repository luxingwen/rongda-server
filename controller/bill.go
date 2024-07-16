package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type BillController struct {
	BillService *service.BillService
}

func NewBillController() *BillController {
	return &BillController{
		BillService: service.NewBillService(),
	}
}

// @Summary 创建发票
// @Description 创建发票
// @Tags 发票
// @Accept  json
// @Produce  json
// @Param param body model.Bill true "发票参数"
// @Success 200 {object} model.Bill
// @Router /api/v1/bill/create [post]
func (t *BillController) CreateBill(ctx *app.Context) {
	var param model.Bill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.BillService.CreateBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 获取发票
// @Description 获取发票
// @Tags 发票
// @Accept  json
// @Produce  json
// @Param uuid path string true "发票UUID"
// @Success 200 {object} model.Bill
// @Router /api/v1/bill/{uuid} [get]
func (t *BillController) GetBill(ctx *app.Context) {
	uuid := ctx.Param("uuid")
	bill, err := t.BillService.GetBill(ctx, uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(bill)
}

// @Summary 更新发票
// @Description 更新发票
// @Tags 发票
// @Accept  json
// @Produce  json
// @Param param body model.Bill true "发票参数"
// @Success 200 {object} model.Bill
// @Router /api/v1/bill/update [post]
func (t *BillController) UpdateBill(ctx *app.Context) {
	var param model.Bill
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.BillService.UpdateBill(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除发票
// @Description 删除发票
// @Tags 发票
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "发票UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/bill/delete [post]
func (t *BillController) DeleteBill(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.BillService.DeleteBill(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取发票列表
// @Description 获取发票列表
// @Tags 发票
// @Accept  json
// @Produce  json
// @Param param body model.ReqBillQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/bill/list [post]
func (t *BillController) GetBillList(ctx *app.Context) {
	param := &model.ReqBillQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	bills, err := t.BillService.ListBills(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(bills)
}
