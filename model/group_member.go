package model

type GroupMember struct {
	Id        uint   `gorm:"primary_key" json:"id"`               // ID 是团队成员的主键
	UUID      string `gorm:"type:char(36);unique" json:"uuid"`    // UUID 是团队成员的唯一标识符
	GroupUuid string `gorm:"type:char(36)" json:"group_uuid"`     // TeamUUID 是团队的UUID
	UserUUID  string `gorm:"type:char(36)" json:"user_uuid"`      // UserUUID 是用户的UUID
	GroupType string `gorm:"type:varchar(100)" json:"group_type"` // GroupType 是团队类型
	Role      string `gorm:"type:varchar(100)" json:"role"`       // Role 是成员在团队中的角色
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了团队成员加入的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了团队成员信息最后更新的时间
}
