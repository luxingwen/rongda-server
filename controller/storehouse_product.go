package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type StorehouseProductController struct {
	ProductService *service.StorehouseProductService
}

// @Summary 创建仓库物品
// @Description 创建仓库物品
// @Tags 仓库物品
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseProduct true "仓库物品参数"
// @Success 200 {object} model.StorehouseProduct
// @Router /api/v1/storehouse_product/create [post]
func (p *StorehouseProductController) CreateProduct(ctx *app.Context) {
	var param model.StorehouseProduct
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userid := ctx.GetString("user_id")
	if err := p.ProductService.CreateProduct(ctx, userid, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("仓库物品创建成功")
}

// @Summary 获取仓库物品信息
// @Description 获取仓库物品信息
// @Tags 仓库物品
// @Accept  json
// @Produce  json
// @Param id path uint true "仓库物品ID"
// @Success 200 {object} model.StorehouseProduct
// @Router /api/v1/storehouse_product/info/{id} [get]
func (p *StorehouseProductController) GetProductInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	product, err := p.ProductService.GetProduct(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(product)
}

// @Summary 获取物品库存操作记录
// @Description 获取物品库存操作记录
// @Tags 仓库物品
// @Accept  json
// @Produce  json
// @Param param body model.ReqStorehouseProductOpLogQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/storehouse_product/op_log [post]
func (p *StorehouseProductController) GetProductOpLog(ctx *app.Context) {
	param := &model.ReqStorehouseProductOpLogQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	opLogs, err := p.ProductService.GetProductOpLog(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(opLogs)
}

// @Summary 更新仓库物品
// @Description 更新仓库物品
// @Tags 仓库物品
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseProduct true "仓库物品参数"
// @Success 200 {object} model.StorehouseProduct
// @Router /api/v1/storehouse_product/update [post]
func (p *StorehouseProductController) UpdateProduct(ctx *app.Context) {
	var param model.StorehouseProduct
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	ctx.Logger.Info("user_id:", userId)
	if err := p.ProductService.UpdateProduct(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("仓库物品更新成功")
}

// @Summary 删除仓库物品
// @Description 删除仓库物品
// @Tags 仓库物品
// @Accept  json
// @Produce  json
// @Param param body model.ReqIDParam true "仓库物品ID"
// @Success 200 {string} string "ok"
// @Router /api/v1/storehouse_product/delete [post]
func (p *StorehouseProductController) DeleteProduct(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductService.DeleteProduct(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("仓库物品删除成功")
}

// @Summary 获取仓库物品列表
// @Description 获取仓库物品列表
// @Tags 仓库物品
// @Accept  json
// @Produce  json
// @Param param body model.ReqStorehouseProductQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/storehouse_product/list [post]
func (p *StorehouseProductController) GetProductList(ctx *app.Context) {
	param := &model.ReqStorehouseProductQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	products, err := p.ProductService.ListProducts(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(products)
}
