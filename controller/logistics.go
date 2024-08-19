package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type LogisticsController struct {
	LogisticsService *service.LogisticsService
}

// @Summary 创建物流
// @Description 创建物流
// @Tags 物流
// @Accept  json
// @Produce  json
// @Param param body model.Logistics true "物流参数"
// @Success 200 {object} model.Logistics
// @Router /api/v1/logistics/create [post]
func (l *LogisticsController) CreateLogistics(ctx *app.Context) {
	var param model.Logistics
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := l.LogisticsService.CreateLogistics(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新物流
// @Description 更新物流
// @Tags 物流
// @Accept  json
// @Produce  json
// @Param param body model.Logistics true "物流参数"
// @Success 200 {object} model.Logistics
// @Router /api/v1/logistics/update [post]
func (l *LogisticsController) UpdateLogistics(ctx *app.Context) {
	var param model.Logistics
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := l.LogisticsService.UpdateLogistics(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除物流
// @Description 删除物流
// @Tags 物流
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "物流UUID"
// @Success 200 {string} string	"ok"
// @Router /api/v1/logistics/delete [post]
func (l *LogisticsController) DeleteLogistics(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := l.LogisticsService.DeleteLogistics(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取物流信息
// @Description 获取物流信息
// @Tags 物流
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "物流UUID"
// @Success 200 {object} model.Logistics
// @Router /api/v1/logistics/info [post]
func (l *LogisticsController) GetLogisticsInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	logistics, err := l.LogisticsService.GetLogisticsByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(logistics)
}

// @Summary 获取物流列表
// @Description 获取物流列表
// @Tags 物流
// @Accept  json
// @Produce  json
// @Param param body model.ReqLogisticsQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/logistics/list [post]
func (l *LogisticsController) GetLogisticsList(ctx *app.Context) {
	param := &model.ReqLogisticsQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	logisticsList, err := l.LogisticsService.GetLogisticsList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(logisticsList)
}
