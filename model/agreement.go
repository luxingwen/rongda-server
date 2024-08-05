package model

const (
	AgreementTypeSales    = "1"
	AgreementTypePurchase = "2"
	AgreementTypeService  = "3"
	AgreementTypeOther    = "4"
)

type ReqAgreementCreate struct {
	OrderNo string `json:"order_no" binding:"required" gorm:"comment:'订单号'"` // 订单号
	Title   string `json:"title" binding:"required" gorm:"comment:'标题'"`     // 标题
	Date    string `json:"date" binding:"required" gorm:"comment:'日期'"`      // 日期
	Content string `json:"content" binding:"-" gorm:"comment:'内容'"`          // 内容
	Type    string `json:"type" binding:"required" gorm:"comment:'合同类型'"`    // 合同类型 1:销售合同 2:采购合同 3:服务合同 4:其他
	// 附件
	// 合同源文件
	SourceFile string `json:"source_file" binding:"required" gorm:"comment:'合同源文件'"` // 合同源文件
	// 签署位置列表
	SignaturePositionList []SignaturePosition `json:"signature_position_list" binding:"required" gorm:"comment:'签署位置列表'"` // 签署位置列表
}

// 合同
type Agreement struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	// 订单号
	OrderNo string `json:"order_no" gorm:"comment:'订单号'"` // 订单号
	Title   string `json:"title" gorm:"comment:'标题'"`     // 标题
	Date    string `json:"date" gorm:"comment:'日期'"`      // 日期
	Content string `json:"content" gorm:"comment:'内容'"`   // 内容
	Creater string `json:"creater" gorm:"comment:'创建人'"`  // 创建人
	// 合同类型
	Type string `json:"type" gorm:"comment:'合同类型'"` // 合同类型 1:销售合同 2:采购合同 3:服务合同 4:其他
	// 附件
	Attachment string `json:"attachment" gorm:"comment:'附件'"` // 附件

	// 合同源文件
	SourceFile string `json:"source_file" gorm:"comment:'合同源文件'"` // 合同源文件

	// 签署过后的文件
	SignatureFile string `json:"signature_file" gorm:"comment:'签署过后的文件'"` // 签署过后的文件

	// 签署图片
	SignatureImage string `json:"signature_image" gorm:"comment:'签署图片'"` // 签署图片
	// 签署时间
	SignatureTime string `json:"signature_time" gorm:"comment:'签署时间'"` // 签署时间
	// 签署用户
	SignatureUser string `json:"signature_user" gorm:"comment:'签署用户'"` // 签署用户

	// 签署位置
	SignaturePosition string `json:"signature_position" gorm:"comment:'签署位置'"` // 签署位置

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// 签署位置
type SignaturePosition struct {
	// 页数
	Page int `json:"page" gorm:"comment:'页数'"` // 页数
	// x坐标
	X int `json:"x" gorm:"comment:'x坐标'"` // x坐标
	// y坐标
	Y int `json:"y" gorm:"comment:'y坐标'"` // y坐标
	// 宽度
	Width int `json:"width" gorm:"comment:'宽度'"` // 宽度
	// 高度
	Height int `json:"height" gorm:"comment:'高度'"` // 高度
	// 签署用户
	SignatureUser string `json:"signature_user" gorm:"comment:'签署用户'"` // 签署用户
}
