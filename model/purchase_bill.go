package model

type PurchaseBillReq struct {
	PurchaseOrderNo string  `json:"purchase_order_no" form:"purchase_order_no" binding:"required"` // 采购单号
	StockInOrderNo  string  `json:"stock_in_order_no" form:"stock_in_order_no" binding:"required"` // 入库单号
	SupplierUuid    string  `json:"supplier_uuid" form:"supplier_uuid" binding:"-"`                // 供应商UUID
	BankAccount     string  `json:"bank_account" form:"bank_account" binding:"required"`           // 银行账号
	BankName        string  `json:"bank_name" form:"bank_name" binding:"required"`                 // 银行名称
	BankAccountName string  `json:"bank_account_name" form:"bank_account_name" binding:"required"` // 银行账户名
	Amount          float64 `json:"amount" form:"amount" binding:"required"`                       // 金额
	PaymentDate     string  `json:"payment_date" form:"payment_date" binding:"required"`           // 付款日期
	PaymentMethod   string  `json:"payment_method" form:"payment_method" binding:"required"`       // 付款方式
	Remark          string  `json:"remark" form:"remark"`                                          // 备注
}

type PurchaseBill struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	// 采购单号
	PurchaseOrderNo string `json:"purchase_order_no" gorm:"comment:'采购单号'"` // 采购单号
	// 入库单号
	StockInOrderNo string `json:"stock_in_order_no" gorm:"comment:'入库单号'"` // 入库单号
	// 供应商ID
	SupplierUuid string `json:"supplier_uuid" gorm:"type:char(36);index;comment:'供应商UUID'"` // 供应商UUID
	// 银行账号
	BankAccount string `json:"bank_account" gorm:"comment:'银行账号'"` // 银行账号
	// 银行名称
	BankName string `json:"bank_name" gorm:"comment:'银行名称'"` // 银行名称
	// 银行账户名
	BankAccountName string `json:"bank_account_name" gorm:"comment:'银行账户名'"` // 银行账户名
	// 金额
	Amount float64 `json:"amount" gorm:"comment:'金额'"` // 金额
	// 付款日期
	PaymentDate string `json:"payment_date" gorm:"comment:'付款日期'"` // 付款日期
	// 付款方式
	PaymentMethod string `json:"payment_method" gorm:"comment:'付款方式'"` // 付款方式 1:现金 2:转账 3:支票 4:其他
	// 备注
	Remark string `json:"remark" gorm:"comment:'备注'"` // 备注

	// 财务人员
	FinanceStaff string `json:"finance_staff" gorm:"comment:'财务人员'"` // 财务人员
	// 财务审核日期
	FinanceAuditDate string `json:"finance_audit_date" gorm:"comment:'财务审核日期'"` // 财务审核日期
	// 财务审核状态
	FinanceAuditStatus int `json:"finance_audit_status" gorm:"comment:'财务审核状态'"` // 财务审核状态 1:待审核 2:已审核 3:已驳回
	// 财务审核备注
	FinanceAuditRemark string `json:"finance_audit_remark" gorm:"comment:'财务审核备注'"` // 财务审核备注

	BillUuid string `json:"bill_uuid" gorm:"type:char(36);index;comment:'发票UUID'"` // 发票UUID
	// 状态
	Status    int    `json:"status" gorm:"comment:'状态'"`                      // 状态 1:待付款 2:已付款 3:已取消
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
