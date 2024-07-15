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
