package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
)

type SysBankInfoController struct {
	SysBankInfoService *service.SysBankInfoService
}

// @Summary 创建银行信息
// @Description 创建银行信息
// @Tags 银行信息
// @Accept  json
// @Produce  json
// @Param param body model.SysBankInfo true "银行信息参数"
// @Success 200 {object} model.SysBankInfo
// @Router /api/v1/bankinfo/create [post]
func (s *SysBankInfoController) CreateSysBankInfo(ctx *app.Context) {
	var param model.SysBankInfo
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SysBankInfoService.CreateSysBankInfo(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 更新银行信息
// @Description 更新银行信息
// @Tags 银行信息
// @Accept  json
// @Produce  json
// @Param param body model.SysBankInfo true "银行信息参数"
// @Success 200 {object} model.SysBankInfo
// @Router /api/v1/bankinfo/update [post]
func (s *SysBankInfoController) UpdateSysBankInfo(ctx *app.Context) {
	var param model.SysBankInfo
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SysBankInfoService.UpdateSysBankInfo(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除银行信息
// @Description 删除银行信息
// @Tags 银行信息
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "银行信息UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/bankinfo/delete [post]
func (s *SysBankInfoController) DeleteSysBankInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := s.SysBankInfoService.DeleteSysBankInfo(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取银行信息
// @Description 获取银行信息
// @Tags 银行信息
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "银行信息UUID"
// @Success 200 {object} model.SysBankInfo
// @Router /api/v1/bankinfo/info [post]
func (s *SysBankInfoController) GetSysBankInfo(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	bankInfo, err := s.SysBankInfoService.GetSysBankInfoByUUID(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(bankInfo)
}

// @Summary 获取银行信息列表
// @Description 获取银行信息列表
// @Tags 银行信息
// @Accept  json
// @Produce  json
// @Param param body model.ReqSysBankInfoQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/bankinfo/list [post]
func (s *SysBankInfoController) GetSysBankInfoList(ctx *app.Context) {
	param := &model.ReqSysBankInfoQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	bankInfos, err := s.SysBankInfoService.GetSysBankInfoList(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(bankInfos)
}

func (s *SysBankInfoController) GetAvailableSysBankInfoList(ctx *app.Context) {

	bankInfos, err := s.SysBankInfoService.GetAvailableSysBankInfoList(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(bankInfos)
}
