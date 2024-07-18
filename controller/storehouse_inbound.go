package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type StorehouseInboundController struct {
	InboundService *service.StorehouseInboundService
}

// @Summary 创建入库
// @Description 创建入库
// @Tags 入库
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseInboundReq true "入库参数"
// @Success 200 {object} model.StorehouseInbound
// @Router /api/v1/inbound/create [post]
func (p *StorehouseInboundController) CreateInbound(ctx *app.Context) {
	var param model.StorehouseInboundReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	param.Status = model.StorehouseInboundStatusPending
	if err := p.InboundService.CreateInbound(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("入库创建成功")
}

// @Summary 获取入库信息
// @Description 获取入库信息
// @Tags 入库
// @Accept  json
// @Produce  json
// @Param uuid path string true "入库UUID"
// @Success 200 {object} model.StorehouseInbound
// @Router /api/v1/inbound/info/{uuid} [get]
func (p *StorehouseInboundController) GetInboundInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	inbound, err := p.InboundService.GetInbound(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(inbound)
}

// @Summary 获取入库明细
// @Description 获取入库明细
// @Tags 入库
// @Accept  json
// @Produce  json
// @Param uuid path string true "入库UUID"
// @Success 200 {object} model.StorehouseInboundDetail
// @Router /api/v1/inbound/detail/{uuid} [get]
func (p *StorehouseInboundController) GetInboundDetail(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	inboundDetail, err := p.InboundService.GetInboundDetail(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(inboundDetail)
}

// @Summary 更新入库
// @Description 更新入库
// @Tags 入库
// @Accept  json
// @Produce  json
// @Param param body model.StorehouseInboundReq true "入库参数"
// @Success 200 {object} model.StorehouseInbound
// @Router /api/v1/inbound/update [post]
func (p *StorehouseInboundController) UpdateInbound(ctx *app.Context) {
	var param model.StorehouseInboundUpdateReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.InboundService.UpdateInbound(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("入库更新成功")
}

// @Summary 删除入库
// @Description 删除入库
// @Tags 入库
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "入库UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/inbound/delete [post]
func (p *StorehouseInboundController) DeleteInbound(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.InboundService.DeleteInbound(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("入库删除成功")
}

// @Summary 获取入库列表
// @Description 获取入库列表
// @Tags 入库
// @Accept  json
// @Produce  json
// @Param param body model.ReqStorehouseInboundQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/inbound/list [post]
func (p *StorehouseInboundController) GetInboundList(ctx *app.Context) {
	param := &model.ReqStorehouseInboundQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	inbounds, err := p.InboundService.ListInbounds(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(inbounds)
}
