package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type ProductCategoryController struct {
	ProductCategoryService *service.ProductCategoryService
}

// @Summary 创建商品类别
// @Description 创建商品类别
// @Tags 商品类别
// @Accept  json
// @Produce  json
// @Param param body model.ProductCategory true "商品类别参数"
// @Success 200 {object} model.ProductCategory
// @Router /api/v1/product-category/create [post]
func (p *ProductCategoryController) CreateProductCategory(ctx *app.Context) {
	var param model.ProductCategory
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductCategoryService.CreateProductCategory(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 获取商品类别信息
// @Description 获取商品类别信息
// @Tags 商品类别
// @Accept  json
// @Produce  json
// @Param uuid path string true "商品类别UUID"
// @Success 200 {object} model.ProductCategory
// @Router /api/v1/product-category/info/{uuid} [get]
func (p *ProductCategoryController) GetProductCategoryInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	category, err := p.ProductCategoryService.GetProductCategoryByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(category)
}

// @Summary 更新商品类别
// @Description 更新商品类别
// @Tags 商品类别
// @Accept  json
// @Produce  json
// @Param param body model.ProductCategory true "商品类别参数"
// @Success 200 {object} model.ProductCategory
// @Router /api/v1/product-category/update [post]
func (p *ProductCategoryController) UpdateProductCategory(ctx *app.Context) {
	var param model.ProductCategory
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.ProductCategoryService.UpdateProductCategory(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除商品类别
// @Description 删除商品类别
// @Tags 商品类别
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "商品类别UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/product-category/delete [post]
func (p *ProductCategoryController) DeleteProductCategory(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	if err := p.ProductCategoryService.DeleteProductCategory(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取商品类别列表
// @Description 获取商品类别列表
// @Tags 商品类别
// @Accept  json
// @Produce  json
// @Param param body model.ReqProductCategoryQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/product-category/list [post]
func (p *ProductCategoryController) GetProductCategoryList(ctx *app.Context) {
	param := &model.ReqProductCategoryQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	categories, err := p.ProductCategoryService.GetProductCategoryList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(categories)
}

// @Summary 获取所有商品类别
// @Description 获取所有商品类别
// @Tags 商品类别
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/product-category/all [get]
func (p *ProductCategoryController) GetAllProductCategories(ctx *app.Context) {
	categories, err := p.ProductCategoryService.GetProductCategoryAll(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(categories)
}
