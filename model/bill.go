package model

type Bill struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID
	// 发票公司
	InvoiceCompany string `json:"invoice_company" gorm:"comment:'发票公司'"` // 发票公司
	// 开票人
	// 申请人
	Applicant string `json:"applicant" gorm:"comment:'申请人'"` // 申请人
	// 发票号
	InvoiceNo string `json:"invoice_no" gorm:"comment:'发票号'"` // 发票号
	// 发票代码
	InvoiceCode string `json:"invoice_code" gorm:"comment:'发票代码'"` // 发票代码
	// 发票类型
	InvoiceType string `json:"invoice_type" gorm:"comment:'发票类型'"` // 发票类型 1:增值税专用发票 2:增值税普通发票
	// 开票日期
	InvoiceDate string `json:"invoice_date" gorm:"comment:'开票日期'"` // 开票日期
	// 金额
	Amount float64 `json:"amount" gorm:"comment:'金额'"` // 金额
	// 税率
	TaxRate float64 `json:"tax_rate" gorm:"comment:'税率'"` // 税率
	// 税额
	TaxAmount float64 `json:"tax_amount" gorm:"comment:'税额'"` // 税额
	// 价税合计
	TotalAmount float64 `json:"total_amount" gorm:"comment:'价税合计'"` // 价税合计
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
	// 状态
	Status    int    `json:"status" gorm:"comment:'状态'"`                      // 状态 1:待付款 2:已付款 3:已取消
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

const (
	// 支付账单类型
	PaymentBillTypeDeposit = "定金" // 定金
	PaymentBillTypeFinal   = "尾款" // 尾款
	// 结算款
	PaymentBillTypeSettlement = "结算款"
)

const (
	// 支付账单状态
	PaymentBillStatusPendingPayment = "待支付" // 待支付
	// 已支付待确认
	PaymentBillStatusPaidPendingConfirm = "已支付待确认" // 已支付待确认
	PaymentBillStatusPaid               = "已支付"    // 已支付
	PaymentBillStatusCancelled          = "已取消"    // 已取消
)

// 支付账单
type PaymentBill struct {
	ID   uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`            // 主键ID
	Uuid string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"` // UUID

	TeamUuid string `json:"team_uuid" gorm:"type:char(36);comment:'团队UUID'"` // 团队UUID
	// 订单号
	OrderNo string `json:"order_no" gorm:"comment:'订单号'"` // 订单号

	// 合同号
	AgreementNo string `json:"agreement_no" gorm:"comment:'合同号'"` // 合同号
	// PI合同号
	PiAgreementNo string `json:"pi_agreement_no" gorm:"comment:'PI合同号'"` // PI合同号

	// 柜号
	CabinetNo string `json:"cabinet_no" gorm:"comment:'柜号'"` // 柜号

	// 类型
	Type string `json:"type" gorm:"comment:'类型'"` // 类型  定金  尾款  全款  结算款 其他

	// 原币金额
	OriginalAmount float64 `json:"original_amount" gorm:"comment:'原币金额'"` // 原币金额
	// 原币币种
	OriginalCurrency string `json:"original_currency" gorm:"comment:'原币币种'"` // 原币币种

	// 汇率
	ExchangeRate float64 `json:"exchange_rate" gorm:"comment:'汇率'"` // 汇率

	// 应付金额
	Amount float64 `json:"amount" gorm:"comment:'应付金额'"` // 应付金额
	// 实际付款金额
	PaymentAmount float64 `json:"payment_amount" gorm:"comment:'实际付款金额'"` // 实际付款金额

	// 未付金额
	UnpaidAmount float64 `json:"unpaid_amount" gorm:"comment:'未付金额'"` // 未付金额

	// 可垫资额度
	AdvanceAmount float64 `json:"advance_amount" gorm:"comment:'可垫资额度'"` // 可垫资额度
	// 是否垫资
	IsAdvance int `json:"is_advance" gorm:"comment:'是否垫资'"` // 是否垫资 1:是 0:否

	// 状态
	Status string `json:"status" gorm:"comment:'状态'"` // 状态 1:待付款 2:已付款 3:已取消

	// 锁汇汇率
	LockExchangeRate float64 `json:"lock_exchange_rate" gorm:"comment:'锁汇汇率'"` // 锁汇汇率

	// 锁汇开始日期
	LockExchangeStartDate string `json:"lock_exchange_start_date" gorm:"comment:'锁汇开始日期'"` // 锁汇开始日期
	// 锁汇结束日期

	LockExchangeEndDate string `json:"lock_exchange_end_date" gorm:"comment:'锁汇结束日期'"` // 锁汇结束日期

	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间

	IsDeleted int `json:"is_deleted" gorm:"comment:'是否删除'"` // 是否删除 1:是 0:否
}
