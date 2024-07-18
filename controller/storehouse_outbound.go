package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type StorehouseOutboundController struct {
	OutboundService *service.StorehouseOutboundService
}

// @Summary 创建出库信息
// @Description 创建出库信息
// @Tags 仓库出库
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseOutboundReq true "出库信息参数"
// @Success 200 {object} model.StorehouseOutbound
// @Router /api/v1/storehouse_outbound/create [post]
func (p *StorehouseOutboundController) CreateOutbound(ctx *app.Context) {
	var param model.StorehouseOutboundReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	param.Status = model.StorehouseOutboundStatusPending
	if err := p.OutboundService.CreateOutbound(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("出库信息创建成功")
}

// @Summary 获取出库信息
// @Description 获取出库信息
// @Tags 仓库出库
// @Accept  json
// @Produce  json
// @Param id path uint true "出库ID"
// @Success 200 {object} model.StorehouseOutbound
// @Router /api/v1/storehouse_outbound/info/{id} [get]
func (p *StorehouseOutboundController) GetOutboundInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	outbound, err := p.OutboundService.GetOutbound(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(outbound)
}

// @Summary 获取出库明细
// @Description 获取出库明细
// @Tags 仓库出库
// @Accept  json
// @Produce  json
// @Param id path uint true "出库ID"
// @Success 200 {object} model.StorehouseOutboundDetail
// @Router /api/v1/storehouse_outbound/detail/{id} [get]
func (p *StorehouseOutboundController) GetOutboundDetail(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	details, err := p.OutboundService.GetOutboundDetail(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(details)
}

// @Summary 更新出库信息
// @Description 更新出库信息
// @Tags 仓库出库
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseOutbound true "出库信息参数"
// @Success 200 {object} model.StorehouseOutbound
// @Router /api/v1/storehouse_outbound/update [post]
func (p *StorehouseOutboundController) UpdateOutbound(ctx *app.Context) {
	var param model.StorehouseOutbound
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.OutboundService.UpdateOutbound(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("出库信息更新成功")
}

// @Summary 删除出库信息
// @Description 删除出库信息
// @Tags 仓库出库
// @Accept  json
// @Produce  json
// @Param param body model.ReqIDParam true "出库ID"
// @Success 200 {string} string "ok"
// @Router /api/v1/storehouse_outbound/delete [post]
func (p *StorehouseOutboundController) DeleteOutbound(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.OutboundService.DeleteOutbound(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("出库信息删除成功")
}

// @Summary 获取出库信息列表
// @Description 获取出库信息列表
// @Tags 仓库出库
// @Accept  json
// @Produce  json
// @Param param body model.ReqStorehouseOutboundQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/storehouse_outbound/list [post]
func (p *StorehouseOutboundController) GetOutboundList(ctx *app.Context) {
	param := &model.ReqStorehouseOutboundQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	outbounds, err := p.OutboundService.ListOutbounds(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(outbounds)
}
