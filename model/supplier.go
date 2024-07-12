package model

type Supplier struct {
	ID                 uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"`             // 主键ID
	Uuid               string  `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`  // UUID
	Name               string  `json:"name" gorm:"comment:'企业名称'"`                      // 企业名称
	Address            string  `json:"address" gorm:"comment:'地址'"`                     // 地址
	CountryNo          string  `json:"country_no" gorm:"comment:'国家厂号'"`                // 国家
	Phone              string  `json:"contact_info" gorm:"comment:'联系方式'"`              // 联系方式
	SettlementCurrency string  `json:"settlement_currency" gorm:"comment:'结算币种'"`       // 结算币种
	DepositRate        float64 `json:"deposit_rate" gorm:"comment:'定金比率'"`              // 定金比率
	Status             int     `json:"status" gorm:"comment:'状态'"`                      // 状态 0:未启用 1:启用
	CreatedAt          string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt          string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
