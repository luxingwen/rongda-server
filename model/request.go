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
	CustomerUuid string `json:"customer_uuid"`
	Pagination
}

type ReqAgreementQueryParam struct {
	Type string `json:"type"`
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
	OrderNo      string `json:"order_no"`
	CustomerUuid string `json:"customer_uuid"`
	Pagination
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
