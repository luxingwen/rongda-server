package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type StorehouseController struct {
	StorehouseService *service.StorehouseService
}

// @Summary 创建仓库
// @Description 创建仓库
// @Tags 仓库
// @Accept  json
// @Produce  json
// @Param param body model.Storehouse true "仓库参数"
// @Success 200 {object} model.Storehouse
// @Router /api/v1/storehouse/create [post]
func (p *StorehouseController) CreateStorehouse(ctx *app.Context) {
	var param model.Storehouse
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.StorehouseService.CreateStorehouse(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 获取仓库信息
// @Description 获取仓库信息
// @Tags 仓库
// @Accept  json
// @Produce  json
// @Param uuid path string true "仓库UUID"
// @Success 200 {object} model.Storehouse
// @Router /api/v1/storehouse/info/{uuid} [get]
func (p *StorehouseController) GetStorehouseInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	storehouse, err := p.StorehouseService.GetStorehouseByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(storehouse)
}

// @Summary 更新仓库
// @Description 更新仓库
// @Tags 仓库
// @Accept  json
// @Produce  json
// @Param param body model.Storehouse true "仓库参数"
// @Success 200 {object} model.Storehouse
// @Router /api/v1/storehouse/update [post]
func (p *StorehouseController) UpdateStorehouse(ctx *app.Context) {
	var param model.Storehouse
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.StorehouseService.UpdateStorehouse(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除仓库
// @Description 删除仓库
// @Tags 仓库
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "仓库UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/storehouse/delete [post]
func (p *StorehouseController) DeleteStorehouse(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.StorehouseService.DeleteStorehouse(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取仓库列表
// @Description 获取仓库列表
// @Tags 仓库
// @Accept  json
// @Produce  json
// @Param param body model.ReqStorehouseQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/storehouse/list [post]
func (p *StorehouseController) GetStorehouseList(ctx *app.Context) {
	param := &model.ReqStorehouseQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	storehouses, err := p.StorehouseService.GetStorehouseList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(storehouses)
}

// @Summary 获取所有可用仓库列表
// @Description 获取所有可用仓库列表
// @Tags 仓库
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Storehouse
// @Router /api/v1/storehouse/all [get]
func (p *StorehouseController) GetAllStorehouse(ctx *app.Context) {
	storehouses, err := p.StorehouseService.GetAvailableStorehouses(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(storehouses)
}

// UpdateStorehouseItem
func (p *StorehouseController) UpdateStorehouseItem(ctx *app.Context) {
	var param model.ReqStorehouseUpdateItem
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.StorehouseService.UpdateStorehouseItem(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// UpdateStorehouseItemByMap
func (p *StorehouseController) UpdateStorehouseItemByMap(ctx *app.Context) {
	mdata := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&mdata); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	storehouseUuid := mdata["storehouse_uuid"].(string)
	if storehouseUuid == "" {
		ctx.JSONError(http.StatusBadRequest, "storehouse_uuid is required")
		return
	}
	delete(mdata, "storehouse_uuid")
	if err := p.StorehouseService.UpdateStorehouseItemByMap(ctx, storehouseUuid, mdata); err != nil {

		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}
