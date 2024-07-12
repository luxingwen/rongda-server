package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SupplierController struct {
	SupplierService *service.SupplierService
}

// @Summary 创建供应商
// @Description 创建供应商
// @Tags 供应商
// @Accept  json
// @Produce  json
// @Param param body model.Supplier true "供应商参数"
// @Success 200 {object} model.Supplier
// @Router /api/v1/supplier/create [post]
func (s *SupplierController) CreateSupplier(ctx *app.Context) {
	var param model.Supplier
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SupplierService.CreateSupplier(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新供应商
// @Description 更新供应商
// @Tags 供应商
// @Accept  json
// @Produce  json
// @Param param body model.Supplier true "供应商参数"
// @Success 200 {object} model.Supplier
// @Router /api/v1/supplier/update [post]
func (s *SupplierController) UpdateSupplier(ctx *app.Context) {
	var param model.Supplier
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SupplierService.UpdateSupplier(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除供应商
// @Description 删除供应商
// @Tags 供应商
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "供应商UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/supplier/delete [post]
func (s *SupplierController) DeleteSupplier(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SupplierService.DeleteSupplier(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取供应商信息
// @Description 获取供应商信息
// @Tags 供应商
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "供应商UUID"
// @Success 200 {object} model.Supplier
// @Router /api/v1/supplier/info [post]
func (s *SupplierController) GetSupplierInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	supplier, err := s.SupplierService.GetSupplierByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(supplier)
}

// @Summary 获取供应商列表
// @Description 获取供应商列表
// @Tags 供应商
// @Accept  json
// @Produce  json
// @Param param body model.ReqSupplierQueryParam true "查询参数"
// @Success 200 {object} model.SupplierQueryResponse
// @Router /api/v1/supplier/list [post]
func (s *SupplierController) GetSupplierList(ctx *app.Context) {
	param := &model.ReqSupplierQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	suppliers, err := s.SupplierService.GetSupplierList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(suppliers)
}
