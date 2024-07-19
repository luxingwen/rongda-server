package model

type Customer struct {
	ID      uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid    string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	Name    string `json:"name" gorm:"comment:'企业名称'"`                     // 企业名称
	Address string `json:"address" gorm:"comment:'地址'"`                    // 地址
	// 联系人
	ContactPerson string `json:"contact_person" gorm:"comment:'联系人'"` // 联系人
	ContactInfo   string `json:"contact_info" gorm:"comment:'联系方式'"`  // 联系方式
	BankAccount   string `json:"bank_account" gorm:"comment:'银行账号'"`  // 银行账号
	// 开户行
	BankName     string `json:"bank_name" gorm:"comment:'开户行'"`      // 开户行
	CreditStatus string `json:"credit_status" gorm:"comment:'信用状态'"` // 信用状态
	// 信用分
	CreditScore int     `json:"credit_score" gorm:"comment:'信用分'"`               // 信用分
	Discount    float64 `json:"discount" gorm:"comment:'折扣'"`                    // 折扣
	Status      int     `json:"status" gorm:"comment:'状态'"`                      // 状态 1：正常 2：停用
	CreatedAt   string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt   string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	IsDeleted   int     `json:"is_deleted" gorm:"comment:'是否删除'"`                // 是否删除 1:删除 0:未删除
}
