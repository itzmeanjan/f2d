package user

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - Create new user i.e. APIKey insertion is actually being
// performed here
func Create(db *gorm.DB, apiKey common.Hash, address common.Address) bool {

	if err := db.Transaction(func(tx *gorm.DB) error {

		return tx.Create(&schema.Users{
			APIKey:    apiKey.Hex(),
			Address:   address.Hex(),
			TimeStamp: time.Now().UTC(),
		}).Error

	}); err != nil {

		log.Printf("[❗️] Failed to create user : %s\n", err.Error())
		return false

	}

	return true

}
