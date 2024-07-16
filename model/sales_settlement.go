package model

type SalesSettlementReq struct {
	OrderUuid string `json:"order_uuid" form:"order_uuid" binding:"required"` // 订单UUID
	// 付款方式
	PaymentMethod string `json:"payment_method" form:"payment_method" binding:"required"` // 付款方式
	// 付款日期
	PaymentDate string `json:"payment_date" form:"payment_date" binding:"required"` // 付款日期
	// 金额
	Amount float64 `json:"amount" form:"amount" binding:"required"` // 金额
	// 备注
	Remark string `json:"remark" form:"remark"` // 备注
	// 付款凭证
	PaymentVoucher string `json:"payment_voucher" form:"payment_voucher"` // 付款凭证

	// 财务人员
	FinanceStaff string `json:"finance_staff" form:"finance_staff"` // 财务人员
	// 财务审核日期
	FinanceAuditDate string `json:"finance_audit_date" form:"finance_audit_date"` // 财务审核日期
	// 财务审核状态
	FinanceAuditStatus int `json:"finance_audit_status" form:"finance_audit_status"` // 财务审核状态 1:待审核 2:已审核 3:已驳回
	// 财务审核备注
	FinanceAuditRemark string `json:"finance_audit_remark" form:"finance_audit_remark"` // 财务审核备注
	// 结算状态
	Status  int    `json:"status" form:"status"`   // 结算状态 1:待结算 2:已结算 3:已取消
	Creater string `json:"creater" form:"creater"` // 创建人
}

type SalesSettlement struct {
	ID        uint   `json:"id" gorm:"primaryKey;comment:'主键ID'"`                    // 主键ID
	Uuid      string `json:"uuid" gorm:"type:char(36);index;comment:'UUID'"`         // UUID
	OrderUuid string `json:"order_uuid" gorm:"type:char(36);index;comment:'订单UUID'"` // 订单UUID
	// 付款方式
	PaymentMethod string `json:"payment_method" gorm:"comment:'付款方式'"` // 付款方式 1:现金 2:转账 3:支票 4:其他
	// 付款日期
	PaymentDate string `json:"payment_date" gorm:"comment:'付款日期'"` // 付款日期
	// 金额
	Amount float64 `json:"amount" gorm:"comment:'金额'"` // 金额
	// 备注
	Remark string `json:"remark" gorm:"comment:'备注'"` // 备注
	// 付款凭证
	PaymentVoucher string `json:"payment_voucher" gorm:"comment:'付款凭证'"` // 付款凭证

	// 财务人员
	FinanceStaff string `json:"finance_staff" gorm:"comment:'财务人员'"` // 财务人员
	// 财务审核日期
	FinanceAuditDate string `json:"finance_audit_date" gorm:"comment:'财务审核日期'"` // 财务审核日期
	// 财务审核状态
	FinanceAuditStatus int `json:"finance_audit_status" gorm:"comment:'财务审核状态'"` // 财务审核状态 1:待审核 2:已审核 3:已驳回
	// 财务审核备注
	FinanceAuditRemark string `json:"finance_audit_remark" gorm:"comment:'财务审核备注'"` // 财务审核备注
	// 结算状态
	Status    int    `json:"status" gorm:"comment:'结算状态'"`                    // 结算状态 1:待结算 2:已结算 3:已取消
	Creater   string `json:"creater" gorm:"comment:'创建人'"`                    // 创建人
	CreatedAt string `json:"created_at" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
