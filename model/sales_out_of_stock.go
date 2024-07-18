package model

type SalesOutOfStockReq struct {
	// 出库日期
	OutOfStockDate string `json:"out_of_stock_date"` // OutOfStockDate 是出库日期
	// 销售单号
	SalesOrderNo string `json:"sales_order_no"` // SalesOrderNo 是销售单号
	CustomerUuid string `json:"customer_uuid"`  // CustomerUuid 是客户的UUID
	BatchNo      string `json:"batch_no"`       // BatchNo 是批次号
	// 登记人
	Registrant string `json:"registrant"` // Registrant 是登记人
	// 仓库uuid
	StorehouseUuid string `json:"storehouse_uuid"` // StorehouseUuid 是仓库的UUID
	// 备注
	Remark string `json:"remark"` // Remark 是备注
	// 状态
	Status int                      `json:"status"` // Status 是状态 1：未出库 2：已出库
	Items  []SalesOutOfStockItemReq `json:"items"`  // Items 是出库单明细
}

type SalesOutOfStockItemReq struct {
	ProductUuid string `json:"product_uuid"` // ProductUuid 是产品的UUID
	SkuUuid     string `json:"sku_uuid"`     // SkuUuid 是SKU的UUID
	// 数量
	Quantity int `json:"quantity"` // Quantity 是数量
	// 价格
	Price float64 `json:"price"` // Price 是价格
	// 总金额
	TotalAmount float64 `json:"total_amount"` // TotalAmount 是总金额
}

type SalesOutOfStock struct {
	ID   uint   `gorm:"primary_key" json:"id"`            // ID 主键
	UUID string `gorm:"type:char(36);unique" json:"uuid"` // UUID 是出库单的唯一标识符
	// 出库日期
	OutOfStockDate string `gorm:"type:date" json:"out_of_stock_date"` // OutOfStockDate 是出库日期
	// 销售单号
	SalesOrderNo string `gorm:"type:varchar(100)" json:"sales_order_no"`  // SalesOrderNo 是销售单号
	CustomerUuid string `gorm:"type:char(36);index" json:"customer_uuid"` // CustomerUuid 是客户的UUID
	BatchNo      string `gorm:"type:varchar(100)" json:"batch_no"`        // BatchNo 是批次号
	// 登记人
	Registrant string `gorm:"type:varchar(100)" json:"registrant"` // Registrant 是登记人
	// 仓库uuid
	StorehouseUuid string `gorm:"type:char(36);index" json:"storehouse_uuid"` // StorehouseUuid 是仓库的UUID
	// 备注
	Remark string `gorm:"type:varchar(255)" json:"remark"` // Remark 是备注
	// 状态
	Status int `gorm:"type:int" json:"status"` // Status 是状态 1：未出库 2：已出库
	// 创建时间
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // CreatedAt 记录了创建的时间
	// 更新时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"` // UpdatedAt 记录了最后更新的时间
}

type SalesOutOfStockRes struct {
	SalesOutOfStock
	CustomerInfo   *Customer   `json:"customer_info"`   // CustomerInfo 是客户信息
	StorehouseInfo *Storehouse `json:"storehouse_info"` // StorehouseInfo 是仓库信息
	RegistrantInfo *User       `json:"registrant_info"` // RegistrantInfo 是登记人信息
}

type SalesOutOfStockItem struct {
	ID                uint   `gorm:"primary_key" json:"id"`                            // ID 主键
	SalesOutOfStockNo string `gorm:"type:char(36);index" json:"sales_out_of_stock_no"` // SalesOutOfStockNo 是出库单号
	ProductUuid       string `gorm:"type:char(36);index" json:"product_uuid"`          // ProductUuid 是产品的UUID
	// 产品名称
	ProductName string `gorm:"type:varchar(100)" json:"product_name"` // ProductName 是产品名称
	SkuUuid     string `gorm:"type:char(36);index" json:"sku_uuid"`   // SkuUuid 是SKU的UUID
	// SKU名称
	SkuName string `gorm:"type:varchar(100)" json:"sku_name"` // SkuName 是SKU名称
	// 数量
	Quantity int `gorm:"type:int" json:"quantity"` // Quantity 是数量
	// 价格
	Price float64 `gorm:"type:decimal(10,2)" json:"price"` // Price 是价格
	// 总金额
	TotalAmount float64 `gorm:"type:decimal(10,2)" json:"total_amount"` // TotalAmount 是总金额
	// 创建时间
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // CreatedAt 记录了创建的时间
	// 更新时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"` // UpdatedAt 记录了最后更新的时间
}

type SalesOutOfStockItemRes struct {
	SalesOutOfStockItem
	ProductInfo *Product `json:"product_info"` // ProductInfo 是产品信息
	SkuInfo     *Sku     `json:"sku_info"`     // SkuInfo 是SKU信息
}
