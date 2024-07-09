package model

import (
	"time"
)

// 采购管理模块
type ProcurementOrder struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Contract    string    `json:"contract" gorm:"comment:'合同信息'"`      // 合同信息
	ProductList string    `json:"product_list" gorm:"comment:'商品清单'"`  // 商品清单
	Tax         float64   `json:"tax" gorm:"comment:'税费'"`             // 税费
	Deposit     float64   `json:"deposit" gorm:"comment:'定金'"`         // 定金
	Supplier    string    `json:"supplier" gorm:"comment:'供应商'"`       // 供应商
	OrderDate   time.Time `json:"order_date" gorm:"comment:'订单日期'"`    // 订单日期
	Remarks     string    `json:"remarks" gorm:"comment:'备注'"`         // 备注
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了创建的时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了信息最后更新的时间
}

type ProcurementArrival struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	OrderID     uint      `json:"order_id" gorm:"comment:'关联的采购订单ID'"` // 关联的采购订单ID
	Quantity    int       `json:"quantity" gorm:"comment:'到货数量'"`      // 到货数量
	ArrivalDate time.Time `json:"arrival_date" gorm:"comment:'到货日期'"`  // 到货日期
	Inspector   string    `json:"inspector" gorm:"comment:'检验员'"`      // 检验员
	Batch       string    `json:"batch" gorm:"comment:'批次号'"`          // 批次号
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了创建的时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了信息最后更新的时间
}

type Payment struct {
	ID               uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"`     // 主键ID
	OrderID          uint      `json:"order_id" gorm:"comment:'关联的采购订单ID'"`     // 关联的采购订单ID
	Amount           float64   `json:"amount" gorm:"comment:'支付金额'"`            // 支付金额
	BankAccount      string    `json:"bank_account" gorm:"comment:'银行账号'"`      // 银行账号
	PaymentDate      time.Time `json:"payment_date" gorm:"comment:'支付日期'"`      // 支付日期
	Description      string    `json:"description" gorm:"comment:'描述'"`         // 描述
	ConfirmedBy      string    `json:"confirmed_by" gorm:"comment:'确认人'"`       // 确认人
	ConfirmationDate time.Time `json:"confirmation_date" gorm:"comment:'确认日期'"` // 确认日期
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`        // CreatedAt 记录了创建的时间
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`        // UpdatedAt 记录了信息最后更新的时间
}

type ThirdPartyWarehouse struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	ProductInfo string    `json:"product_info" gorm:"comment:'商品信息'"`  // 商品信息
	Quantity    int       `json:"quantity" gorm:"comment:'数量'"`        // 数量
	Location    string    `json:"location" gorm:"comment:'位置'"`        // 位置
	ArrivalDate time.Time `json:"arrival_date" gorm:"comment:'到货日期'"`  // 到货日期
	StorageFee  float64   `json:"storage_fee" gorm:"comment:'存储费'"`    // 存储费
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了创建的时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了信息最后更新的时间
}

// 仓库管理模块
type Warehouse struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Location  string    `json:"location" gorm:"comment:'仓库位置'"`      // 仓库位置
	Capacity  int       `json:"capacity" gorm:"comment:'容量'"`        // 容量
	Manager   string    `json:"manager" gorm:"comment:'管理员'"`        // 管理员
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了创建的时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了信息最后更新的时间
}

type InventoryCheck struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	WarehouseID uint      `json:"warehouse_id" gorm:"comment:'仓库ID'"`  // 仓库ID
	ProductList string    `json:"product_list" gorm:"comment:'商品清单'"`  // 商品清单
	CheckDate   time.Time `json:"check_date" gorm:"comment:'盘点日期'"`    // 盘点日期
	CheckedBy   string    `json:"checked_by" gorm:"comment:'盘点人'"`     // 盘点人
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`    // CreatedAt 记录了创建的时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`    // UpdatedAt 记录了信息最后更新的时间
}

type DamageReport struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	ProductID   uint      `json:"product_id" gorm:"comment:'商品ID'"`    // 商品ID
	Description string    `json:"description" gorm:"comment:'描述'"`     // 描述
	Quantity    int       `json:"quantity" gorm:"comment:'数量'"`        // 数量
	ReportDate  time.Time `json:"report_date" gorm:"comment:'报告日期'"`   // 报告日期
	ReportedBy  string    `json:"reported_by" gorm:"comment:'报告人'"`    // 报告人
}

// 销售管理模块
type SalesOrder struct {
	ID            uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"`  // 主键ID
	Contract      string    `json:"contract" gorm:"comment:'合同信息'"`       // 合同信息
	ProductList   string    `json:"product_list" gorm:"comment:'商品清单'"`   // 商品清单
	Tax           float64   `json:"tax" gorm:"comment:'税费'"`              // 税费
	Deposit       float64   `json:"deposit" gorm:"comment:'定金'"`          // 定金
	PaymentMethod string    `json:"payment_method" gorm:"comment:'支付方式'"` // 支付方式
	OrderDate     time.Time `json:"order_date" gorm:"comment:'订单日期'"`     // 订单日期
	Remarks       string    `json:"remarks" gorm:"comment:'备注'"`          // 备注
}

type Shipment struct {
	ID           uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"`       // 主键ID
	SalesOrderID uint      `json:"sales_order_id" gorm:"comment:'关联的销售订单ID'"` // 关联的销售订单ID
	Quantity     int       `json:"quantity" gorm:"comment:'发货数量'"`            // 发货数量
	ShipmentDate time.Time `json:"shipment_date" gorm:"comment:'发货日期'"`       // 发货日期
	Recipient    string    `json:"recipient" gorm:"comment:'收件人'"`            // 收件人
	Address      string    `json:"address" gorm:"comment:'地址'"`               // 地址
}

type Installment struct {
	ID           uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"`       // 主键ID
	SalesOrderID uint      `json:"sales_order_id" gorm:"comment:'关联的销售订单ID'"` // 关联的销售订单ID
	Amount       float64   `json:"amount" gorm:"comment:'分期金额'"`              // 分期金额
	BankAccount  string    `json:"bank_account" gorm:"comment:'银行账号'"`        // 银行账号
	PaymentDate  time.Time `json:"payment_date" gorm:"comment:'支付日期'"`        // 支付日期
	Description  string    `json:"description" gorm:"comment:'描述'"`           // 描述
}

// 财务管理模块
type Invoice struct {
	ID           uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"`       // 主键ID
	SalesOrderID uint      `json:"sales_order_id" gorm:"comment:'关联的销售订单ID'"` // 关联的销售订单ID
	InvoiceDate  time.Time `json:"invoice_date" gorm:"comment:'发票日期'"`        // 发票日期
	Amount       float64   `json:"amount" gorm:"comment:'金额'"`                // 金额
	Remarks      string    `json:"remarks" gorm:"comment:'备注'"`               // 备注
}

type Receipt struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	PaymentID   uint      `json:"payment_id" gorm:"comment:'关联的支付ID'"` // 关联的支付ID
	ReceiptDate time.Time `json:"receipt_date" gorm:"comment:'收据日期'"`  // 收据日期
	Amount      float64   `json:"amount" gorm:"comment:'金额'"`          // 金额
	Remarks     string    `json:"remarks" gorm:"comment:'备注'"`         // 备注
}

// 统计分析模块
type Analysis struct {
	ID          uint      `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	ReportType  string    `json:"report_type" gorm:"comment:'报告类型'"`   // 报告类型
	Data        string    `json:"data" gorm:"comment:'数据'"`            // 数据
	GeneratedAt time.Time `json:"generated_at" gorm:"comment:'生成日期'"`  // 生成日期
}

// 资料管理模块
type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Name        string  `json:"name" gorm:"comment:'商品名称'"`          // 商品名称
	SKU         string  `json:"sku" gorm:"comment:'SKU'"`            // SKU
	Description string  `json:"description" gorm:"comment:'描述'"`     // 描述
	Price       float64 `json:"price" gorm:"comment:'价格'"`           // 价格
	Cost        float64 `json:"cost" gorm:"comment:'成本'"`            // 成本
}

type Supplier struct {
	ID          uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Name        string `json:"name" gorm:"comment:'企业名称'"`          // 企业名称
	Country     string `json:"country" gorm:"comment:'国家'"`         // 国家
	ContactInfo string `json:"contact_info" gorm:"comment:'联系方式'"`  // 联系方式
}

type Agent struct {
	ID          uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Name        string `json:"name" gorm:"comment:'企业名称'"`          // 企业名称
	Address     string `json:"address" gorm:"comment:'地址'"`         // 地址
	ContactInfo string `json:"contact_info" gorm:"comment:'联系方式'"`  // 联系方式
}

type Customer struct {
	ID           uint    `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Name         string  `json:"name" gorm:"comment:'企业名称'"`          // 企业名称
	Address      string  `json:"address" gorm:"comment:'地址'"`         // 地址
	ContactInfo  string  `json:"contact_info" gorm:"comment:'联系方式'"`  // 联系方式
	BankAccount  string  `json:"bank_account" gorm:"comment:'银行账号'"`  // 银行账号
	CreditStatus string  `json:"credit_status" gorm:"comment:'信用状态'"` // 信用状态
	Discount     float64 `json:"discount" gorm:"comment:'折扣'"`        // 折扣
}

// 员工管理模块
type Employee struct {
	ID          uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"` // 主键ID
	Department  string `json:"department" gorm:"comment:'部门'"`      // 部门
	Role        string `json:"role" gorm:"comment:'职位'"`            // 职位
	Name        string `json:"name" gorm:"comment:'姓名'"`            // 姓名
	ContactInfo string `json:"contact_info" gorm:"comment:'联系方式'"`  // 联系方式
	Permissions string `json:"permissions" gorm:"comment:'权限管理'"`   // 权限管理
}
