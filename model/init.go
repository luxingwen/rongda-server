package model

import "gorm.io/gorm"

func MigrateDbTable(db *gorm.DB) {
	err := db.AutoMigrate(
		&AppPermission{},
		&API{},
		&App{},
		&Log{},
		&Menu{},
		&Role{},
		&User{},
		&RoleMenuPermission{},
		&UserRole{},
		&Team{},
		&TeamMember{},
		&VerificationCode{},
		&Customer{},
		&Agent{},
		&Supplier{},
		&SettlementCurrency{},
		&Sku{},
		&Department{},
		&Product{},
		&Storehouse{},
		&StorehouseInbound{},
		&StorehouseInboundDetail{},
		&StorehouseProduct{},
		&StorehouseOutbound{},
		&StorehouseOutboundDetail{},
		&StorehouseInventoryCheck{},
		&StorehouseInventoryCheckDetail{},
		&PurchaseOrder{},
		&PurchaseOrderItem{},
		&Agreement{},
		&PurchaseArrival{},
		&PurchaseArrivalItem{},
		&PurchaseBill{},
		&SalesOrder{},
		&SalesOrderItem{},
		&SalesOutOfStock{},
		&SalesOutOfStockItem{},
		&SalesSettlement{},
		&Bill{},
		&StorehouseProductOpLog{},
		&ProductCategory{},
		&ProductManage{},
		&SysBankInfo{},
		&SysLoginLog{},
	)
	if err != nil {
		panic(err)
	}
}
