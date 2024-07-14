package model

const (
	ProductStatusAvailable = 1 // 商品状态可用
)

type Product struct {
	ID            uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Uuid          string  `json:"uuid" gorm:"comment:'商品唯一标识'"`        // 商品唯一标识
	Name          string  `json:"name" gorm:"comment:'商品名称'"`          // 商品名称
	Category      string  `json:"category" gorm:"comment:'商品类别'"`      // 商品类别
	Specification string  `json:"specification" gorm:"comment:'规格'"`   // 规格
	Description   string  `json:"description" gorm:"comment:'描述'"`     // 描述
	Price         float64 `json:"price" gorm:"comment:'价格'"`           // 价格
	Cost          float64 `json:"cost" gorm:"comment:'成本'"`            // 成本
	Supplier      string  `json:"supplier" gorm:"comment:'供应商'"`       // 供应商
	Creater       string  `json:"creater" gorm:"comment:'创建人'"`        // 创建人
	CreatedAt     string  `json:"created_at" gorm:"comment:'创建时间'"`    // 创建时间
	UpdatedAt     string  `json:"updated_at" gorm:"comment:'更新时间'"`    // 更新时间
}

type ProductRes struct {
	Product
	SupplierInfo *Supplier `json:"supplier_info"`
}
