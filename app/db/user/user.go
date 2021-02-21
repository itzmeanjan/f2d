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

// mutateState - Attempts to mutate state of existing APIKey
func mutateState(db *gorm.DB, apiKey common.Hash, state bool) bool {

	// Wrap mutation operation inside tx
	if err := db.Transaction(func(tx *gorm.DB) error {

		return tx.Model(&schema.Users{}).
			Where("apikey = ?", apiKey.Hex()).
			Update("enabled = ?", state).Error

	}); err != nil {

		log.Printf("[❗️] Failed to mutate state of APIKey : %s\n", err.Error())
		return false

	}

	return true

}

// Enable - Given exisiting ( disabled ) APIKey, attempts to activate it
func Enable(db *gorm.DB, apiKey common.Hash) bool {

	return mutateState(db, apiKey, true)

}

// Disable - Given exisiting ( enabled ) APIKey, attempts to deactivate it
func Disable(db *gorm.DB, apiKey common.Hash) bool {

	return mutateState(db, apiKey, false)

}
