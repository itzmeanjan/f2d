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
			CreateBatchSize:        10, // when performing multi DML ops, they will be splitted into batch size of 10
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true, // all db writing to be wrapped inside transaction manually
		})
	if err != nil {

		log.Printf("[❗️] Failed to connect to db : %s\n", err.Error())
		return nil

	}

	_db.AutoMigrate(&Users{}, &Tasks{}, &EventLogs{}, &TaskResults{})

	return _db

}
