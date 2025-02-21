package model

import "time"

const (
	RoleTypeAdmin = "admin"
	// 客户
	RoleTypeCustomer = "customer"
	// 销售
	RoleTypeSales = "sales"
	// 运营
	RoleTypeOperation = "operation"
	// 法务
	RoleTypeLegal = "legal"
	// 财务
	RoleTypeFinance = "finance"
	// 银行人员
	RoleTypeBank = "bank"
)

// Role 结构体定义了角色的基础信息
type Role struct {
	ID        uint      `gorm:"primary_key" json:"id"`                // ID 是角色的主键
	Uuid      string    `gorm:"type:varchar(100);unique" json:"uuid"` // Uuid 是角色的唯一标识
	Name      string    `gorm:"type:varchar(100);unique" json:"name"` // Name 是角色的名称，它在系统中是唯一的
	Type      string    `gorm:"type:varchar(100)" json:"type"`        // Type 是角色的类型，它标识了角色的职能
	Desc      string    `gorm:"type:varchar(255)" json:"desc"`        // Desc 是对角色的描述
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`              // CreatedAt 记录了角色创建的时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`              // UpdatedAt 记录了角色最后更新的时间
	IsActive  bool      `gorm:"default:true" json:"is_active"`        // IsActive 标识角色是否是活跃的
}
