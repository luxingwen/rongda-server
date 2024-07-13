package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type DepartmentController struct {
	DepartmentService *service.DepartmentService
}

// @Summary 创建部门
// @Description 创建部门
// @Tags 部门
// @Accept  json
// @Produce  json
// @Param param body model.Department true "部门参数"
// @Success 200 {object} model.Department
// @Router /api/v1/department/create [post]
func (d *DepartmentController) CreateDepartment(ctx *app.Context) {
	var param model.Department
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := d.DepartmentService.CreateDepartment(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新部门
// @Description 更新部门
// @Tags 部门
// @Accept  json
// @Produce  json
// @Param param body model.Department true "部门参数"
// @Success 200 {object} model.Department
// @Router /api/v1/department/update [post]
func (d *DepartmentController) UpdateDepartment(ctx *app.Context) {
	var param model.Department
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := d.DepartmentService.UpdateDepartment(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除部门
// @Description 删除部门
// @Tags 部门
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "部门UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/department/delete [post]
func (d *DepartmentController) DeleteDepartment(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := d.DepartmentService.DeleteDepartment(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取部门信息
// @Description 获取部门信息
// @Tags 部门
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "部门UUID"
// @Success 200 {object} model.Department
// @Router /api/v1/department/info [post]
func (d *DepartmentController) GetDepartmentInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	department, err := d.DepartmentService.GetDepartmentByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(department)
}

// @Summary 获取部门列表
// @Description 获取部门列表
// @Tags 部门
// @Accept  json
// @Produce  json
// @Param param body model.ReqDepartmentQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/department/list [post]
func (d *DepartmentController) GetDepartmentList(ctx *app.Context) {
	param := &model.ReqDepartmentQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	departments, err := d.DepartmentService.GetDepartmentList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(departments)
}
