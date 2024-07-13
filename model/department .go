package model

type Department struct {
	Id               uint   `gorm:"primary_key" json:"id"`
	UUID             string `gorm:"type:char(36);unique" json:"uuid"`     // UUID 是部门的唯一标识符
	Name             string `gorm:"type:varchar(100)" json:"name"`        // Name 是部门的名称
	Description      string `gorm:"type:varchar(255)" json:"description"` // Description 是对部门的描述
	Manager          string `gorm:"type:varchar(100)" json:"manager"`     // Manager 是部门的负责人
	ParentDepartment string `gorm:"type:varchar(100)" json:"parent_uuid"` // ParentDepartment 是部门的上级部门
	Status           int    `gorm:"type:int" json:"status"`               // Status 是部门的状态 0:未启用 1:启用
	CreatedAt        string `gorm:"autoCreateTime" json:"created_at"`     // CreatedAt 记录了部门创建的时间
	UpdatedAt        string `gorm:"autoUpdateTime" json:"updated_at"`     // UpdatedAt 记录了部门信息最后更新的时间
}
