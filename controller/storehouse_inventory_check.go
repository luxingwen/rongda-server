package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type StorehouseInventoryCheckController struct {
	InventoryCheckService *service.StorehouseInventoryCheckService
}

func NewStorehouseInventoryCheckController() *StorehouseInventoryCheckController {
	return &StorehouseInventoryCheckController{
		InventoryCheckService: service.NewStorehouseInventoryCheckService(),
	}
}

// @Summary 创建盘点
// @Description 创建盘点
// @Tags 盘点
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseInventoryCheckReq true "盘点参数"
// @Success 200 {object} model.StorehouseInventoryCheck
// @Router /api/v1/storehouse_inventory_check/create [post]
func (t *StorehouseInventoryCheckController) CreateInventoryCheck(ctx *app.Context) {
	var param model.StorehouseInventoryCheckReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.InventoryCheckService.CreateInventoryCheck(ctx, ctx.GetString("user_id"), &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取盘点
// @Description 获取盘点
// @Tags 盘点
// @Accept  json
// @Produce  json
// @Param uuid path string true "盘点UUID"
// @Success 200 {object} model.StorehouseInventoryCheck
// @Router /api/v1/storehouse_inventory_check/{uuid} [get]
func (t *StorehouseInventoryCheckController) GetInventoryCheck(ctx *app.Context) {
	uuid := ctx.Param("uuid")
	check, err := t.InventoryCheckService.GetInventoryCheck(ctx, uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(check)
}

// @Summary 更新盘点
// @Description 更新盘点
// @Tags 盘点
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseInventoryCheck true "盘点参数"
// @Success 200 {object} model.StorehouseInventoryCheck
// @Router /api/v1/storehouse_inventory_check/update [post]
func (t *StorehouseInventoryCheckController) UpdateInventoryCheck(ctx *app.Context) {
	var param model.StorehouseInventoryCheck
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.InventoryCheckService.UpdateInventoryCheck(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除盘点
// @Description 删除盘点
// @Tags 盘点
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "盘点UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/storehouse_inventory_check/delete [post]
func (t *StorehouseInventoryCheckController) DeleteInventoryCheck(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.InventoryCheckService.DeleteInventoryCheck(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取盘点列表
// @Description 获取盘点列表
// @Tags 盘点
// @Accept  json
// @Produce  json
// @Param param body model.ReqInventoryCheckQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/storehouse_inventory_check/list [post]
func (t *StorehouseInventoryCheckController) GetInventoryCheckList(ctx *app.Context) {
	param := &model.ReqInventoryCheckQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	checks, err := t.InventoryCheckService.ListInventoryChecks(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(checks)
}
