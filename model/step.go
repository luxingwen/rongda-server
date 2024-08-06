package model

const (
	// 等待
	StepStatusWait = "wait"
	// 处理中
	StepStatusProcess = "process"
	// 完成
	StepStatusFinish = "finish"
	// 拒绝
	StepStatusReject = "reject"
	// error
	StepStatusError = "error"
)

const (
	// 销售订单类型
	StepRefTypeSalesOrder = "sales_order"
)

const (
	// 创建步骤
	StepTypeCreate = "create"
	// 详情步骤
	StepTypeDetail = "detail"
)

type Step struct {
	ID int `gorm:"primary_key" json:"id"`
	// 步骤链ID
	ChainId     string `gorm:"type:char(36);index" json:"chain_id"`   // 步骤链ID
	Uuid        string `gorm:"type:char(36);unique" json:"uuid"`      // 用户唯一标识
	Title       string `gorm:"type:varchar(100);" json:"title"`       // 步骤名称
	Description string `gorm:"type:varchar(100);" json:"description"` // 步骤内容
	RefId       string `gorm:"type:char(36);" json:"ref_id"`          // 关联ID
	RefType     string `gorm:"type:varchar(100);" json:"ref_type"`    // 关联类型
	StepType    string `gorm:"type:varchar(100);" json:"step_type"`   // 步骤类型
	Status      string `gorm:"type:varchar(100);" json:"status"`      // 步骤状态
	StepOrder   int    `gorm:"type:varchar(100);" json:"step_order"`  // 步骤顺序
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"`      // 创建时间
	UpdatedAt   string `gorm:"autoUpdateTime" json:"updated_at"`      // 更新时间
}

const (
	StepChainRefTypeSalesOrder    = "sales_order"    // 销售订单
	StepChainRefTypePurchaseOrder = "purchase_order" // 采购订单
)

const (
	StepChainStatusWait = "wait" // 等待
	// 处理中
	StepChainStatusProcessing = "process"
	// 完成
	StepChainStatusFinish = "finish"
	// 拒绝
	StepChainStatusReject = "reject"
)

// 步骤链
type StepChain struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Uuid        string `gorm:"type:char(36);unique" json:"uuid"`       // 用户唯一标识
	RefId       string `gorm:"type:char(36);unique" json:"ref_id"`     // 关联ID
	RefType     string `gorm:"type:varchar(100);" json:"ref_type"`     // 关联类型
	ChainName   string `gorm:"type:varchar(100);" json:"chain_name"`   // 链名称
	ChainType   string `gorm:"type:varchar(100);" json:"chain_type"`   // 链类型
	ChainStatus string `gorm:"type:varchar(100);" json:"chain_status"` // 链状态
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"`       // 创建时间
	UpdatedAt   string `gorm:"autoUpdateTime" json:"updated_at"`       // 更新时间
}

var SalesSteps = []Step{
	{Title: "创建订单", StepOrder: 1},
	{Title: "订单确认", Description: "客户确认订单", StepOrder: 2},
	{Title: "创建合同", Description: "创建订单合同", StepOrder: 3},
	{Title: "签署合同", StepOrder: 4},
	{Title: "创建定金合同", StepOrder: 5},
	{Title: "签署定金合同", StepOrder: 6},
	{Title: "支付定金", StepOrder: 7},
	{Title: "更新单据信息", StepOrder: 8},
	{Title: "船期更新", StepOrder: 9},
	{Title: "创建尾款合同", StepOrder: 10},
	{Title: "签署尾款合同", StepOrder: 11},
	{Title: "支付尾款", StepOrder: 12},
	{Title: "等待货物到港清关", StepOrder: 13},
	{Title: "货物流向", StepOrder: 14},
	{Title: "货物海关放行", StepOrder: 15},
	{Title: "入库或倒柜直提", StepOrder: 16},
	{Title: "预约提货", StepOrder: 17},
	{Title: "账单结算", StepOrder: 18},
	{Title: "账单确认", StepOrder: 19},
	{Title: "货款支付", StepOrder: 20},
	{Title: "货物放行", StepOrder: 21},
	{Title: "完成", StepOrder: 22},
}
