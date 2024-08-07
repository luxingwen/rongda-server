package model

// EntrustOrder 委托订单
type EntrustOrder struct {
	ID       uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint(20);not null" json:"id"`
	OrderId  string `gorm:"unique;column:order_id;type:varchar(64);not null" json:"order_id"`  // 订单ID
	UserUuid string `gorm:"index;column:user_uuid;type:varchar(64);not null" json:"user_uuid"` // 用户UUID
	TeamUuid string `gorm:"index;column:team_uuid;type:varchar(64);not null" json:"team_uuid"` // 团队UUID
	Content  string `gorm:"column:content;type:varchar(255);not null" json:"content"`          // 订单内容
	Status   string `gorm:"column:status;type:varchar(32);not null" json:"status"`             // 订单状态 待处理 处理中 已处理 已完成
	// 处理人
	Handler         string `gorm:"column:handler;type:varchar(64);" json:"handler"`
	PurchaseOrderNo string `gorm:"index;column:purchase_order_no;type:varchar(64);" json:"purchase_order_no"` // 采购订单号
	SalesOrderNo    string `gorm:"index;column:sales_order_no;type:varchar(64);" json:"sales_order_no"`       // 销售订单号

	CreatedAt string `gorm:"autoCreateTime" json:"created_at"` // CreatedAt 记录了创建的时间
	UpdatedAt string `gorm:"autoUpdateTime" json:"updated_at"` // UpdatedAt 记录了信息最后更新的时间
	IsDeleted int    `json:"is_deleted" gorm:"comment:'是否删除'"` // 是否删除 1:删除 0:未删除

}
