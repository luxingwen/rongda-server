package model

const (
	StorehouseStatusEnabled  = 1 // 仓库状态启用
	StorehouseStatusDisabled = 2 // 仓库状态未启用
)

const (
	StorehouseProductOpLogOpTypeInbound        = 1 // 入库
	StorehouseProductOpLogOpTypeOutbound       = 2 // 出库
	StorehouseProductOpLogOpTypeInventoryCheck = 3 // 盘点
	StorehouseProductOpLogOpTypeTransfer       = 4 // 调拨
	// 更新库存
	StorehouseProductOpLogOpTypeUpdate = 5 // 更新库存
)

type Storehouse struct {
	ID      uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid    string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	Name    string `json:"name" gorm:"comment:'名称'"`                       // 名称
	Address string `json:"address" gorm:"comment:'地址'"`                    // 地址
	// 联系人
	ContactPerson string `json:"contact_person" gorm:"comment:'联系人'"` // 联系人
	// 联系电话
	ContactPhone string `json:"contact_phone" gorm:"comment:'联系电话'"` // 联系电话
	// 银行账号
	BankAccount string `json:"bank_account" gorm:"comment:'银行账号'"` // 银行账号
	// 开户行
	BankName string `json:"bank_name" gorm:"comment:'开户行'"` // 开户行
	Status   int    `json:"status" gorm:"comment:'状态'"`     // 状态 1:启用 2:未启用
	// 仓库类型
	Type      string `json:"type" gorm:"comment:'仓库类型'"`                      // 仓库类型 1:自有仓库 2:第三方仓库
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	IsDeleted int    `json:"is_deleted" gorm:"comment:'是否删除'"`                // 是否删除 1:删除 0:未删除
}

// 仓库物品信息
type StorehouseProduct struct {
	ID             uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`                         // 主键ID
	Uuid           string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`              // UUID
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` //

	PurchaseOrderNo string `json:"purchase_order_no" gorm:"comment:'采购订单'"` // 采购订单

	PurchaseOrderProductNo string `json:"purchase_order_product_no" gorm:"comment:'采购订单物品编号'"` // 采购订单物品编号
	// 采购订单物品类型
	PurchaseOrderProductType string `json:"purchase_order_product_type" gorm:"comment:'采购订单物品类型'"` // 采购订单物品类型 1：期货 2：现货
	// 客户uuid
	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"` //
	// 物品类型

	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 库存数量
	Quantity float64 `json:"quantity" gorm:"comment:'库存数量'"` // 库存数量
	BoxNum   float64 `json:"box_num" gorm:"comment:'箱数'"`    // 箱数
	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"`                  // 柜号
	InDate    string `json:"in_date" gorm:"comment:'入库日期'"`                   // 入库日期
	MaxDate   string `json:"max_date" gorm:"comment:'最大日期'"`                  // 最大库存日期
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// 仓库物品操作日志
type StorehouseProductOpLog struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	// 采购订单
	PurchaseOrderNo string `json:"purchase_order_no" gorm:"comment:'采购订单'"` // 采购订单
	// 入库单号
	InboundOrderNo        string `json:"inbound_order_no" gorm:"comment:'入库单号'"`            // 入库单号
	InboudItemDdetailUuid string `json:"inboud_item_detail_uuid" gorm:"comment:'入库明细UUID'"` // 入库明细UUID
	// 出库单号
	OutboundOrderNo       string `json:"outbound_order_no" gorm:"comment:'出库单号'"`        // 出库单号
	InventoryCheckOrderNo string `json:"inventory_check_order_no" gorm:"comment:'盘点单号'"` // 盘点单号

	StorehouseUuid        string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"`           //
	StorehouseProductUuid string `json:"storehouse_product_uuid" gorm:"type:char(36);index;comment:'仓库物品UUID'"` // 仓库物品UUID

	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID

	SkuUuid string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"` // SKU UUID

	TeamUuid string `json:"team_uuid" gorm:"type:char(36);index;comment:'团队UUID'"` // 团队UUID

	// 操作之前库存数量
	BeforeQuantity float64 `json:"before_quantity" gorm:"comment:'操作之前库存数量'"` // 操作之前库存数量
	BeforeBoxNum   float64 `json:"before_box_num" gorm:"comment:'操作之前箱数'"`    // 操作之前箱数
	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号
	// 库存数量
	Quantity   float64 `json:"quantity" gorm:"comment:'库存数量'"`                  // 库存数量
	BoxNum     float64 `json:"box_num" gorm:"comment:'箱数'"`                     // 箱数
	LogType    string  `json:"log_type" gorm:"comment:'日志类型'"`                  // 记录类型 采购单入库  销售单入库
	OpType     float64 `json:"op_type" gorm:"comment:'操作类型'"`                   // 操作类型 1:入库 2:出库 3:盘点 4:调拨
	OpQuantity float64 `json:"op_quantity" gorm:"comment:'操作数量'"`               // 操作数量
	OpBoxNum   float64 `json:"op_box_num" gorm:"comment:'操作箱数'"`                // 操作箱数
	OpBy       string  `json:"op_by" gorm:"comment:'操作人'"`                      // 操作人
	OpDesc     string  `json:"op_desc" gorm:"comment:'操作描述'"`                   // 操作描述
	CreatedAt  string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}

type StorehouseProductOpLogRes struct {
	StorehouseProductOpLog
	ProductInfo    Product    `json:"product_info"`
	SkuInfo        Sku        `json:"sku_info"`
	StorehouseInfo Storehouse `json:"storehouse_info"`
	OpByUser       User       `json:"op_by_user"`
}

type StorehouseProductRes struct {
	StorehouseProduct
	Storehouse   Storehouse `json:"storehouse"`
	Product      Product    `json:"product"`
	Sku          Sku        `json:"sku"`
	CustomerInfo Customer   `json:"customer_info"`
	// 库存天数
	StockDays int `json:"stock_days"` // 库存天数
}

const (
	StorehouseInboundTypePurchaseInbound = 1 // 采购入库
	StorehouseInboundTypeReturnInbound   = 2 // 退货入库
	StorehouseInboundTypeManualInbound   = 3 // 手工入库

	StorehouseInboundStatusPending   = 1 // 待处理
	StorehouseInboundStatusHandled   = 2 // 已处理
	StorehouseInboundStatusCanceled  = 3 // 已取消
	StorehouseInboundStatusCompleted = 4 // 已完成
)

// 请求入库信息
type StorehouseInboundReq struct {
	StorehouseUuid string `json:"storehouse_uuid" binding:"required"` // 仓库UUID
	Title          string `json:"title" binding:"required"`           // 标题
	// 采购订单
	PurchaseOrderNo string `json:"purchase_order_no" binding:"-"` // 采购订单

	// 销售订单
	SalesOrderNo string `json:"sales_order_no" binding:"-"` // 销售订单

	// 订单类型
	PurchaseOrderProductType string                       `json:"purchase_order_product_type" binding:"-"` // 采购订单物品类型 1：期货 2：现货
	CustomerUuid             string                       `json:"customer_uuid" binding:"-"`               // 客户UUID
	InboundType              string                       `json:"inbound_type" binding:"required"`         // 入库类型 1:采购入库 2:退货入库 3:手工入库
	Status                   int                          `json:"status" binding:"-"`                      // 状态 1:待处理 2: 已处理 3:已取消 4:已完成
	InDate                   string                       `json:"in_date" binding:"-"`                     // 入库日期
	Detail                   []StorehouseInboundDetailReq `json:"detail" binding:"required"`               // 入库明细
}

type StorehouseInboundUpdateReq struct {
	StorehouseInboundReq
	InboundOrderNo string `json:"inbound_order_no" binding:"required"` // 入库单号
}

// 请求入库明细信息
type StorehouseInboundDetailReq struct {
	PurchaseOrderProductNo string  `json:"purchase_order_product_no" binding:"-"` // 采购订单物品编号
	ProductUuid            string  `json:"product_uuid" binding:"required"`       // 商品UUID
	SkuUuid                string  `json:"sku_uuid" binding:"required"`           // SKU UUID
	Quantity               float64 `json:"quantity" binding:"required"`           // 入库数量
	BoxNum                 float64 `json:"box_num" binding:"-"`                   // 箱数
	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号

}

// 入库
type StorehouseInbound struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID

	// 采购订单
	PurchaseOrderNo string `json:"purchase_order_no" gorm:"comment:'采购订单'"`

	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"` //

	// 入库单号
	InboundOrderNo string `json:"inbound_order_no" gorm:"comment:'入库单号'"` // 入库单号

	Title string `json:"title" gorm:"comment:'标题'"` // 标题

	// 入库日期
	InboundDate string `json:"inbound_date" gorm:"comment:'入库日期'"` // 入库日期
	// 入库类型
	InboundType string `json:"inbound_type" gorm:"comment:'入库类型'"` // 入库类型 1:采购入库 2:退货入库 3:手工入库
	// 入库状态
	Status int `json:"status" gorm:"comment:'状态'"` // 状态 1:已入库 2:未入库
	// 入库人
	InboundBy string `json:"inbound_by" gorm:"comment:'入库人'"`                 // 入库人
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type StorehouseInboundRes struct {
	StorehouseInbound
	Storehouse    Storehouse `json:"storehouse"`
	InboundByUser User       `json:"inbound_by_user"`
}

type StorehouseInboundItemRes struct {
	StorehouseInbound
	StorehouseInboundDetailRes
	CustomerInfo  Customer   `json:"customer_info"`
	Storehouse    Storehouse `json:"storehouse"`
	InboundByUser User       `json:"inbound_by_user"`
}

// 入库明细
type StorehouseInboundDetail struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID

	StorehouseUuid         string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID
	PurchaseOrderProductNo string `json:"purchase_order_product_no" gorm:"comment:'采购订单物品编号'"`         // 采购订单物品编号

	// 入库单号
	InboundOrderNo string `json:"inbound_order_no" gorm:"comment:'入库单号'"` // 入库单号
	// 采购订单物品类型
	PurchaseOrderProductType string `json:"purchase_order_product_type" gorm:"comment:'采购订单物品类型'"` // 采购订单物品类型 1：期货 2：现货
	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号
	// 入库数量
	Quantity float64 `json:"quantity" gorm:"comment:'入库数量'"` // 入库数量
	BoxNum   float64 `json:"box_num" gorm:"comment:'箱数'"`    // 箱数

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type StorehouseInboundDetailRes struct {
	StorehouseInboundDetail
	Product Product `json:"product"`
	Sku     Sku     `json:"sku"`
}

type StorehouseInboundDetailInfoRes struct {
	StorehouseInboundDetail
	StorehouseInbound StorehouseInboundRes `json:"storehouse_inbound"`
	PurchaseOrderInfo PurchaseOrder        `json:"purchase_order_info"`
	Product           Product              `json:"product"`
	Sku               Sku                  `json:"sku"`
	CustomerInfo      Customer             `json:"customer_info"`
}

const (
	StorehouseOutboundTypeSalesOutbound  = 1 // 销售出库
	StorehouseOutboundTypeReturnOutbound = 2 // 退货出库
	StorehouseOutboundTypeManualOutbound = 3 // 手工出库

	StorehouseOutboundStatusPending   = 1 // 待处理
	StorehouseOutboundStatusHandled   = 2 // 已处理
	StorehouseOutboundStatusCanceled  = 3 // 已取消
	StorehouseOutboundStatusCompleted = 4 // 已完成
)

// 请求出库信息
type StorehouseOutboundReq struct {
	StorehouseUuid string `json:"storehouse_uuid" binding:"required"` // 仓库UUID
	OutboundType   string `json:"outbound_type" binding:"-"`          // 出库类型 1:销售出库 2:退货出库 3:手工出库
	OutboundDate   string `json:"outbound_date" binding:"-"`          // 出库日期
	Status         int    `json:"status" binding:"-"`                 // 状态 1:待处理 2: 已处理 3:已取消 4:已完成
	// 销售订单
	SalesOrderNo string `json:"sales_order_no" binding:"-"` // 销售订单
	// 订单类型
	SalesOrderProductType string                        `json:"sales_order_product_type" binding:"-"` // 销售订单物品类型 1：期货 2：现货
	CustomerUuid          string                        `json:"customer_uuid" binding:"-"`            // 客户UUID
	Detail                []StorehouseOutboundDetailReq `json:"detail" binding:"required"`            // 出库明细
}

type StorehouseOutboundDetailReq struct {
	StorehouseProductUuid string  `json:"storehouse_product_uuid" binding:"required"` // 仓库物品UUID
	ProductUuid           string  `json:"product_uuid" binding:"required"`            // 商品UUID
	SkuUuid               string  `json:"sku_uuid" binding:"required"`                // SKU UUID
	Quantity              float64 `json:"quantity" binding:"required"`                // 入库数量
	BoxNum                float64 `json:"box_num" binding:"-"`                        // 箱数
	// 柜号
	CabinetNo string `json:"cabinet_no"` // 柜号
}

// 出库
type StorehouseOutbound struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID

	// 销售订单
	SalesOrderNo string `json:"sales_order_no" gorm:"comment:'销售订单'"`
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号

	// 出库日期
	OutboundDate string `json:"outbound_date" gorm:"comment:'出库日期'"` // 出库日期
	// 出库类型
	OutboundType string `json:"outbound_type" gorm:"comment:'出库类型'"` // 出库类型

	SalesOrderProductType string `json:"sales_order_product_type" gorm:"comment:'销售订单物品类型'"` // 销售订单物品类型 1：期货 2：现货

	// 出库状态
	Status       int    `json:"status" gorm:"comment:'状态'"`                                // 状态 1:已出库 2:未出库
	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"` // 客户UUID
	// 出库人
	OutboundBy string `json:"outbound_by" gorm:"comment:'出库人'"`                // 出库人
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间

}

type StorehouseOutboundRes struct {
	StorehouseOutbound
	StorehouseOutboundDetailRes
	Storehouse     Storehouse `json:"storehouse"`
	CustomerInfo   Customer   `json:"customer_info"`
	OutboundByUser User       `json:"outbound_by_user"`
}

// 出库明细
type StorehouseOutboundDetail struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID

	Uuid           string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`              // UUID
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID
	SalesOrderType string `json:"sales_order_type" gorm:"comment:'销售订单类型'"`                    // 销售订单类型 1:期货 2:现货
	// 销售订单
	SalesOrderNo string `json:"sales_order_no" gorm:"comment:'销售订单'"`
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号

	// 产品库存uuid
	StorehouseProductUuid string `json:"storehouse_product_uuid" gorm:"type:char(36);index;comment:'仓库物品UUID'"` // 仓库物品UUID

	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号
	// 出库数量
	Quantity  float64 `json:"quantity" gorm:"comment:'出库数量'"`                  // 出库数量
	BoxNum    float64 `json:"box_num" gorm:"comment:'箱数'"`                     // 箱数
	CreatedAt string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type StorehouseOutboundDetailRes struct {
	StorehouseOutboundDetail
	Product Product `json:"product"`
	Sku     Sku     `json:"sku"`
}

// 仓库盘点请求
type StorehouseInventoryCheckReq struct {
	StorehouseUuid string                              `json:"storehouse_uuid" binding:"required"` // 仓库UUID
	CheckDate      string                              `json:"check_date" binding:"required"`      // 盘点日期
	Status         int                                 `json:"status" binding:"required"`          // 状态 1:已盘点 2:未盘点
	Detail         []StorehouseInventoryCheckDetailReq `json:"detail" binding:"required"`          // 盘点明细
}

type StorehouseInventoryCheckDetailReq struct {
	StorehouseProductUuid string  `json:"storehouse_product_uuid" binding:"required"` // 仓库物品UUID
	ProductUuid           string  `json:"product_uuid" binding:"required"`            // 商品UUID
	SkuUuid               string  `json:"sku_uuid" binding:"required"`                // SKU UUID
	Quantity              float64 `json:"quantity" binding:"required"`                // 盘点数量
	BoxNum                float64 `json:"box_num" binding:"-"`                        // 箱数
	// 差异op
	DifferenceOp string `json:"difference_op" gorm:"comment:'差异操作'"` // 差异操作 1:盘盈 2:盘亏
}

// 仓库盘点
type StorehouseInventoryCheck struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID
	// 盘点单号
	CheckOrderNo string `json:"check_order_no" gorm:"comment:'盘点单号'"` // 盘点单号
	// 盘点日期
	CheckDate string `json:"check_date" gorm:"comment:'盘点日期'"` // 盘点日期
	// 盘点状态
	Status int `json:"status" gorm:"comment:'状态'"` // 状态 1:已盘点 2:未盘点
	// 盘点人
	CheckBy   string `json:"check_by" gorm:"comment:'盘点人'"`                   // 盘点人
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type StorehouseInventoryCheckRes struct {
	StorehouseInventoryCheckDetailRes
	StorehouseInventoryCheck StorehouseInventoryCheck `json:"storehouse_inventory_check"`
	Storehouse               Storehouse               `json:"storehouse"`
	StorehouseProduct        StorehouseProduct        `json:"storehouse_product"`
	CheckByUser              User                     `json:"check_by_user"`
}

// 仓库盘点明细
type StorehouseInventoryCheckDetail struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 盘点单号
	Uuid                  string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`                        // UUID
	CheckOrderNo          string `json:"check_order_no" gorm:"comment:'盘点单号'"`                                  // 盘点单号
	StorehouseUuid        string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"`           // 仓库UUID
	StorehouseProductUuid string `json:"storehouse_product_uuid" gorm:"type:char(36);index;comment:'仓库物品UUID'"` // 仓库物品UUID
	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 盘点数量

	Quantity float64 `json:"quantity" gorm:"comment:'盘点数量'"` // 盘点数量
	BoxNum   float64 `json:"box_num" gorm:"comment:'箱数'"`    // 箱数

	// 差异op
	DifferenceOp string `json:"difference_op" gorm:"comment:'差异操作'"` // 差异操作 1:盘盈 2:盘亏

	DifferenceBoxNumOp string `json:"difference_box_num_op" gorm:"comment:'差异箱数操作'"` // 差异箱数操作 1:盘盈 2:盘亏

	// 差异数量
	DifferenceQuantity float64 `json:"difference_quantity" gorm:"comment:'差异数量'"` // 差异数量

	DifferenceBoxNum float64 `json:"difference_box_num" gorm:"comment:'差异箱数'"` // 差异箱数

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type StorehouseInventoryCheckDetailRes struct {
	StorehouseInventoryCheckDetail
	Product Product `json:"product"`
	Sku     Sku     `json:"sku"`
}

type ReqStorehouseOutboundOrder struct {
	StorehouseProductUuids []string `json:"storehouse_product_uuids" binding:"required"` // 仓库物品UUID
	// 是否整柜
	IsFullCabinet bool `json:"is_full_cabinet" binding:"-"` // 是否整柜

	ApplyBy string `json:"apply_by" binding:"-"` // 申请人
}

type ReqStorehouseOutboundOrderQueryParam struct {
	Pagination
}

// 出货单
type StorehouseOutboundOrder struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID
	// 销售订单
	SalesOrderNo string `json:"sales_order_no" gorm:"comment:'销售订单'"`
	// 出库日期
	OutboundDate string `json:"outbound_date" gorm:"comment:'出库日期'"` // 出库日期
	// 出库类型
	OutboundType string `json:"outbound_type" gorm:"comment:'出库类型'"` // 出库类型 自提
	// 出库状态
	Status       string `json:"status" gorm:"comment:'状态'"`                                // 状态  申请中  已出库
	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"` // 客户UUID

	// 出库人
	OutboundBy string `json:"outbound_by" gorm:"comment:'出库人'"`                // 出库人
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// 出货单明细
type StorehouseOutboundOrderDetail struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号

	// 销售订单
	SalesOrderNo string `json:"sales_order_no" gorm:"comment:'销售订单'"`

	// 采购单号
	PurchaseOrderNo string `json:"purchase_order_no" gorm:"comment:'采购订单'"`

	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID

	// 产品库存uuid
	StorehouseProductUuid string `json:"storehouse_product_uuid" gorm:"type:char(36);index;comment:'仓库物品UUID'"` // 仓库物品UUID

	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号

	// 申请人
	ApplyBy string `json:"apply_by" gorm:"comment:'申请人'"` // 申请人

	// 出库状态
	Status       string `json:"status" gorm:"comment:'状态'"`                                // 状态  申请中  已出库
	CustomerUuid string `json:"customer_uuid" gorm:"type:char(36);index;comment:'客户UUID'"` // 客户UUID
	// 出库数量
	Quantity  float64 `json:"quantity" gorm:"comment:'出库数量'"`                  // 出库数量
	BoxNum    float64 `json:"box_num" gorm:"comment:'箱数'"`                     // 箱数
	CreatedAt string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

type StorehouseOutboundOrderDetailRes struct {
	StorehouseOutboundOrderDetail
	Product           Product       `json:"product"`
	Sku               Sku           `json:"sku"`
	PurchaseOrderInfo PurchaseOrder `json:"purchase_order_info"`
	StorehouseInfo    Storehouse    `json:"storehouse_info"`
}
