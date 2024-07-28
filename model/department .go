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

// 部门员工
type DepartmentStaff struct {
	Id uint `gorm:"primary_key" json:"id"`

	Uuid string `gorm:"type:char(36);unique" json:"uuid"` // UUID 是部门员工的唯一标识符

	DepartmentUuid string `gorm:"type:char(36);index" json:"department_uuid"` // DepartmentUuid 是部门员工所属部门的 UUID
	// StaffUuid 是部门员工的 UUID
	StaffUuid string `gorm:"type:char(36);index" json:"staff_uuid"`
	// StaffName 是部门员工的名称
	StaffName string `gorm:"type:varchar(100)" json:"staff_name"`
	// StaffNo 是部门员工的工号
	StaffNo string `gorm:"type:varchar(100)" json:"staff_no"`
	// StaffPosition 是部门员工的职位
	StaffPosition string `gorm:"type:varchar(100)" json:"staff_position"`

	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // CreatedAt 记录了部门员工创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"` // UpdatedAt 记录了部门员工信息最后更新的时间
}
