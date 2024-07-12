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
	)
	if err != nil {
		panic(err)
	}
}
