package model

// Team 结构体定义了团队的基础信息
type Team struct {
	Id        uint   `gorm:"primary_key" json:"id"`                // ID 是团队的主键
	UUID      string `gorm:"type:char(36);unique" json:"uuid"`     // UUID 是团队的唯一标识符
	Name      string `gorm:"type:varchar(100);unique" json:"name"` // Name 是团队的名称，它在系统中是唯一的
	Desc      string `gorm:"type:varchar(255)" json:"desc"`        // Desc 是对团队的描述
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"`     // CreatedAt 记录了团队创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"`     // UpdatedAt 记录了团队最后更新的时间
	IsActive  bool   `gorm:"default:true" json:"is_active"`        // IsActive 标识团队是否是活跃的
	Creater   string `gorm:"type:char(36)" json:"creater"`         // Creater 标识团队的创建者
}

type ResTeam struct {
	TeamUuid string `json:"team_uuid"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Category string `json:"category"`
}

const (
	// 客户类型
	TeamCategoryCustomer = "customer"
	// 供应商类型
	TeamCategorySupplier = "supplier"
	// 代理商类型
	TeamCategoryAgent = "agent"
)

type TeamRef struct {
	Id        uint   `gorm:"primary_key" json:"id"`                    // ID 是团队的主键
	TeamUuid  string `gorm:"type:char(36);unique" json:"team_uuid"`    // UUID 是团队的唯一标识符
	Category  string `gorm:"type:varchar(100);unique" json:"category"` // Name 是类型
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"`         // CreatedAt 记录了团队创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"`         // UpdatedAt 记录了团队最后更新的时间
}

// 团队邀请信息
type TeamInvite struct {
	Id       uint   `gorm:"primary_key" json:"id"`                 // ID 是团队的主键
	TeamUuid string `gorm:"type:char(36);unique" json:"team_uuid"` // UUID 是团队的唯一标识符
	// 邀请码
	InviteCode string `gorm:"type:varchar(100);unique" json:"invite_code"`
	// 邀请人
	Inviter string `gorm:"type:char(36)" json:"inviter"`
}

type ReqUpdateInviteStatus struct {
	InviteCode string `json:"invite_code"`
	Status     int    `json:"status"`
}

type ReqInviteCodeParam struct {
	InviteCode string `json:"invite_code"`
}
