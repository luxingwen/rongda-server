package model

type Pagination struct {
	PageSize  int    `form:"pageSize" json:"pageSize"`
	Current   int    `form:"current" json:"current"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (p *Pagination) GetOffset() int {
	return p.PageSize * (p.Current - 1)
}

// 角色查询参数
type ReqRoleQueryParam struct {
	Name     string `form:"name"`
	IsActive bool   `form:"is_active"`
	Pagination
}

type ReqUserLogin struct {
	// 用户名或邮箱
	Username string `json:"username" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
}

type ReqUserQueryParam struct {
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Nickname string `json:"nickname"` // 昵称
	Sex      int    `json:"sex"`      // 性别
	Username string `json:"username"` // 用户名
	Status   int    `json:"status"`   // 状态
	Uuid     string `json:"uuid"`     // uuid
	Pagination
}

// 菜单查询参数
type ReqMenuQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 删除用户删除参数
type ReqUserDeleteParam struct {
	Uuid string `json:"uuid" binding:"required"`
}

// 删除菜单参数
type ReqMenuDeleteParam struct {
	Uuid string `json:"uuid" binding:"required"`
}

type ReqApiQueryParam struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	Method string `json:"method"`
	Status int    `json:"status"`
	Pagination
}

// 查询app的参数
type ReqAppQueryParam struct {
	Name   string `json:"name"`    // 名称
	ApiKey string `json:"api_key"` // api_key
	Status int    `json:"status"`  // 状态
	Pagination
}

// uuid参数
type ReqUuidParam struct {
	Uuid  string   `json:"uuid"`
	Uuids []string `json:"uuids"`
}

// api权限参数
type ReqApiPermissionParam struct {
	AppId string `json:"app_id"`
	Pagination
}

type ReqVerificationCodeParam struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type ReqRegisterParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

// 服务查询参数
type ReqServerQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 团队查询参数
type ReqTeamQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 创建团队成员参数
type ReqTeamMemberCreateParam struct {
	TeamUUID string `json:"team_uuid"`
	UserUUID string `json:"user_uuid"`
}

// 团队成员查询参数
type ReqTeamMemberQueryParam struct {
	TeamUUID string `json:"team_uuid"`
	Pagination
}

type ReqTeamMemberUpdateRoleParam struct {
	TeamUUID string `json:"team_uuid"`
	UserUUID string `json:"user_uuid"`
	Role     string `json:"role"`
}

// 创建用户角色参数
type ReqUserRole struct {
	UserUUID  string   `json:"user_uuid"`
	RoleUUIDs []string `json:"role_uuids"`
}

type ReqCustomerQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqAgentQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqSupplierQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqSettlementCurrencyQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqSkuQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqDepartmentQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqProductQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqStorehouseQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqStorehouseInboundQueryParam struct {
	StorehouseUuid           string `json:"storehouse_uuid"`
	PurchaseOrderProductType string `json:"purchase_order_product_type" binding:"-"` // 采购订单物品类型 1：期货 2：现货
	CustomerUuid             string `json:"customer_uuid" binding:"-"`               // 客户uuid
	ProductUuid              string `json:"product_uuid" binding:"-"`                // 商品UUID
	Pagination
}

type ReqStorehouseProductQueryParam struct {
	StorehouseUuid           string `json:"storehouse_uuid"`
	PurchaseOrderProductType string `json:"purchase_order_product_type" binding:"-"` // 采购订单物品类型 1：期货 2：现货
	CustomerUuid             string `json:"customer_uuid" binding:"-"`               // 客户uuid
	ProductUuid              string `json:"product_uuid" binding:"-"`                // 商品UUID
	TeamUuid                 string `json:"team_uuid" binding:"-"`                   // 团队uuid
	Pagination
}

type ReqStorehouseProductSalesOrderQueryParam struct {
	StorehouseUuid string `json:"storehouse_uuid"`
	SalesOrderNo   string `json:"sales_order_no"`
}

type ReqStorehouseOutboundQueryParam struct {
	StorehouseUuid        string `json:"storehouse_uuid"`
	SalesOrderProductType string `json:"sales_order_product_type" binding:"-"` // 销售订单物品类型 1：期货 2：现货
	CustomerUuid          string `json:"customer_uuid" binding:"-"`            // 客户uuid
	ProductUuid           string `json:"product_uuid" binding:"-"`             // 商品UUID
	Pagination
}

type ReqInventoryCheckQueryParam struct {
	CheckOrderNo   string `json:"check_order_no"`  // 盘点单号
	StorehouseUuid string `json:"storehouse_uuid"` // 仓库uuid
	ProductUuid    string `json:"product_uuid"`    // 商品uuid
	CustomerUuid   string `json:"customer_uuid"`   // 客户uuid
	Pagination
}

type ReqPurchaseOrderQueryParam struct {
	OrderNo      string `json:"order_no"`
	Title        string `json:"title"`
	SupplierUuid string `json:"supplier_uuid"`
	CustomerUuid string `json:"customer_uuid"` // 客户uuid  对应小程序的team_id
	Status       string `json:"status"`        // 状态
	Pagination
}

type ReqAgreementQueryParam struct {
	Type     string `json:"type"`
	Status   string `json:"status"`    // 状态 未签署 已签署 已拒绝
	TeamUuid string `json:"team_uuid"` // 团队uuid
	Pagination
}

type ReqPurchaseArrivalQueryParam struct {
	PurchaseOrderNo string `json:"order_no"`
	Pagination
}

type ReqPurchaseBillQueryParam struct {
	PurchaseOrderNo string `json:"purchase_order_no"`
	SupplierUuid    string `json:"supplier_uuid"`
	Pagination
}

type ReqSalesOrderQueryParam struct {
	// 委托单号
	EntrustOrderNo string `json:"entrust_order_no"`
	// 合同号
	AgreementNo string `json:"agreement_no"`

	EtaStartDate string `json:"eta_start_date"` // 预计到货开始日期
	EtaEndDate   string `json:"eta_end_date"`   // 预计到货结束日期

	// 柜号
	CabinetNo string `json:"cabinet_no"`

	OrderNo      string `json:"order_no"`
	CustomerUuid string `json:"customer_uuid"`
	Status       string `json:"status"`
	StartDate    string `json:"start_date"` // 开始日期
	EndDate      string `json:"end_date"`   // 结束日期
	Pagination
}

type ReqSalesOrderUpdateStatusParam struct {
	OrderNo string `json:"order_no"`
	Status  string `json:"status"`
}

type ReqSalesOutOfStockQueryParam struct {
	SalesOrderNo   string `json:"order_no"`
	StorehouseUuid string `json:"storehouse_uuid"`
	Pagination
}

type ReqSalesSettlementQueryParam struct {
	SalesOrderNo string `json:"order_no"`
	OrderUuid    string `json:"order_uuid"`
	Pagination
}

type ReqBillQueryParam struct {
	InvoiceCompany string `json:"order_no"`
	Pagination
}

type ReqStorehouseProductOpLogQueryParam struct {
	Uuid string `json:"uuid"`
	Pagination
}

type ReqProductCategoryQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqSysBankInfoQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqLoginLogQueryParam struct {
	Username string `json:"username"`
	Pagination
}

type ReqAPIQueryParam struct {
	Name   string `json:"name"`
	Module string `json:"module"`
	Status int    `json:"status"`
	Pagination
}

type ReqOpLogQueryParam struct {
	UserName string `form:"user_name"`
	Path     string `form:"path"`
	Method   string `form:"method"`
	Status   int    `form:"status"`
	Pagination
}

type ReqIdParam struct {
	Id int64 `json:"id"`
}

type PurchaseOrderStatusReq struct {
	OrderNo string `json:"order_no"`
	Status  string `json:"status"`
}

type ReqUpdateOrderStatus struct {
	OrderNo string `json:"order_no"`
	Status  string `json:"status"`
}

type ReqPurchaseOrderStatusParam struct {
	StatusList []string `json:"status_list"`
}

type ReqDeleteStorehouseCheckOrderDetail struct {
	Uuid         string `json:"uuid"`
	CheckOrderNo string `json:"check_order_no"`
}

type ReqDepartmentStaffQueryParam struct {
	DepartmentUuid string `json:"department_uuid"`
	Pagination
}

type ReqPermissionQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqUserPermissionQueryParam struct {
	UserUuid       string `json:"user_uuid"`
	PermissionUuid string `json:"permission_uuid"`
	Pagination
}

type ReqPermissionMenuQueryParam struct {
	PermissionUuid string `json:"permission_uuid"`
	MenuUuid       string `json:"menu_uuid"`
	Pagination
}

type ReqMenuAPIQueryParam struct {
	MenuUUID string `json:"menu_uuid"`
	APIUUID  string `json:"api_uuid"`
	Pagination
}

type ReqEntrustOrderQueryParam struct {
	UserUuid  string `json:"user_uuid"`
	TeamUuid  string `json:"team_uuid"`
	Status    string `json:"status"`
	OrderId   string `json:"order_id"`   // 订单ID
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
	Pagination
}

type ReqWxUserQueryParam struct {
	NickName string `json:"nick_name"`
	Pagination
}

type ReqWxUserChangePasswordParam struct {
	Phone    string `json:"phone"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

type ReqWxUserLoginByPasswordParam struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ReqInviteQueryParam struct {
	InviteCode string `json:"invite_code"`
	Pagination
}

type ReqConfigQueryParam struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Pagination
}

type ReqWxUserRealNameAuthParam struct {
	CertificateType   string `json:"certificate_type"`   // 证件类型 身份证 护照 港澳通行证 台湾通行证 其他
	CertificateNumber string `json:"certificate_number"` // 证件号码
	CertificateImage  string `json:"certificate_image"`  // 证件图片
	Name              string `json:"name"`               // 姓名
}

type ReqWxUserUpdateIsRealNameAuthParam struct {
	Uuid       string `json:"uuid"`
	IsRealName int    `json:"is_real_name"`
}

type ReqWxUserUpdateEmailParam struct {
	Email string `json:"email"` // 邮箱
}

// 订单确认
type ReqSalesOrderConfirmParam struct {
	OrderNoList []string `json:"order_no_list"`
	Op          string   `json:"op"` // confirm:确认  cancel:取消
}

type ReqPaymentBillQueryParam struct {
	OrderNo     string `json:"order_no"`
	AgreementNo string `json:"agreement_no"`
	TeamUuid    string `json:"team_uuid"`

	Status string `json:"status"` //
	Type   string `json:"type"`   // 类型  定金  尾款  全款  结算款 其他
	// 是否垫资
	IsAdvance int `json:"is_advance"` // 是否垫资 1:是 0:否
	Pagination
}

type ReqOrderAgreementQueryParam struct {
	OrderNo string `json:"order_no"`
	Type    string `json:"type"`
}

type ReqUpdatePaymentBillStatusParam struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

type ReqPaymentBillOrderStatusPaidComfirm struct {
	Uuid string `json:"uuid"` // 支付账单uuid
	// 实际付款金额
	PaymentAmount float64 `json:"payment_amount"`
}

type ReqSettlementQueryParam struct {
	OrderNo         string `json:"order_no"`
	PurchaseOrderNo string `json:"purchase_order_no"`
	TeamUuid        string `json:"team_uuid"`
	Status          string `json:"status"` // 状态
	Pagination
}

type ReqStorehouseProductOpLogListParam struct {
	TeamUuid       string `json:"team_uuid"`
	StorehouseUuid string `json:"storehouse_uuid"`
	OpTypes        []int  `json:"op_types"` // 操作类型 1:入库 2:出库 3:盘点
	Pagination
}

type ReqUpdatePaymentBillIsAdvanceParam struct {
	// uuid列表
	Uuids []string `json:"uuids"`
	// 是否垫资
	IsAdvance int `json:"is_advance"` // 是否垫资 1:是 0:否
}

type ReqUpdatePurchaseOrderStorehouseParam struct {
	OrderNo         string `json:"order_no"`
	PurchaseOrderNo string `json:"purchase_order_no"`
	StorehouseUuid  string `json:"storehouse_uuid"`
}

type ReqRemittanceBillQueryParam struct {
	OrderNo     string `json:"order_no"`
	AgreementNo string `json:"agreement_no"`
	TeamUuid    string `json:"team_uuid"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Pagination
}

// 更新状态
type ReqUpdateStatus struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

type ReqOrderNoParam struct {

	// 订单号
	OrderNo string `json:"order_no"`
}
