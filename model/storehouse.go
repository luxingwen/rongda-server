package model

type Storehouse struct {
	ID      uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid    string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	Name    string `json:"name" gorm:"comment:'名称'"`                       // 名称
	Address string `json:"address" gorm:"comment:'地址'"`                    // 地址
	// 联系人
	ContactPerson string `json:"contact_person" gorm:"comment:'联系人'"` // 联系人
	// 联系电话
	ContactPhone string `json:"contact_phone" gorm:"comment:'联系电话'"` // 联系电话
	Status       int    `json:"status" gorm:"comment:'状态'"`          // 状态 1:启用 2:未启用
	// 仓库类型
	Type      string `json:"type" gorm:"comment:'仓库类型'"`                      // 仓库类型 1:自有仓库 2:第三方仓库
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// 入库
type StorehouseInbound struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID
	// 入库单号
	InboundOrderNo string `json:"inbound_order_no" gorm:"comment:'入库单号'"` // 入库单号
	// 入库日期
	InboundDate string `json:"inbound_date" gorm:"comment:'入库日期'"` // 入库日期
	// 入库类型
	InboundType string `json:"inbound_type" gorm:"comment:'入库类型'"` // 入库类型
	// 入库状态
	Status int `json:"status" gorm:"comment:'状态'"` // 状态 1:已入库 2:未入库
	// 入库人
	InboundBy string `json:"inbound_by" gorm:"comment:'入库人'"`                 // 入库人
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间

}

// 入库明细
type StorehouseInboundDetail struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 入库单号
	InboundOrderNo string `json:"inbound_order_no" gorm:"comment:'入库单号'"` // 入库单号
	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 入库数量
	Quantity  int    `json:"quantity" gorm:"comment:'入库数量'"`                  // 入库数量
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// 出库
type StorehouseOutbound struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid" gorm:"type:char(36);index;comment:'仓库UUID'"` // 仓库UUID
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号
	// 出库日期
	OutboundDate string `json:"outbound_date" gorm:"comment:'出库日期'"` // 出库日期
	// 出库类型
	OutboundType string `json:"outbound_type" gorm:"comment:'出库类型'"` // 出库类型
	// 出库状态
	Status int `json:"status" gorm:"comment:'状态'"` // 状态 1:已出库 2:未出库
	// 出库人
	OutboundBy string `json:"outbound_by" gorm:"comment:'出库人'"`                // 出库人
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间

}

// 出库明细
type StorehouseOutboundDetail struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 出库单号
	OutboundOrderNo string `json:"outbound_order_no" gorm:"comment:'出库单号'"` // 出库单号
	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 出库数量
	Quantity  int    `json:"quantity" gorm:"comment:'出库数量'"`                  // 出库数量
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
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

// 仓库盘点明细
type StorehouseInventoryCheckDetail struct {
	ID uint `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	// 盘点单号
	CheckOrderNo string `json:"check_order_no" gorm:"comment:'盘点单号'"` // 盘点单号
	// 商品uuid
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'商品UUID'"` // 商品UUID
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"`   // SKU UUID
	// 盘点数量
	Quantity int `json:"quantity" gorm:"comment:'盘点数量'"` // 盘点数量

	// 差异op
	DifferenceOp string `json:"difference_op" gorm:"comment:'差异操作'"` // 差异操作 1:盘盈 2:盘亏
	// 差异数量
	DifferenceQuantity int `json:"difference_quantity" gorm:"comment:'差异数量'"` // 差异数量

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
