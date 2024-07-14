package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type ProductController struct {
	ProductService *service.ProductService
}

// @Summary 创建商品
// @Description 创建商品
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param param body model.Product true "商品参数"
// @Success 200 {object} model.Product
// @Router /api/v1/product/create [post]
func (p *ProductController) CreateProduct(ctx *app.Context) {
	var param model.Product
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductService.CreateProduct(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 获取商品信息
// @Description 获取商品信息
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param uuid path string true "商品UUID"
// @Success 200 {object} model.Product
// @Router /api/v1/product/info/{uuid} [get]
func (p *ProductController) GetProductInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	product, err := p.ProductService.GetProductByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(product)
}

// @Summary 更新商品
// @Description 更新商品
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param param body model.Product true "商品参数"
// @Success 200 {object} model.Product
// @Router /api/v1/product/update [post]
func (p *ProductController) UpdateProduct(ctx *app.Context) {
	var param model.Product
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductService.UpdateProduct(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除商品
// @Description 删除商品
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "商品UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/product/delete [post]
func (p *ProductController) DeleteProduct(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductService.DeleteProduct(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取商品列表
// @Description 获取商品列表
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param param body model.ReqProductQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/product/list [post]
func (p *ProductController) GetProductList(ctx *app.Context) {
	param := &model.ReqProductQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	products, err := p.ProductService.GetProductList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(products)
}
