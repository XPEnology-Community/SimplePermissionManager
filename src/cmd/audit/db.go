package audit

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	dsn = "/var/packages/SimplePermissionManager/var/audit.db"
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect audit database, dsn: %s", dsn))
	}

	db.AutoMigrate(&Audit{})
}
