package model

type SkuReq struct {
	Name                string `json:"name" binding:"required"`                  // Name 是SKU的名称
	ProductCategoryUuid string `json:"product_category_uuid" binding:"required"` // ProductCategoryUuid 是商品分类的UUID
	Code                string `json:"code" binding:"-"`                         // Code 是SKU的编码
	// 规格
	Specification string `json:"specification" binding:"-"` // Specification 是SKU的规格
	// 单位
	Unit string `json:"unit" binding:"-"` // Unit 是SKU的单位
	// 国家
	Country string `json:"country" binding:"required"` // Country 是SKU的国家
	// 厂号
	FactoryNo string `json:"factory_no" binding:"required"` // FactoryNo 是SKU的厂号
}

type Sku struct {
	Id                  uint   `gorm:"primary_key" json:"id"`                            // ID 是SKU的主键
	UUID                string `gorm:"type:char(36);unique" json:"uuid"`                 // UUID 是SKU的唯一标识符
	ProductUuid         string `gorm:"type:char(36);index" json:"product_uuid"`          // ProductUuid 是商品的UUID
	ProductName         string `gorm:"type:varchar(100)" json:"product_name"`            // ProductName 是商品的名称
	ProductCategoryUuid string `gorm:"type:char(36);index" json:"product_category_uuid"` // ProductCategoryUuid 是商品分类的UUID
	Code                string `gorm:"type:varchar(100)" json:"code"`                    // Name 是SKU的Code
	// 规格
	Specification string `gorm:"type:varchar(100)" json:"specification"` // Specification 是SKU的规格
	// 单位
	// 国家
	Country string `gorm:"type:varchar(100)" json:"country"` // Country 是SKU的国家
	// 厂号
	FactoryNo string `gorm:"type:varchar(100)" json:"factory_no"` // FactoryNo 是SKU的厂号
	Unit      string `gorm:"type:varchar(100)" json:"unit"`       // Unit 是SKU的单位
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了最后更新的时间
	IsDeleted int    `gorm:"default:0" json:"is_deleted"`         // IsDeleted 标识SKU是否被删除
}

type SkuRes struct {
	Sku
	Product         Product         `json:"product"`
	ProductCategory ProductCategory `json:"product_category"`
}
