package model

type Agent struct {
	ID           uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"`             // 主键ID
	Uuid         string  `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`  // UUID
	Name         string  `json:"name" gorm:"comment:'企业名称'"`                      // 企业名称
	Address      string  `json:"address" gorm:"comment:'地址'"`                     // 地址
	ContactInfo  string  `json:"contact_info" gorm:"comment:'联系方式'"`              // 联系方式
	BankAccount  string  `json:"bank_account" gorm:"comment:'银行账号'"`              // 银行账号
	CreditStatus string  `json:"credit_status" gorm:"comment:'信用状态'"`             // 信用状态
	Rate         float64 `json:"rate" gorm:"comment:'折扣'"`                        // 费率
	Status       int     `json:"status" gorm:"comment:'状态'"`                      // 状态
	CreatedAt    string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt    string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
