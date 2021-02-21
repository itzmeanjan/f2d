package event

import (
	"log"

	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - Creates new event log entry(ies) in DB
func Create(db *gorm.DB, events []*schema.EventLogs) bool {

	if events == nil || len(events) == 0 {

		log.Printf("[❗️] No event log(s) to insert\n")
		return false

	}

	// Wrap write operation inside tx
	if err := db.Transaction(func(tx *gorm.DB) error {

		return tx.Create(events).Error

	}); err != nil {

		log.Printf("[❗️] Failed to insert event log(s) : %s\n", err.Error())
		return false

	}

	return true

}
