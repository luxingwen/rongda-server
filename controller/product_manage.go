package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type ProductManageController struct {
	ProductManageService *service.ProductManageService
	SkuService           *service.SkuService
}

// @Summary 创建商品
// @Description 创建商品
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param param body model.ProductManage true "商品参数"
// @Success 200 {object} model.ProductManage
// @Router /api/v1/product_manage/create [post]
func (p *ProductManageController) CreateProduct(ctx *app.Context) {
	var param model.ProductManage
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductManageService.CreateProduct(ctx, &param); err != nil {
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
// @Success 200 {object} model.ProductManage
// @Router /api/v1/product_manage/info/{uuid} [get]
func (p *ProductManageController) GetProductInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	product, err := p.ProductManageService.GetProductByUUID(ctx, param.Uuid)
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
// @Param param body model.ProductManage true "商品参数"
// @Success 200 {object} model.ProductManage
// @Router /api/v1/product_manage/update [post]
func (p *ProductManageController) UpdateProduct(ctx *app.Context) {
	var param model.ProductManage
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductManageService.UpdateProduct(ctx, &param); err != nil {
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
// @Router /api/v1/product_manage/delete [post]
func (p *ProductManageController) DeleteProduct(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if err := p.ProductManageService.DeleteProduct(ctx, param.Uuid); err != nil {
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
// @Router /api/v1/product_manage/list [post]
func (p *ProductManageController) GetProductList(ctx *app.Context) {
	param := &model.ReqProductQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	products, err := p.ProductManageService.GetProductList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(products)
}

// @Summary 获取所有商品
// @Description 获取所有商品
// @Tags 商品
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/product_manage/all [get]
func (p *ProductManageController) GetAllProduct(ctx *app.Context) {
	products, err := p.ProductManageService.GetAvailableProductList(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(products)
}

// @Summary 获取商品SKU列表
// @Description 获取商品SKU列表
// @Tags 商品
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "商品UUID"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/product_manage/sku/list [post]
func (p *ProductManageController) GetProductSkuList(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	skulist, err := p.SkuService.GetSkuListByProductUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(skulist)
}
