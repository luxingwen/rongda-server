package model

// 用户关系
type UserRelation struct {
	// 自增ID
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`                     // 主键ID
	Uuid string `json:"uuid" gorm:"uniqueIndex;type:varchar(50);comment:'UUID'"` // UUID
	// 用户ID
	UserUuid string `json:"user_uuid" gorm:"comment:'用户ID'"` // 用户ID

	FromModule string `json:"from_module" gorm:"comment:'来源模块'"` // 来

	RefModule string `json:"ref_module" gorm:"comment:'关联模块'"` // 关联模块

	RefUuid string `json:"ref_uuid" gorm:"comment:'关联UUID'"` // 关联UUID

	CreatedAt string `gorm:"autoCreateTime" json:"create_at"` // CreatedAt 记录了创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"create_at"` // UpdatedAt 记录了信息最后更新的时间
}
