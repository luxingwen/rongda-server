package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
	"strings"

	"github.com/google/uuid"
)

type OrderFileController struct {
	OrderFileService *service.OrderFileService
}

// 上传销售订单文件
func (o *OrderFileController) UploadSalesOrderFile(ctx *app.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orderNo := ctx.PostForm("order_no")

	if orderNo == "" {
		ctx.JSONError(http.StatusBadRequest, "order_no is required")
		return
	}

	extfile := filepath.Ext(file.Filename)

	filename := "/sales_order/receipt_file/" + orderNo + "_" + uuid.New().String() + extfile

	err = ctx.SaveUploadedFile(file, ctx.Config.Upload.Dir+filename)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	orderfile := model.OrderFile{
		OrderNo:   orderNo,
		Name:      strings.TrimSuffix(file.Filename, extfile),
		Filename:  file.Filename,
		Url:       filename,
		OrderType: "sales_order",
	}

	if err := o.OrderFileService.CreateOrderFile(ctx, &orderfile); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orderfile)
}

// 删除订单文件
func (o *OrderFileController) DeleteOrderFile(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orderFileInfo, err := o.OrderFileService.GetOrderFileByUuid(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	os.Remove(ctx.Config.Upload.Dir + orderFileInfo.Url)

	if err := o.OrderFileService.DeleteOrderFile(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(nil)
}

// 获取销售订单文件列表
func (o *OrderFileController) GetOrderFileList(ctx *app.Context) {
	var param model.ReqOrderNoParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orderFiles, err := o.OrderFileService.GetOrderFileListByOrderNo(ctx, param.OrderNo)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orderFiles)
}
