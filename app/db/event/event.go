package event

import (
	"log"

	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - Creates new event log entry in DB
func Create(db *gorm.DB, event *schema.EventLogs) bool {

	if event == nil {

		log.Printf("[❗️] No event log to insert\n")
		return false

	}

	// Wrap write operation inside tx
	if err := db.Transaction(func(tx *gorm.DB) error {

		return tx.Create(event).Error

	}); err != nil {

		log.Printf("[❗️] Failed to insert event log : %s\n", err.Error())
		return false

	}

	return true

}
