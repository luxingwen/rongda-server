package model

type Sku struct {
	Id   uint   `gorm:"primary_key" json:"id"`            // ID 是SKU的主键
	UUID string `gorm:"type:char(36);unique" json:"uuid"` // UUID 是SKU的唯一标识符
	Name string `gorm:"type:varchar(100)" json:"name"`    // Name 是SKU的名称
	// 单位
	Unit      string `gorm:"type:varchar(100)" json:"unit"`    // Unit 是SKU的单位
	Num       int    `gorm:"type:int" json:"num"`              // Num 是SKU的数量
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // CreatedAt 记录了创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"` // UpdatedAt 记录了最后更新的时间
}
