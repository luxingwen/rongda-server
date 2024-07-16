package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SalesOutOfStockController struct {
	SalesOutOfStockService *service.SalesOutOfStockService
}

func NewSalesOutOfStockController() *SalesOutOfStockController {
	return &SalesOutOfStockController{
		SalesOutOfStockService: service.NewSalesOutOfStockService(),
	}
}

// @Summary 创建销售出库单
// @Description 创建销售出库单
// @Tags 销售出库
// @Accept  json
// @Produce  json
// @Param param body model.SalesOutOfStockReq true "销售出库单参数"
// @Success 200 {object} model.SalesOutOfStock
// @Router /api/v1/sales_out_of_stock/create [post]
func (t *SalesOutOfStockController) CreateSalesOutOfStock(ctx *app.Context) {
	var param model.SalesOutOfStockReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOutOfStockService.CreateSalesOutOfStock(ctx, ctx.GetString("userId"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取销售出库单
// @Description 获取销售出库单
// @Tags 销售出库
// @Accept  json
// @Produce  json
// @Param uuid path string true "出库单UUID"
// @Success 200 {object} model.SalesOutOfStock
// @Router /api/v1/sales_out_of_stock/{uuid} [get]
func (t *SalesOutOfStockController) GetSalesOutOfStock(ctx *app.Context) {
	uuid := ctx.Param("uuid")
	outOfStock, err := t.SalesOutOfStockService.GetSalesOutOfStock(ctx, uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(outOfStock)
}

// @Summary 更新销售出库单
// @Description 更新销售出库单
// @Tags 销售出库
// @Accept  json
// @Produce  json
// @Param param body model.SalesOutOfStock true "销售出库单参数"
// @Success 200 {object} model.SalesOutOfStock
// @Router /api/v1/sales_out_of_stock/update [post]
func (t *SalesOutOfStockController) UpdateSalesOutOfStock(ctx *app.Context) {
	var param model.SalesOutOfStock
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOutOfStockService.UpdateSalesOutOfStock(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除销售出库单
// @Description 删除销售出库单
// @Tags 销售出库
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "出库单UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/sales_out_of_stock/delete [post]
func (t *SalesOutOfStockController) DeleteSalesOutOfStock(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.SalesOutOfStockService.DeleteSalesOutOfStock(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取销售出库单列表
// @Description 获取销售出库单列表
// @Tags 销售出库
// @Accept  json
// @Produce  json
// @Param param body model.ReqSalesOutOfStockQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/sales_out_of_stock/list [post]
func (t *SalesOutOfStockController) GetSalesOutOfStockList(ctx *app.Context) {
	param := &model.ReqSalesOutOfStockQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	outOfStocks, err := t.SalesOutOfStockService.ListSalesOutOfStocks(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(outOfStocks)
}
