package model

type OrderFile struct {
	ID        uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Uuid      string `json:"uuid" gorm:"type:char(36);index"`     // UUID 是订单文件的唯一标识符
	OrderNo   string `json:"order_no" gorm:"comment:'订单编号'"`      // 订单编号
	OrderType string `json:"order_type" gorm:"comment:'订单类型'"`    // 订单类型
	Name      string `json:"name" gorm:"comment:'文件名'"`           // 文件名
	Filename  string `json:"filename" gorm:"comment:'文件名'"`       // 文件名
	Url       string `json:"url" gorm:"comment:'文件地址'"`           // 文件地址
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`    // CreatedAt 记录了创建的时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`    // UpdatedAt 记录了最后更新的时间
}
