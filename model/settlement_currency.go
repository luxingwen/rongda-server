package model

const (
	SettlementCurrencyStatusEnabled  = 1 // 启用
	SettlementCurrencyStatusDisabled = 0 // 未启用
)

type SettlementCurrency struct {
	ID           uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"`             // 主键ID
	Uuid         string  `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`  // UUID
	Name         string  `json:"name" gorm:"comment:'名称'"`                        // 名称
	Code         string  `json:"code" gorm:"comment:'代码'"`                        // 代码
	ExchangeRate float64 `json:"exchange_rate" gorm:"comment:'汇率'"`               // 汇率
	Status       int     `json:"status" gorm:"comment:'状态'"`                      // 状态 0:未启用 1:启用
	CreatedAt    string  `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt    string  `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
