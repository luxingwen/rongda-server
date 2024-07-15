package model

type PurchaseOrderReq struct {
	// 标题
	Title string `json:"title" gorm:"comment:'标题'"`  // 标题
	Date  string `json:"date" gorm:"comment:'采购日期'"` // 采购日期
	// 供应商
	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID
	// 定金
	Deposit float64 `json:"deposit" gorm:"comment:'定金'"` // 定金
	// 税费
	Tax float64 `json:"tax" gorm:"comment:'税费'"` // 税费
	// 总金额
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"` // 总金额
	// 采购合同
	AgreementUuid string `json:"agreement_uuid" gorm:"type:char(36);index;comment:'采购合同UUID'"` // 采购合同UUID

	Details []PurchaseOrderItemReq `json:"details" gorm:"comment:'采购单明细'"` // 采购单明细
}

type PurchaseOrderItemReq struct {
	ProductUuid string `json:"product_uuid" gorm:"type:char(36);index;comment:'产品UUID'"` // 产品UUID
	// SKU UUID
	SkuUuid string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"` // SKU UUID
	// 产品名称
	ProductName string `json:"product_name" gorm:"comment:'产品名称'"` // 产品名称
	// SKU名称
	SkuName string `json:"sku_name" gorm:"comment:'SKU名称'"` // SKU名称
	// 数量
	Quantity int `json:"quantity" gorm:"comment:'数量'"` // 数量
	// 价格
	Price float64 `json:"price" gorm:"comment:'价格'"` // 价格
	// 总金额
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"` // 总金额
}

// PurchaseOrder 采购单
type PurchaseOrder struct {
	ID           uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`                        // 主键ID
	Title        string `json:"title" gorm:"comment:'标题'"`                                  // 标题
	OrderNo      string `json:"order_no" gorm:"comment:'采购单号'"`                             // 采购单号
	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID
	// 采购日期
	Date string `json:"date" gorm:"comment:'采购日期'"` // 采购日期
	// 采购单号

	// 定金
	Deposit float64 `json:"deposit" gorm:"comment:'定金'"` // 定金
	// 税费
	Tax float64 `json:"tax" gorm:"comment:'税费'"` // 税费
	// 总金额
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"` // 总金额
	// 采购人
	Purchaser string `json:"purchaser" gorm:"comment:'采购人'"` // 采购人

	// 采购单状态
	Status int `json:"status" gorm:"comment:'采购单状态'"` // 采购单状态 1： 待处理 2：已处理 3：已取消 4：已完成

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// PurchaseOrderItem 采购单明细
type PurchaseOrderItem struct {
	ID              uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`                          // 主键ID
	PurchaseOrderNo string `json:"purchase_order_no" gorm:"type:char(36);index;comment:'采购单单号'"` // 采购单号
	ProductUuid     string `json:"product_uuid" gorm:"type:char(36);index;comment:'产品UUID'"`     // 产品UUID
	// 产品名称
	ProductName string `json:"product_name" gorm:"comment:'产品名称'"` // 产品名称

	SkuUuid string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"` // SKU UUID
	// SKU名称
	SkuName string `json:"sku_name" gorm:"comment:'SKU名称'"` // SKU名称
	// 数量
	Quantity int `json:"quantity" gorm:"comment:'数量'"` // 数量
	// 价格
	Price float64 `json:"price" gorm:"comment:'价格'"` // 价格
	// 总金额
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"`               // 总金额
	CreatedAt   string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt   string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// PurchaseOrderResp 采购单响应
type PurchaseOrderRes struct {
	PurchaseOrder
	Supplier Supplier `json:"supplier" gorm:"comment:'供应商'"` // 供应商
}
