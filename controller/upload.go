package controller

import (
	"net/http"
	"os"
	"sgin/model"
	"sgin/pkg/app"
)

type UploadController struct {
}

// 文件上传
// @Summary 文件上传
// @Tags 上传
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "token"
// @Param file formData file true "文件"
// @Success 200 {string} app.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/upload [post]
func (u *UploadController) UploadFile(ctx *app.Context) {
	// Multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.Logger.Error("上传文件失败:", err)
		ctx.JSONError(http.StatusBadRequest, "上传文件失败")
		return
	}

	for _, files := range form.File {
		for _, file := range files {
			file.Filename = ctx.Config.Upload.Dir + "/" + file.Filename
			if err := ctx.SaveUploadedFile(file, file.Filename); err != nil {
				ctx.Logger.Error("上传文件失败:", err)
				ctx.JSONError(http.StatusBadRequest, "上传文件失败")
				return
			}
		}
	}

	ctx.JSONSuccess("上传文件成功")
}

// 删除文件
// @Summary 删除文件
// @Tags 上传
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body model.ReqUuidParam true "文件路径"
// @Success 200 {string} app.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/upload/delete [post]
func (u *UploadController) DeleteFile(ctx *app.Context) {
	var param model.ReqFileDeleteParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	err := os.Remove(ctx.Config.Upload.Dir + param.Filename)
	if err != nil {
		ctx.Logger.Error("删除文件失败:", err)
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess("删除文件成功")
}
