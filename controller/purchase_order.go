package controller

import (
	"net/http"
	"sgin/model"
	"sgin/pkg/app"
	"sgin/service"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type PurchaseOrderController struct {
	PurchaseOrderService *service.PurchaseOrderService
}

func NewPurchaseOrderController() *PurchaseOrderController {
	return &PurchaseOrderController{
		PurchaseOrderService: service.NewPurchaseOrderService(),
	}
}

// @Summary 创建采购单
// @Description 创建采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseOrderReq true "采购单参数"
// @Success 200 {object} model.PurchaseOrder
// @Router /api/v1/purchase_order/create [post]
func (t *PurchaseOrderController) CreatePurchaseOrderFutures(ctx *app.Context) {
	var param model.PurchaseOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}
	if err := t.PurchaseOrderService.CreatePurchaseOrderFutures(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

func (t *PurchaseOrderController) CreatePurchaseOrderSpot(ctx *app.Context) {
	var param model.PurchaseOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}
	if err := t.PurchaseOrderService.CreatePurchaseOrderSpot(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// 上传excel文件
func (t *PurchaseOrderController) UploadFuturesItemsExcel(ctx *app.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	f, err := file.Open()
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	// 读取excel文件
	xlsx, err := excelize.OpenReader(f)
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	// 获取第一个工作表
	sheetName := xlsx.GetSheetName(0)
	// 获取列名
	headers, err := xlsx.GetCols(sheetName)

	if err != nil {
		ctx.Logger.Error("读取excel文件失败", err)
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	columnIndexMap := make(map[string]int)
	for colIndex, colValues := range headers {
		if len(colValues) > 0 {
			columnIndexMap[colValues[0]] = colIndex
		}
	}

	// 读取数据
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		ctx.Logger.Error("读取excel文件失败", err)
		return
	}

	var purchaseOrderItems []model.PurchaseOrderItemReq

	for rowIndex, row := range rows {
		// Skip the header row
		if rowIndex == 0 {
			continue
		}
		var item model.PurchaseOrderItemReq

		item.ProductName = safeGetCellString(row, columnIndexMap, "产品名称")
		item.SkuCode = safeGetCellString(row, columnIndexMap, "SKU编码")
		item.SkuSpec = safeGetCellString(row, columnIndexMap, "SKU规格")

		item.Quantity, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "数量"))
		if err != nil {
			ctx.Logger.Error("数量转换失败", err)
		}

		item.Price, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "价格"), 64)
		if err != nil {
			ctx.Logger.Error("价格转换失败", err)
		}
		item.TotalAmount, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "总金额"), 64)
		if err != nil {
			ctx.Logger.Error("总金额转换失败", err)
		}
		item.PIBoxNum, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "PI箱数"))
		if err != nil {
			ctx.Logger.Error("PI箱数转换失败", err)
		}

		item.PIQuantity, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "PI数量"))
		if err != nil {
			ctx.Logger.Error("PI数量转换失败", err)
		}
		item.PIUnitPrice, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "PI单价"), 64)
		if err != nil {
			ctx.Logger.Error("PI单价转换失败", err)
		}
		item.PITotalAmount, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "PI总金额"), 64)
		if err != nil {
			ctx.Logger.Error("PI总金额转换失败", err)
		}
		item.CabinetNo = safeGetCellString(row, columnIndexMap, "柜号")
		item.BillOfLadingNo = safeGetCellString(row, columnIndexMap, "提单号")
		item.ShipName = safeGetCellString(row, columnIndexMap, "船名")
		item.Voyage = safeGetCellString(row, columnIndexMap, "航次")
		item.CIInvoiceNo = safeGetCellString(row, columnIndexMap, "CI发票号")
		item.CIBoxNum, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "CI箱数"))
		if err != nil {
			ctx.Logger.Error("CI箱数转换失败", err)
		}
		item.CIQuantity, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "CI数量"))
		if err != nil {
			ctx.Logger.Error("CI数量转换失败", err)
		}
		item.CIUnitPrice, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "CI单价"), 64)
		if err != nil {
			ctx.Logger.Error("CI单价转换失败", err)
		}
		item.CITotalAmount, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "CI总金额"), 64)
		if err != nil {
			ctx.Logger.Error("CI总金额转换失败", err)
		}
		item.ProductionDate = safeGetCellString(row, columnIndexMap, "生产日期")
		item.EstimatedArrivalDate = safeGetCellString(row, columnIndexMap, "预计到港日期")
		item.Tariff, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "关税"), 64)
		if err != nil {
			ctx.Logger.Error("关税转换失败", err)
		}
		item.VAT, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "增值税"), 64)
		if err != nil {
			ctx.Logger.Error("增值税转换失败", err)
		}

		item.PaymentDate = safeGetCellString(row, columnIndexMap, "缴费日期")
		purchaseOrderItems = append(purchaseOrderItems, item)
	}

	rPuchaseOrderItems, err := t.PurchaseOrderService.CompletePurchaseOrderItem(ctx, purchaseOrderItems)
	if err != nil {
		ctx.Logger.Error("Failed to complete purchase order items", err)
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(rPuchaseOrderItems)

}

func safeGetCellString(row []string, columnIndexMap map[string]int, keyname string) string {
	if index, ok := columnIndexMap[keyname]; ok {
		if len(row) > index {
			return row[index]
		}
	}
	return ""
}

func (t *PurchaseOrderController) UploadSpotItemsExcel(ctx *app.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	f, err := file.Open()
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	// 读取excel文件
	xlsx, err := excelize.OpenReader(f)
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	// 获取第一个工作表
	sheetName := xlsx.GetSheetName(0)
	// 获取列名
	headers, err := xlsx.GetCols(sheetName)

	if err != nil {
		ctx.Logger.Error("读取excel文件失败", err)
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	columnIndexMap := make(map[string]int)
	for colIndex, colValues := range headers {
		if len(colValues) > 0 {
			columnIndexMap[colValues[0]] = colIndex
		}
	}

	// 读取数据
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		ctx.Logger.Error("读取excel文件失败", err)
		return
	}

	var purchaseOrderItems []model.PurchaseOrderItemReq

	for rowIndex, row := range rows {
		// Skip the header row
		if rowIndex == 0 {
			continue
		}
		var item model.PurchaseOrderItemReq
		item.ProductName = safeGetCellString(row, columnIndexMap, "产品名称")
		item.SkuCode = safeGetCellString(row, columnIndexMap, "SKU编码")
		item.SkuSpec = safeGetCellString(row, columnIndexMap, "SKU规格")

		item.Quantity, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "数量"))
		if err != nil {
			ctx.Logger.Error("数量转换失败", err)
		}

		item.BoxNum, err = strconv.Atoi(safeGetCellString(row, columnIndexMap, "箱数"))
		if err != nil {
			ctx.Logger.Error("箱数转换失败", err)
		}

		item.Price, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "价格"), 64)
		if err != nil {
			ctx.Logger.Error("价格转换失败", err)
		}

		item.TotalAmount, err = strconv.ParseFloat(safeGetCellString(row, columnIndexMap, "总金额"), 64)
		if err != nil {
			ctx.Logger.Error("总金额转换失败", err)
		}

		item.CabinetNo = safeGetCellString(row, columnIndexMap, "柜号")
		item.ProductionDate = safeGetCellString(row, columnIndexMap, "生产日期")
		item.Desc = safeGetCellString(row, columnIndexMap, "描述")
		purchaseOrderItems = append(purchaseOrderItems, item)
	}

	rPuchaseOrderItems, err := t.PurchaseOrderService.CompletePurchaseOrderItem(ctx, purchaseOrderItems)
	if err != nil {
		ctx.Logger.Error("Failed to complete purchase order items", err)
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(rPuchaseOrderItems)

}

// @Summary 获取采购单
// @Description 获取采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param orderNo path string true "采购单号"
// @Success 200 {object} model.PurchaseOrderRes
// @Router /api/v1/purchase_order/{orderNo} [get]
func (t *PurchaseOrderController) GetPurchaseOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	order, err := t.PurchaseOrderService.GetPurchaseOrder(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(order)
}

// @Summary 获取采购单商品列表
// @Description 获取采购单商品列表
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param orderNo path string true "采购单号"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/purchase_order/{orderNo}/products [get]
func (t *PurchaseOrderController) GetPurchaseOrderProducts(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	products, err := t.PurchaseOrderService.GetPurchaseOrderItems(ctx, param.Uuid)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(products)
}

// @Summary 更新采购单
// @Description 更新采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.PurchaseOrder true "采购单参数"
// @Success 200 {object} model.PurchaseOrder
// @Router /api/v1/purchase_order/update [post]
func (t *PurchaseOrderController) UpdatePurchaseOrder(ctx *app.Context) {
	var param model.PurchaseOrder
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseOrderService.UpdatePurchaseOrder(ctx, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

func (t *PurchaseOrderController) UpdatePurchaseOrderFutures(ctx *app.Context) {
	var param model.PurchaseOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}
	if err := t.PurchaseOrderService.UpdatePurchaseOrderFutures(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

func (t *PurchaseOrderController) UpdatePurchaseOrderSpot(ctx *app.Context) {
	var param model.PurchaseOrderReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		ctx.JSONError(http.StatusUnauthorized, "用户未登录")
		return
	}
	if err := t.PurchaseOrderService.UpdatePurchaseOrderSpot(ctx, userId, &param); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 删除采购单
// @Description 删除采购单
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.ReqUuidParam true "采购单UUID"
// @Success 200 {string} string "ok"
// @Router /api/v1/purchase_order/delete [post]
func (t *PurchaseOrderController) DeletePurchaseOrder(ctx *app.Context) {
	var param model.ReqUuidParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseOrderService.DeletePurchaseOrder(ctx, param.Uuid); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// @Summary 获取采购单列表
// @Description 获取采购单列表
// @Tags 采购单
// @Accept  json
// @Produce  json
// @Param param body model.ReqPurchaseOrderQueryParam true "查询参数"
// @Success 200 {object} model.PagedResponse
// @Router /api/v1/purchase_order/list [post]
func (t *PurchaseOrderController) GetPurchaseOrderList(ctx *app.Context) {
	param := &model.ReqPurchaseOrderQueryParam{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}

	orders, err := t.PurchaseOrderService.ListPurchaseOrders(ctx, param)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orders)
}

func (t *PurchaseOrderController) GetAvailablePurchaseOrderList(ctx *app.Context) {

	orders, err := t.PurchaseOrderService.GetAvailablePurchaseOrderList(ctx)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSONSuccess(orders)
}

// 更改订单状态
func (t *PurchaseOrderController) UpdatePurchaseOrderStatus(ctx *app.Context) {
	var param model.PurchaseOrderStatusReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	if err := t.PurchaseOrderService.UpdatePurchaseOrderStatus(ctx, param.OrderNo, param.Status); err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(nil)
}

// 根据订单状态获取订单列表
func (t *PurchaseOrderController) GetPurchaseOrderByStatus(ctx *app.Context) {
	var param model.ReqPurchaseOrderStatusParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSONError(http.StatusBadRequest, err.Error())
		return
	}
	orders, err := t.PurchaseOrderService.GetPurchaseOrderListByStatus(ctx, param.StatusList)
	if err != nil {
		ctx.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONSuccess(orders)
}
