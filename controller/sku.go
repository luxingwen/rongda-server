package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SkuController struct {
	SkuService *service.SkuService
}

// @Summary 创建SKU
// @Description 创建SKU
// @Tags SKU
// @Accept  json
// @Produce  json
// @Param param body model.Sku true "SKU参数"
// @Success 200 {object} model.SkuInfoResponse
// @Router /api/v1/sku/create [post]
func (s *SkuController) CreateSku(ctx *app.Context) {
	var param model.SkuReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SkuService.CreateSku(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新SKU
// @Description 更新SKU
// @Tags SKU
// @Accept  json
// @Produce  json
// @Param param body model.Sku true "SKU参数"
// @Success 200 {object} model.SkuInfoResponse
// @Router /api/v1/sku/update [post]
func (s *SkuController) UpdateSku(ctx *app.Context) {
	var param model.Sku
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SkuService.UpdateSku(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除SKU
// @Description 删除SKU
// @Tags SKU
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "SKU UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/sku/delete [post]
func (s *SkuController) DeleteSku(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SkuService.DeleteSku(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取SKU信息
// @Description 获取SKU信息
// @Tags SKU
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "SKU UUID"
// @Success 200 {object} model.SkuInfoResponse
// @Router /api/v1/sku/info [post]
func (s *SkuController) GetSkuInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	sku, err := s.SkuService.GetSkuByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(sku)
}

// @Summary 获取SKU列表
// @Description 获取SKU列表
// @Tags SKU
// @Accept  json
// @Produce  json
// @Param param body model.ReqSkuQueryParam true "查询参数"
// @Success 200 {object} model.SkuQueryResponse
// @Router /api/v1/sku/list [post]
func (s *SkuController) GetSkuList(ctx *app.Context) {
	param := &model.ReqSkuQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	skus, err := s.SkuService.GetSkuList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(skus)
}
