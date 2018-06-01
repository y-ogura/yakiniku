package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/y-ogura/yakiniku/account"
	"github.com/y-ogura/yakiniku/db"
)

// Execute run migration
func Execute() {
	odb := db.ConnectDB()
	Migrate(odb)
	odb.Close()
}

// Migrate migrate
func Migrate(odb *gorm.DB) {
	odb.AutoMigrate(
		&account.Account{},
	)
}
