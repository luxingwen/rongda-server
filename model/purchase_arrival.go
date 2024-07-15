package model

type PurchaseArrivalReq struct {
	PurchaseOrderNo  string `json:"purchase_order_no" form:"purchase_order_no" binding:"required"` // 采购单号
	Batch            string `json:"batch" form:"batch" binding:"required"`                         // 批次
	ArrivalDate      string `json:"arrival_date" form:"arrival_date" binding:"required"`           // 到货日期
	SupplierUuid     string `json:"supplier_uuid" form:"supplier_uuid" binding:"-"`                // 供应商UUID
	Acceptor         string `json:"acceptor" form:"acceptor" binding:"required"`                   // 验收人
	AcceptanceResult string `json:"acceptance_result" form:"acceptance_result" binding:"required"` // 验收结果
	// 仓库UUID
	StorageUuid string                   `json:"storage_uuid" form:"storage_uuid" binding:"-"` // 仓库UUID
	Remark      string                   `json:"remark" form:"remark"`                         // 备注
	Items       []PurchaseArrivalItemReq `json:"items" form:"items" binding:"required"`        // 到货明细
}

type PurchaseArrivalItemReq struct {
	ProductUuid string  `json:"product_uuid" form:"product_uuid" binding:"required"` // 产品UUID
	ProductName string  `json:"product_name" form:"product_name" binding:"required"` // 产品名称
	SkuUuid     string  `json:"sku_uuid" form:"sku_uuid" binding:"required"`         // SKU UUID
	SkuName     string  `json:"sku_name" form:"sku_name" binding:"required"`         // SKU名称
	Quantity    int     `json:"quantity" form:"quantity" binding:"required"`         // 数量
	Price       float64 `json:"price" form:"price" binding:"required"`               // 价格
	TotalAmount float64 `json:"total_amount" form:"total_amount" binding:"required"` // 总金额
}

type PurchaseArrival struct {
	ID              uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid            string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	PurchaseOrderNo string `json:"purchase_order_no" gorm:"comment:'采购单号'"`        // 采购单号
	// 供应商
	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID

	// 批次
	Batch string `json:"batch" gorm:"comment:'批次'"` // 批次
	// 到货日期
	ArrivalDate string `json:"arrival_date" gorm:"comment:'到货日期'"` // 到货日期
	// 验收人
	Acceptor string `json:"acceptor" gorm:"comment:'验收人'"` // 验收人
	// 验收结果
	AcceptanceResult string `json:"acceptance_result" gorm:"comment:'验收结果'"` // 验收结果 1：合格 2：不合格
	Remark           string `json:"remark" gorm:"comment:'备注'"`              // 备注
	// 到货状态
	Status      int     `json:"status" gorm:"comment:'到货状态'"`      // 到货状态 1：待处理 2：已处理 3：已取消 4：已完成
	TotalAmount float64 `json:"total_amount" gorm:"comment:'总金额'"` // 总金额

	// 存储仓库uuid
	StorageUuid string `json:"storage_uuid" gorm:"type:char(36);index;comment:'存储仓库UUID'"` // 存储仓库UUID

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// 到货明细
type PurchaseArrivalItem struct {
	ID                uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`                           // 主键ID
	PurchaseArrivalNo string `json:"purchase_arrival_no" gorm:"type:char(36);index;comment:'到货单号'"` // 到货单号
	ProductUuid       string `json:"product_uuid" gorm:"type:char(36);index;comment:'产品UUID'"`      // 产品UUID
	// 产品名称
	ProductName string `json:"product_name" gorm:"comment:'产品名称'"`                     // 产品名称
	SkuUuid     string `json:"sku_uuid" gorm:"type:char(36);index;comment:'SKU UUID'"` // SKU UUID
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
