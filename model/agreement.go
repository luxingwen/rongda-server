package model

const (
	AgreementTypeSales    = "1"
	AgreementTypePurchase = "2"
	AgreementTypeService  = "3"
	AgreementTypeOther    = "4"
)

// 合同
type Agreement struct {
	ID      uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid    string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	Title   string `json:"title" gorm:"comment:'标题'"`                      // 标题
	Date    string `json:"date" gorm:"comment:'日期'"`                       // 日期
	Content string `json:"content" gorm:"comment:'内容'"`                    // 内容
	Creater string `json:"creater" gorm:"comment:'创建人'"`                   // 创建人
	// 合同类型
	Type string `json:"type" gorm:"comment:'合同类型'"` // 合同类型 1:销售合同 2:采购合同 3:服务合同 4:其他
	// 附件
	Attachment string `json:"attachment" gorm:"comment:'附件'"`                  // 附件
	CreatedAt  string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt  string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
