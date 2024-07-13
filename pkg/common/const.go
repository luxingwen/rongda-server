package common

type TeamType string // 团队类型

const (
	TeamTypeOrganization TeamType = "organization" // 组织
	TeamTypeDepartment   TeamType = "department"   // 部门
	TeamTypeCustomer     TeamType = "customer"     // 客户
	TeamTypeSupplier     TeamType = "supplier"     // 供应商
	TeamTypeAgent        TeamType = "agent"        // 代理商
)
