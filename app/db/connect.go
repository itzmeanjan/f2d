package db

import (
	"fmt"
	"log"

	"github.com/itzmeanjan/f2d/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect - Attempt to connect to database, then try to run migration
// which will create all tables if not existing already
func Connect() *gorm.DB {

	_db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.GetDbUser(),
		config.GetDbPassword(),
		config.GetDbHost(),
		config.GetDbPort(),
		config.GetDbName())),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true, // all db writing to be wrapped inside transaction manually
		})
	if err != nil {
		log.Fatalf("[!] Failed to connect to db : %s\n", err.Error())
	}

	_db.AutoMigrate(&Users{}, &Tasks{}, &EventLogs{}, &TaskResults{})

	return _db

}
