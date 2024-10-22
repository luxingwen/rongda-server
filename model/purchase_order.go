package model

type PurchaseOrderReq struct {
	OrderNo string `json:"order_no" gorm:"comment:'采购单号'"` // 采购单号

	EntrustOrderId string `json:"entrust_order_id" gorm:"comment:'委托订单ID'"` // 委托订单ID
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
	Quantity float64 `json:"quantity" gorm:"comment:'数量'"` // 数量

	BoxNum float64 `json:"box_num" gorm:"comment:'箱数'"` // 箱数
	// 价格
	Price float64 `json:"price" gorm:"comment:'价格'"` // 价格
	// 总金额
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"` // 总金额

	// PI箱数
	PIBoxNum float64 `json:"pi_box_num" gorm:"comment:'PI箱数'"` // PI箱数

	// PI数量
	PIQuantity float64 `json:"pi_quantity" gorm:"comment:'PI数量'"` // PI数量

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
	CIBoxNum float64 `json:"ci_box_num" gorm:"comment:'CI箱数'"` // CI箱数

	// CI数量
	CIQuantity float64 `json:"ci_quantity" gorm:"comment:'CI数量'"` // CI数量

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

	// 预付款手续费
	PrepaymentFee float64 `json:"prepayment_fee" gorm:"comment:'预付款手续费'"` // 预付款手续费
	// 尾款手续费
	ResidualFee float64 `json:"residual_fee" gorm:"comment:'尾款手续费'"` // 尾款手续费

	Desc string `json:"desc" gorm:"comment:'备注'"` // 备注
}

const (
	OrderTypeFutures = "1"
	OrderTypeSpot    = "2"

	PurchaseOrderStatusPending = "待处理" // 待处理
	// 处理中
	PurchaseOrderStatusProcessing = "处理中"
	PurchaseOrderStatusHandled    = "已处理" // 已处理

	// 待确认
	PurchaseOrderStatusToBeConfirmed = "待确认"
	// 已审核
	PurchaseOrderStatusReviewed = "已审核"
	PurchaseOrderStatusCanceled = "已取消" // 已取消
	PurchaseOrderStatusDone     = "已完成" // 已完成
	PurchaseOrderStatusInStore  = "已入库" // 已入库
)

// PurchaseOrder 采购单
type PurchaseOrder struct {
	ID             uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`      // 主键ID
	Title          string `json:"title" gorm:"comment:'标题'"`                // 标题
	OrderNo        string `json:"order_no" gorm:"comment:'采购单号'"`           // 采购单号
	EntrustOrderId string `json:"entrust_order_id" gorm:"comment:'委托订单ID'"` // 委托订单ID
	// 订单类型
	OrderType    string `json:"order_type" gorm:"comment:'订单类型'"`                           // 订单类型 1:期货订单 2:现货订单
	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID
	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"`  // 客户UUID
	// 采购日期
	Date string `json:"date" gorm:"comment:'采购日期'"` // 采购日期
	// 采购单号

	// 消费使用单位
	ConsumerUnit string `json:"consumer_unit" gorm:"comment:'消费使用单位'"` // 消费使用单位

	// 境内收货人
	DomesticConsignee string `json:"domestic_consignee" gorm:"comment:'境内收货人'"` // 境内收货人

	// 贸易条款
	TradeTerms string `json:"trade_terms" gorm:"comment:'贸易条款'"` // 贸易条款

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

	// 是否海关放行
	IsCustomsClearance string `json:"is_customs_clearance" gorm:"comment:'是否海关放行'"` // 是否海关放行

	// 实际到港日期
	ActualArrivalDate string `json:"actual_arrival_date" gorm:"comment:'实际到港日期'"` // 实际到港日期

	// 预计入库仓库
	EstimatedWarehouse string `json:"estimated_warehouse" gorm:"comment:'预计入库仓库'"` // 预计入库仓库

	// 实际入库仓库
	ActualWarehouse string `json:"actual_warehouse" gorm:"comment:'实际入库仓库'"` // 实际入库仓库

	// 定金金额
	DepositAmount float64 `json:"deposit_amount" gorm:"comment:'定金金额'"` // 定金金额

	// 定金比例
	DepositRatio float64 `json:"deposit_ratio" gorm:"comment:'定金比例'"` // 定金比例

	// CI总金额
	CITotalAmount float64 `json:"ci_total_amount" gorm:"comment:'CI总金额'"` // CI总金额

	// 尾款金额
	ResidualAmount float64 `json:"residual_amount" gorm:"comment:'尾款金额'"` // 尾款金额

	// PI总金额
	PITotalAmount float64 `json:"pi_total_amount" gorm:"comment:'PI总金额'"` // PI总金额

	// 采购人
	Purchaser string `json:"purchaser" gorm:"comment:'采购人'"` // 采购人

	// 更新人
	Updater string `json:"updater" gorm:"comment:'更新人'"` // 更新人

	// 附件
	Attachment string `json:"attachment" gorm:"comment:'附件'"` // 附件

	// 采购单状态
	Status string `json:"status" gorm:"comment:'采购单状态'"` // 采购单状态 1： 待处理 2：已处理 3：已取消 4：已完成 5: 已入库

	// 船公司
	ShipCompany    string `json:"ship_company" gorm:"comment:'船公司'"`      // 船公司
	CabinetNo      string `json:"cabinet_no" gorm:"comment:'柜号'"`         // 柜号
	BillOfLadingNo string `json:"bill_of_lading_no" gorm:"comment:'提单号'"` // 提单号
	ShipName       string `json:"ship_name" gorm:"comment:'船名'"`          // 船名
	Voyage         string `json:"voyage" gorm:"comment:'航次'"`             // 航次
	// 封条号
	SealNo string `json:"seal_no" gorm:"comment:'封条号'"` // 封条号
	// 柜型
	CabinetType string `json:"cabinet_type" gorm:"comment:'柜型'"` // 柜型

	// 预计到港日期
	EstimatedArrivalDate string `json:"estimated_arrival_date" gorm:"comment:'预计到港日期'"` // 预计到港日期
	// 起运港
	DeparturePort string `json:"departure_port" gorm:"comment:'起运港'"` // 起运港

	// 目的港
	DestinationPort string `json:"destination_port" gorm:"comment:'目的港'"` // 目的港

	// 实际到港日期
	ActualArrivalPort string `json:"actual_arrival_port" gorm:"comment:'实际到港日期'"` // 实际到港日期

	// 形式发票/采购订单附件
	InvoiceAttachment     string `json:"invoice_attachment" gorm:"comment:'形式发票/采购订单附件'"`        // 形式发票/采购订单附件
	InvoiceAttachmentTime string `json:"invoice_attachment_time" gorm:"comment:'形式发票/采购订单附件时间'"` // 形式发票/采购订单附件时间
	// 商业发票
	CommercialInvoice     string `json:"commercial_invoice" gorm:"comment:'商业发票'"`        // 商业发票
	CommercialInvoiceTime string `json:"commercial_invoice_time" gorm:"comment:'商业发票时间'"` // 商业发票时间
	// 装箱单
	PackingList     string `json:"packing_list" gorm:"comment:'装箱单'"`        // 装箱单
	PackingListTime string `json:"packing_list_time" gorm:"comment:'装箱单时间'"` // 装箱单时间
	// 船公司提单
	BillOfLading     string `json:"bill_of_lading" gorm:"comment:'船公司提单'"`        // 船公司提单
	BillOfLadingTime string `json:"bill_of_lading_time" gorm:"comment:'船公司提单时间'"` // 船公司提单时间
	// 批次单
	BatchOrder     string `json:"batch_order" gorm:"comment:'批次单'"`        // 批次单
	BatchOrderTime string `json:"batch_order_time" gorm:"comment:'批次单时间'"` // 批次单时间
	// 卫生证
	SanitaryCertificate     string `json:"sanitary_certificate" gorm:"comment:'卫生证'"`        // 卫生证
	SanitaryCertificateTime string `json:"sanitary_certificate_time" gorm:"comment:'卫生证时间'"` // 卫生证时间
	// 产地证
	CertificateOfOrigin     string `json:"certificate_of_origin" gorm:"comment:'产地证'"`        // 产地证
	CertificateOfOriginTime string `json:"certificate_of_origin_time" gorm:"comment:'产地证时间'"` // 产地证时间
	// 报关单
	CustomsDeclaration     string `json:"customs_declaration" gorm:"comment:'报关单'"`        // 报关单
	CustomsDeclarationTime string `json:"customs_declaration_time" gorm:"comment:'报关单时间'"` // 报关单时间
	// 检疫证
	QuarantineCertificate     string `json:"quarantine_certificate" gorm:"comment:'检疫证'"`        // 检疫证
	QuarantineCertificateTime string `json:"quarantine_certificate_time" gorm:"comment:'检疫证时间'"` // 检疫证时间

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type PurchaseOrderReceiptFileInfo struct {
	ID        uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`             // 主键ID
	OrderNo   string `json:"order_no" gorm:"comment:'采购单号'"`                  // 采购单号
	Key       string `json:"key" gorm:"comment:'修改字段'"`                       // 修改字段
	Filename  string `json:"filename" gorm:"comment:'文件名'"`                   // 文件名
	FilePath  string `json:"file_path" gorm:"comment:'文件路径'"`                 // 文件路径
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
	Quantity               float64 `json:"quantity" gorm:"comment:'数量'"`                                           // 数量
	BoxNum                 float64 `json:"box_num" gorm:"comment:'箱数'"`                                            // 箱数
	Price                  float64 `json:"price" gorm:"comment:'价格'"`                                              // 价格
	TotalAmount            float64 `json:"total_amount" gorm:"comment:'总金额'"`                                      // 总金额
	PIBoxNum               float64 `json:"pi_box_num" gorm:"comment:'PI箱数'"`                                       // PI箱数
	PIQuantity             float64 `json:"pi_quantity" gorm:"comment:'PI数量'"`                                      // PI数量
	PIUnitPrice            float64 `json:"pi_unit_price" gorm:"comment:'PI单价'"`                                    // PI单价
	PITotalAmount          float64 `json:"pi_total_amount" gorm:"comment:'PI总金额'"`                                 // PI总金额
	CabinetNo              string  `json:"cabinet_no" gorm:"comment:'柜号'"`                                         // 柜号
	BillOfLadingNo         string  `json:"bill_of_lading_no" gorm:"comment:'提单号'"`                                 // 提单号
	ShipName               string  `json:"ship_name" gorm:"comment:'船名'"`                                          // 船名
	Voyage                 string  `json:"voyage" gorm:"comment:'航次'"`                                             // 航次
	CIInvoiceNo            string  `json:"ci_invoice_no" gorm:"comment:'CI发票号'"`                                   // CI发票号
	CIBoxNum               float64 `json:"ci_box_num" gorm:"comment:'CI箱数'"`                                       // CI箱数
	CIQuantity             float64 `json:"ci_quantity" gorm:"comment:'CI数量'"`                                      // CI数量
	CIUnitPrice            float64 `json:"ci_unit_price" gorm:"comment:'CI单价'"`                                    // CI单价
	CITotalAmount          float64 `json:"ci_total_amount" gorm:"comment:'CI总金额'"`                                 // CI总金额

	// CI尾款金额
	CIResidualAmount float64 `json:"ci_residual_amount" gorm:"comment:'CI尾款金额'"` // CI尾款金额
	//

	ProductionDate       string `json:"production_date" gorm:"comment:'生产日期'"`          // 生产日期
	EstimatedArrivalDate string `json:"estimated_arrival_date" gorm:"comment:'预计到港日期'"` // 预计到港日期

	// RMB定金金额
	RMBDepositAmount float64 `json:"rmb_deposit_amount" gorm:"comment:'RMB定金金额'"` // RMB定金金额
	// RMB定金金额时间
	RMBDepositAmountTime string `json:"rmb_deposit_amount_time" gorm:"comment:'RMB定金金额时间'"` // RMB定金金额时间

	// 预付款手续费
	PrepaymentFee float64 `json:"prepayment_fee" gorm:"comment:'预付款手续费'"` // 预付款手续费
	// 尾款手续费
	ResidualFee float64 `json:"residual_fee" gorm:"comment:'尾款手续费'"` // 尾款手续费

	// RMB尾款金额
	RMBResidualAmount float64 `json:"rmb_residual_amount" gorm:"comment:'RMB尾款金额'"` // RMB尾款金额
	// RMB尾款金额时间
	RMBResidualAmountTime string `json:"rmb_residual_amount_time" gorm:"comment:'RMB尾款金额时间'"` // RMB尾款金额时间
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

type ReqPurchaseOrderUpdateItem struct {
	OrderNo string      `json:"order_no" gorm:"comment:'采购单号'"` // 采购单号
	Key     string      `json:"key" gorm:"comment:'修改字段'"`      // 修改字段
	Value   interface{} `json:"value" gorm:"comment:'修改值'"`     // 修改值
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

// 货运信息
type FreightInfo struct {
	CabinetNo      string `json:"cabinet_no" gorm:"comment:'柜号'"`         // 柜号
	BillOfLadingNo string `json:"bill_of_lading_no" gorm:"comment:'提单号'"` // 提单号
	ShipName       string `json:"ship_name" gorm:"comment:'船名'"`          // 船名
	Voyage         string `json:"voyage" gorm:"comment:'航次'"`             // 航次
}

type ReqPurchaseOrderDeleteReceiptFile struct {
	OrderNo string `json:"order_no" gorm:"comment:'采购单号'"` // 采购单号
	Key     string `json:"key" gorm:"comment:'修改字段'"`      // 修改字段
}
