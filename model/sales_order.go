package model

type SalesOrderReq struct {
	OrderNo            string  `json:"order_no" form:"order_no" binding:"-"`                                // 订单号
	OrderType          string  `json:"order_type" form:"order_type" binding:"required"`                     // 订单类型：1.期货订单 2.现货订单
	Title              string  `json:"title" form:"title" binding:"required"`                               // 标题
	OrderDate          string  `json:"order_date" form:"order_date" binding:"required"`                     // 订单日期
	Deposit            int     `json:"deposit" form:"deposit" binding:"required"`                           // 定金
	DepositRatio       float64 `json:"deposit_ratio" form:"deposit_ratio" binding:"required"`               // 定金比例
	FinalPaymentAmount float64 `json:"final_payment_amount" form:"final_payment_amount" binding:"required"` // 尾款金额

	PurchaseOrderNo string `json:"purchase_order_no" form:"purchase_order_no" binding:"-"` // 采购订单号
	EntrustOrderId  string `json:"entrust_order_id" gorm:"comment:'委托订单ID'"`               // 委托订单ID

	OrderAmount        int    `json:"order_amount" form:"order_amount" binding:"required"`               // 订单金额
	SettlementCurrency string `json:"settlement_currency" form:"settlement_currency" binding:"required"` // 结算币种
	// 备注
	Remarks string `json:"remarks" form:"remarks" binding:"-"`
	// 客户uuid
	CustomerUuid string `json:"customer_uuid" form:"customer_uuid" binding:"required"`

	// 销售人
	Salesman string `json:"salesman" form:"salesman" binding:"-"`

	// 商品列表
	ProductList []SalesOrderItemReq `json:"product_list" form:"product_list" binding:"required"`
}

type SalesOrderItemReq struct {
	PurchaseOrderProductNo string `json:"purchase_order_product_no" form:"purchase_order_product_no" binding:"-"`
	// 商品uuid

	ProductUuid string `json:"product_uuid" form:"product_uuid" binding:"required"`

	SkuUuid string `json:"sku_uuid" form:"sku_uuid" binding:"required"`

	// 商品数量
	ProductQuantity int `json:"product_quantity" form:"product_quantity" binding:"required"`
	// 商品单价
	ProductPrice float64 `json:"product_price" form:"product_price" binding:"required"`
	// 商品金额
	ProductAmount float64 `json:"product_amount" form:"product_amount" binding:"required"`

	// 箱数
	BoxNum int `json:"box_num" form:"box_num" binding:"-"`
}

const (
	// 订单状态
	// 待处理
	OrderStatusPending = "待处理"
	// 处理中
	OrderStatusProcessing     = "处理中"
	OrderStatusPendingPayment = "待支付"
	OrderStatusPaid           = "已支付"
	OrderStatusPendingShipped = "待发货"
	OrderStatusShipped        = "已发货"
	OrderStatusCompleted      = "已完成"
	OrderStatusCancelled      = "已取消"
)

// 销售订单
type SalesOrder struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	OrderNo         string  `json:"order_no" gorm:"uniqueIndex;type:varchar(50);comment:订单号"`
	PurchaseOrderNo string  `json:"purchase_order_no" gorm:"type:varchar(50);comment:采购订单号"` // 采购订单号
	EntrustOrderId  string  `json:"entrust_order_id" gorm:"comment:'委托订单ID'"`                // 委托订单ID
	AgreementUuid   string  `json:"agreement_uuid" gorm:"type:varchar(50);comment:合同uuid"`   // 合同uuid
	Title           string  `json:"title" gorm:"type:varchar(200);comment:标题"`               // 标题
	OrderType       string  `json:"order_type" gorm:"type:varchar(50);comment:订单类型"`         // 订单类型：1.期货订单 2.现货订单
	OrderStatus     string  `json:"order_status" gorm:"type:varchar(50);comment:订单状态"`       // 订单状态：待支付、已支付、已发货、已完成、已取消
	AgreementNo     string  `json:"agreement_no" gorm:"type:varchar(50);comment:合同号"`        // 合同号
	OrderDate       string  `json:"order_date" gorm:"type:varchar(50);comment:订单日期"`         // 订单日期
	DepositAmount   float64 `json:"deposit_amount" gorm:"comment:定金"`
	// 定金比例
	DepositRatio float64 `json:"deposit_ratio" gorm:"comment:定金比例"`
	// 尾款金额
	FinalPaymentAmount float64 `json:"final_payment_amount" gorm:"comment:尾款金额"`

	OrderAmount  float64 `json:"order_amount" gorm:"comment:订单金额"`
	Salesman     string  `json:"salesman" gorm:"type:varchar(50);comment:销售人"`
	CustomerUuid string  `json:"customer_uuid" gorm:"type:varchar(50);comment:客户uuid"`
	// 结算币种
	SettlementCurrency string `json:"settlement_currency" gorm:"type:varchar(50);comment:结算币种"`
	Remarks            string `json:"remarks" gorm:"type:varchar(50);comment:备注"`
	Updater            string `json:"updater" gorm:"type:varchar(50);comment:更新人"`
	CreatedAt          string `json:"created_at" gorm:"autoCreateTime"` // CreatedAt 记录了创建的时间
	UpdatedAt          string `json:"updated_at" gorm:"autoUpdateTime"` // UpdatedAt 记录了最后更新的时间
}

type CustomerSalesOrderRes struct {
	SalesOrder
	PurchaseOrderInfo *PurchaseOrder `json:"purchase_order_info"`
}

type SalesOrderRes struct {
	SalesOrder
	CustomerInfo *Customer `json:"customer_info"`
	SalesmanInfo *User     `json:"salesman_info"`
}

// 销售订单明细
type SalesOrderItem struct {
	ID                     uint    `json:"id" gorm:"primaryKey"`
	Uuid                   string  `json:"uuid" gorm:"uniqueIndex;type:varchar(50);comment:uuid"`
	OrderNo                string  `json:"order_no" gorm:"type:varchar(50);comment:订单号"`
	PurchaseOrderProductNo string  `json:"purchase_order_product_no" gorm:"type:varchar(50);comment:采购订单商品号"` // 采购订单商品号
	ProductUuid            string  `json:"product_uuid" gorm:"type:varchar(50);comment:商品uuid"`
	ProductName            string  `json:"product_name" gorm:"type:varchar(50);comment:商品名称"`
	SkuUuid                string  `json:"sku_uuid" gorm:"type:varchar(50);comment:sku uuid"`
	SkuName                string  `json:"sku_name" gorm:"type:varchar(50);comment:sku名称"`
	ProductQuantity        int     `json:"product_quantity" gorm:"comment:商品数量"`
	BoxNum                 int     `json:"box_num" gorm:"comment:箱数"`
	ProductPrice           float64 `json:"product_price" gorm:"comment:商品单价"`
	ProductAmount          float64 `json:"product_amount" gorm:"comment:商品金额"`
	CreatedAt              string  `json:"created_at" gorm:"autoCreateTime"` // CreatedAt 记录了创建的时间
	UpdatedAt              string  `json:"updated_at" gorm:"autoUpdateTime"` // UpdatedAt 记录了最后更新的时间
}

type SalesOrderItemRes struct {
	SalesOrderItem
	ProductInfo *Product `json:"product"`
	SkuInfo     *Sku     `json:"sku"`
}
