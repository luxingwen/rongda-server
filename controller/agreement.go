package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
	"time"

	"github.com/google/uuid"
)

type AgreementController struct {
	AgreementService *service.AgreementService
}

// @Summary 创建合同
// @Description 创建合同
// @Tags 合同
// @Accept  json
// @Produce  json
// @Param param body model.Agreement true "合同参数"
// @Success 200 {object} model.Agreement
// @Router /api/v1/agreement/create [post]
func (t *AgreementController) CreateAgreement(ctx *app.Context) {
	// Parse the multipart form
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil { // limit your max input length!
		ctx.JSONError(http.StatusBadRequest, "File is too big")
		return
	}

	// Read form fields
	param := model.Agreement{
		Uuid:    uuid.New().String(),
		Date:    ctx.Request.FormValue("date"),
		Content: ctx.Request.FormValue("content"),
		Type:    model.AgreementTypePurchase,
	}

	// Handle file uploads
	form := ctx.Request.MultipartForm
	files := form.File["attachment"]
	var attachments []string

	for _, fileHeader := range files {

		// Create a unique filename and save the file
		filename := "/agreement/" + uuid.New().String() + filepath.Ext(fileHeader.Filename)

		err := ctx.SaveUploadedFile(fileHeader, ctx.Config.Upload.Dir+filename)
		if err != nil {
			ctx.JSONError(http.StatusInternalServerError, "Cannot save file")
			return
		}
		attachments = append(attachments, filename)
	}

	battachments, _ := json.Marshal(attachments)

	param.Attachment = string(battachments)
	param.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	param.UpdatedAt = param.CreatedAt

	// Get user ID from context
	userid := ctx.GetString("user_id")

	if err := t.AgreementService.CreateAgreement(ctx, userid, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 获取合同信息
// @Description 获取合同信息
// @Tags 合同
// @Accept  json
// @Produce  json
// @Param uuid path string true "合同UUID"
// @Success 200 {object} model.Agreement
// @Router /api/v1/agreement/{uuid} [get]
func (t *AgreementController) GetAgreement(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	agreement, err := t.AgreementService.GetAgreement(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(agreement)
}

// @Summary 更新合同
// @Description 更新合同
// @Tags 合同
// @Accept  json
// @Produce  json
// @Param param body model.Agreement true "合同参数"
// @Success 200 {object} model.Agreement
// @Router /api/v1/agreement/update [post]
func (t *AgreementController) UpdateAgreement(ctx *app.Context) {
	var param model.Agreement
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.AgreementService.UpdateAgreement(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(param)
}

// @Summary 删除合同
// @Description 删除合同
// @Tags 合同
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "合同UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/agreement/delete [post]
func (t *AgreementController) DeleteAgreement(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	agreement, err := t.AgreementService.GetAgreement(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	// Delete attachments
	if agreement.Attachment != "" {
		var attachments []string
		if err := json.Unmarshal([]byte(agreement.Attachment), &attachments); err != nil {
			ctx.JSONError(http.StatusInternalServerError, err.Error())
			return
		}

		for _, attachment := range attachments {
			// 删除文件
			os.Remove(ctx.Config.Upload.Dir + attachment)
		}
	}

	if err := t.AgreementService.DeleteAgreement(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess("ok")
}

// @Summary 获取合同列表
// @Description 获取合同列表
// @Tags 合同
// @Accept  json
// @Produce  json
// @Param param body model.ReqAgreementQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/agreement/list [post]
func (t *AgreementController) ListAgreements(ctx *app.Context) {
	param := &model.ReqAgreementQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	agreements, err := t.AgreementService.ListAgreements(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(agreements)
}
