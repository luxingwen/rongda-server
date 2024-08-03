package model

type Feedback struct {
	Id        uint   `gorm:"primary_key" json:"id"`            // ID 是反馈的主键
	UUID      string `gorm:"type:char(36);unique" json:"uuid"` // UUID 是反馈的唯一标识符
	Content   string `gorm:"type:varchar(255)" json:"content"` // Content 是反馈的内容
	From      string `gorm:"type:varchar(100)" json:"from"`    // From 是反馈的来源
	UserUuid  string `gorm:"type:char(36)" json:"user_uuid"`   // UserUuid 是反馈的用户
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // CreatedAt 记录了反馈创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"` // UpdatedAt 记录了反馈最后更新的时间
}
