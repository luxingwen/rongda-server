package model

type PurchaseOrderReq struct {
	OrderNo string `json:"order_no" gorm:"comment:'采购单号'"` // 采购单号
	// 标题
	Title string `json:"title" gorm:"comment:'标题'"`  // 标题
	Date  string `json:"date" gorm:"comment:'采购日期'"` // 采购日期

	Status string `json:"status" gorm:"comment:'采购单状态'"` // 采购单状态

	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"` // 客户UUID

	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID

	PIAgreementNo string `json:"pi_agreement_no" gorm:"comment:'PI合同号'"` // PI合同号

	// 订单币种
	OrderCurrency string `json:"order_currency" gorm:"comment:'订单币种'"` // 订单币种

	// 结算币种
	SettlementCurrency string `json:"settlement_currency" gorm:"comment:'结算币种'"` // 结算币种

	// 起运地
	Departure string `json:"departure" gorm:"comment:'起运地'"` // 起运地

	// 目的地
	Destination string `json:"destination" gorm:"comment:'目的地'"` // 目的地

	// 预计装船日期
	EstimatedShippingDate string `json:"estimated_shipping_date" gorm:"comment:'预计装船日期'"` // 预计装船日期

	// 预计入库仓库
	EstimatedWarehouse string `json:"estimated_warehouse" gorm:"comment:'预计入库仓库'"` // 预计入库仓库

	// 实际入库仓库
	ActualWarehouse string `json:"actual_warehouse" gorm:"comment:'实际入库仓库'"` // 实际入库仓库

	// 定金金额
	DepositAmount float64 `json:"deposit_amount" gorm:"comment:'定金金额'"` // 定金金额

	// 定金比例
	DepositRatio float64 `json:"deposit_ratio" gorm:"comment:'定金比例'"` // 定金比例

	// 附件
	Attachment []FileAttachment `json:"attachment" gorm:"comment:'附件'"` // 附件

	Details []PurchaseOrderItemReq `json:"details" gorm:"comment:'采购单明细'"` // 采购单明细
}

type PurchaseOrderItemReq struct {
	PurchaseOrderProductNo string `json:"purchase_order_product_no" gorm:"type:char(36);index;comment:'采购单产品单号'"` // 采购单产品单号
	ProductUuid            string `json:"product_uuid" gorm:"type:char(36);index;comment:'产品UUID'"`               // 产品UUID
	// SKU UUID
	SkuUuid string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"` // SKU UUID

	// 产品名称
	ProductName string `json:"product_name" gorm:"comment:'产品名称'"` // 产品名称
	// SKU名称
	SkuName string `json:"sku_name" gorm:"comment:'SKU名称'"` // SKU名称

	// SKUCode
	SkuCode string `json:"sku_code" gorm:"comment:'SKU编码'"` // SKU编码

	// SKU规格
	SkuSpec string `json:"sku_spec" gorm:"comment:'SKU规格'"` // SKU规格

	// 数量
	Quantity int `json:"quantity" gorm:"comment:'数量'"` // 数量

	BoxNum int `json:"box_num" gorm:"comment:'箱数'"` // 箱数
	// 价格
	Price float64 `json:"price" gorm:"comment:'价格'"` // 价格
	// 总金额
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"` // 总金额

	// PI箱数
	PIBoxNum int `json:"pi_box_num" gorm:"comment:'PI箱数'"` // PI箱数

	// PI数量
	PIQuantity int `json:"pi_quantity" gorm:"comment:'PI数量'"` // PI数量

	// PI单价
	PIUnitPrice float64 `json:"pi_unit_price" gorm:"comment:'PI单价'"` // PI单价

	// PI总金额
	PITotalAmount float64 `json:"pi_total_amount" gorm:"comment:'PI总金额'"` // PI总金额

	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号

	// 提单号
	BillOfLadingNo string `json:"bill_of_lading_no" gorm:"comment:'提单号'"` // 提单号

	// 船名
	ShipName string `json:"ship_name" gorm:"comment:'船名'"` // 船名

	// 航次
	Voyage string `json:"voyage" gorm:"comment:'航次'"` // 航次

	// CI发票号
	CIInvoiceNo string `json:"ci_invoice_no" gorm:"comment:'CI发票号'"` // CI发票号

	// CI箱数
	CIBoxNum int `json:"ci_box_num" gorm:"comment:'CI箱数'"` // CI箱数

	// CI数量
	CIQuantity int `json:"ci_quantity" gorm:"comment:'CI数量'"` // CI数量

	// CI单价
	CIUnitPrice float64 `json:"ci_unit_price" gorm:"comment:'CI单价'"` // CI单价

	// CI总金额
	CITotalAmount float64 `json:"ci_total_amount" gorm:"comment:'CI总金额'"` // CI总金额

	// CI尾款金额
	CIResidualAmount float64 `json:"ci_residual_amount" gorm:"comment:'CI尾款金额'"` // CI尾款金额

	// 生产日期
	ProductionDate string `json:"production_date" gorm:"comment:'生产日期'"` // 生产日期

	// 预计到港日期
	EstimatedArrivalDate string `json:"estimated_arrival_date" gorm:"comment:'预计到港日期'"` // 预计到港日期

	// RMB定金金额
	RMBDepositAmount float64 `json:"rmb_deposit_amount" gorm:"comment:'RMB定金金额'"` // RMB定金金额
	// RMB尾款金额
	RMBResidualAmount float64 `json:"rmb_residual_amount" gorm:"comment:'RMB尾款金额'"` // RMB尾款金额

	// 定金汇率
	DepositExchangeRate float64 `json:"deposit_exchange_rate" gorm:"comment:'定金汇率'"` // 定金汇率
	// 尾款汇率
	ResidualExchangeRate float64 `json:"residual_exchange_rate" gorm:"comment:'尾款汇率'"` // 尾款汇率

	// 关税
	Tariff float64 `json:"tariff" gorm:"comment:'关税'"` // 关税

	// 增值税
	VAT float64 `json:"vat" gorm:"comment:'增值税'"` // 增值税

	// 缴费日期
	PaymentDate string `json:"payment_date" gorm:"comment:'缴费日期'"` // 缴费日期

	Desc string `json:"desc" gorm:"comment:'备注'"` // 备注
}

const (
	OrderTypeFutures = "1"
	OrderTypeSpot    = "2"

	PurchaseOrderStatusPending = "待处理" // 待处理
	// 处理中
	PurchaseOrderStatusProcessing = "处理中"
	PurchaseOrderStatusHandled    = "已处理" // 已处理
	// 已审核
	PurchaseOrderStatusReviewed = "已审核"
	PurchaseOrderStatusCanceled = "已取消" // 已取消
	PurchaseOrderStatusDone     = "已完成" // 已完成
	PurchaseOrderStatusInStore  = "已入库" // 已入库
)

// PurchaseOrder 采购单
type PurchaseOrder struct {
	ID      uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Title   string `json:"title" gorm:"comment:'标题'"`           // 标题
	OrderNo string `json:"order_no" gorm:"comment:'采购单号'"`      // 采购单号
	// 订单类型
	OrderType    string `json:"order_type" gorm:"comment:'订单类型'"`                           // 订单类型 1:期货订单 2:现货订单
	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID
	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"`  // 客户UUID
	// 采购日期
	Date string `json:"date" gorm:"comment:'采购日期'"` // 采购日期
	// 采购单号

	// PI合同号
	PIAgreementNo string `json:"pi_agreement_no" gorm:"comment:'PI合同号'"` // PI合同号

	// 订单币种
	OrderCurrency string `json:"order_currency" gorm:"comment:'订单币种'"` // 订单币种

	// 结算币种
	SettlementCurrency string `json:"settlement_currency" gorm:"comment:'结算币种'"` // 结算币种

	// 起运地
	Departure string `json:"departure" gorm:"comment:'起运地'"` // 起运地

	// 目的地
	Destination string `json:"destination" gorm:"comment:'目的地'"` // 目的地

	// 预计装船日期
	EstimatedShippingDate string `json:"estimated_shipping_date" gorm:"comment:'预计装船日期'"` // 预计装船日期

	// 预计入库仓库
	EstimatedWarehouse string `json:"estimated_warehouse" gorm:"comment:'预计入库仓库'"` // 预计入库仓库

	// 实际入库仓库
	ActualWarehouse string `json:"actual_warehouse" gorm:"comment:'实际入库仓库'"` // 实际入库仓库

	// 定金金额
	DepositAmount float64 `json:"deposit_amount" gorm:"comment:'定金金额'"` // 定金金额

	// 定金比例
	DepositRatio float64 `json:"deposit_ratio" gorm:"comment:'定金比例'"` // 定金比例

	// 采购人
	Purchaser string `json:"purchaser" gorm:"comment:'采购人'"` // 采购人

	// 更新人
	Updater string `json:"updater" gorm:"comment:'更新人'"` // 更新人

	// 附件
	Attachment string `json:"attachment" gorm:"comment:'附件'"` // 附件

	// 采购单状态
	Status string `json:"status" gorm:"comment:'采购单状态'"` // 采购单状态 1： 待处理 2：已处理 3：已取消 4：已完成 5: 已入库

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// PurchaseOrderItem 采购单明细
type PurchaseOrderItem struct {
	ID                     uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"`                                    // 主键ID
	PurchaseOrderNo        string  `json:"purchase_order_no" gorm:"type:char(36);index;comment:'采购单单号'"`           // 采购单号
	PurchaseOrderProductNo string  `json:"purchase_order_product_no" gorm:"type:char(36);index;comment:'采购单产品单号'"` // 采购单产品单号
	ProductUuid            string  `json:"product_uuid" gorm:"type:char(36);index;comment:'产品UUID'"`               // 产品UUID
	ProductName            string  `json:"product_name" gorm:"comment:'产品名称'"`                                     // 产品名称
	SkuUuid                string  `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`                 // SKU UUID
	SkuName                string  `json:"sku_name" gorm:"comment:'SKU名称'"`                                        // SKU名称
	Quantity               int     `json:"quantity" gorm:"comment:'数量'"`                                           // 数量
	BoxNum                 int     `json:"box_num" gorm:"comment:'箱数'"`                                            // 箱数
	Price                  float64 `json:"price" gorm:"comment:'价格'"`                                              // 价格
	TotalAmount            float64 `json:"total_amount" gorm:"comment:'总金额'"`                                      // 总金额
	PIBoxNum               int     `json:"pi_box_num" gorm:"comment:'PI箱数'"`                                       // PI箱数
	PIQuantity             int     `json:"pi_quantity" gorm:"comment:'PI数量'"`                                      // PI数量
	PIUnitPrice            float64 `json:"pi_unit_price" gorm:"comment:'PI单价'"`                                    // PI单价
	PITotalAmount          float64 `json:"pi_total_amount" gorm:"comment:'PI总金额'"`                                 // PI总金额
	CabinetNo              string  `json:"cabinet_no" gorm:"comment:'柜号'"`                                         // 柜号
	BillOfLadingNo         string  `json:"bill_of_lading_no" gorm:"comment:'提单号'"`                                 // 提单号
	ShipName               string  `json:"ship_name" gorm:"comment:'船名'"`                                          // 船名
	Voyage                 string  `json:"voyage" gorm:"comment:'航次'"`                                             // 航次
	CIInvoiceNo            string  `json:"ci_invoice_no" gorm:"comment:'CI发票号'"`                                   // CI发票号
	CIBoxNum               int     `json:"ci_box_num" gorm:"comment:'CI箱数'"`                                       // CI箱数
	CIQuantity             int     `json:"ci_quantity" gorm:"comment:'CI数量'"`                                      // CI数量
	CIUnitPrice            float64 `json:"ci_unit_price" gorm:"comment:'CI单价'"`                                    // CI单价
	CITotalAmount          float64 `json:"ci_total_amount" gorm:"comment:'CI总金额'"`                                 // CI总金额

	// CI尾款金额
	CIResidualAmount float64 `json:"ci_residual_amount" gorm:"comment:'CI尾款金额'"` // CI尾款金额
	//

	ProductionDate       string `json:"production_date" gorm:"comment:'生产日期'"`          // 生产日期
	EstimatedArrivalDate string `json:"estimated_arrival_date" gorm:"comment:'预计到港日期'"` // 预计到港日期

	// RMB定金金额
	RMBDepositAmount float64 `json:"rmb_deposit_amount" gorm:"comment:'RMB定金金额'"` // RMB定金金额
	// RMB尾款金额
	RMBResidualAmount float64 `json:"rmb_residual_amount" gorm:"comment:'RMB尾款金额'"` // RMB尾款金额

	// 定金汇率
	DepositExchangeRate float64 `json:"deposit_exchange_rate" gorm:"comment:'定金汇率'"` // 定金汇率
	// 尾款汇率
	ResidualExchangeRate float64 `json:"residual_exchange_rate" gorm:"comment:'尾款汇率'"` // 尾款汇率

	Tariff      float64 `json:"tariff" gorm:"comment:'关税'"`                      // 关税
	VAT         float64 `json:"vat" gorm:"comment:'增值税'"`                        // 增值税
	PaymentDate string  `json:"payment_date" gorm:"comment:'缴费日期'"`              // 缴费日期
	Desc        string  `json:"desc" gorm:"comment:'备注'"`                        // 备注
	CreatedAt   string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt   string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// PurchaseOrderResp 采购单响应
type PurchaseOrderRes struct {
	PurchaseOrder
	Supplier               Supplier            `json:"supplier" gorm:"comment:'供应商'"`                  // 供应商
	PurchaserInfo          User                `json:"purchaser_info" gorm:"comment:'采购人'"`            // 采购人
	OrderCurrencyInfo      *SettlementCurrency `json:"order_currency_info" gorm:"comment:'订单币种'"`      // 订单币种
	SettlementCurrencyInfo *SettlementCurrency `json:"settlement_currency_info" gorm:"comment:'结算币种'"` // 结算币种
	// 入库仓库
	EstimatedWarehouseInfo *Storehouse `json:"estimated_warehouse_info" gorm:"comment:'入库仓库'"` // 入库仓库
	CustomerInfo           *Customer   `json:"customer_info" gorm:"comment:'客户'"`              // 客户
	ActualWarehouseInfo    *Storehouse `json:"actual_warehouse_info" gorm:"comment:'实际入库仓库'"`  // 实际入库仓库
}

type PurchaseOrderItemRes struct {
	PurchaseOrderItem
	Product Product `json:"product" gorm:"comment:'产品'"` // 产品
	Sku     Sku     `json:"sku" gorm:"comment:'SKU'"`    // SKU
}
