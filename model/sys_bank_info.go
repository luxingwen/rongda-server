package model

type SysBankInfo struct {
	ID        uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Uuid      string `json:"uuid" gorm:"comment:'银行唯一标识'"`        // 银行唯一标识
	Name      string `json:"name" gorm:"comment:'银行名称'"`          // 银行名称
	Sort      int    `json:"sort" gorm:"comment:'排序'"`            // 排序
	Status    int    `json:"status" gorm:"comment:'状态'"`          // 状态 1:启用 2:禁用
	CreatedAt string `json:"created_at" gorm:"comment:'创建时间'"`    // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"comment:'更新时间'"`    // 更新时间
}
