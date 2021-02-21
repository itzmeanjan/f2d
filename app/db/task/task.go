package task

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - Attempts to create task entry in DB
func Create(db *gorm.DB, user common.Hash, startBlock uint64, contract string, topics map[uint8]string) (bool, string) {

	task := &schema.Tasks{
		Client:     user.Hex(),
		StartBlock: startBlock,
		TimeStamp:  time.Now().UTC(),
	}

	if topics != nil {

		for k, v := range topics {

			switch k {
			case 0:
				task.Topic0 = v
			case 1:
				task.Topic1 = v
			case 2:
				task.Topic2 = v
			case 3:
				task.Topic3 = v

			}
		}

	}

	if len(contract) != 0 {
		task.Contract = contract
	}

	if err := db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(task).Error; err != nil {
			return err
		}

		return nil

	}); err != nil {

		log.Printf("[❗️] Failed to create task : %s\n", err.Error())
		return false, ""

	}

	return true, task.ID

}

// mutateState - Given an existing task id, attempts to mutate
// its activation state ( boolean )
func mutateState(db *gorm.DB, id string, state bool) bool {

	// Wrap mutation operation inside tx
	if err := db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&schema.Tasks{}).Where("id = ?", id).Update("enabled = ?", state).Error; err != nil {
			return err
		}

		return nil

	}); err != nil {

		log.Printf("[❗️] Failed to mutate state of task : %s\n", err.Error())
		return false

	}

	return true

}

// Enable - Given exisiting ( disabled ) task id, attempts to activate it
func Enable(db *gorm.DB, id string) bool {

	return mutateState(db, id, true)

}

// Disable - Given exisiting ( enabled ) task id, attempts to deactivate it
func Disable(db *gorm.DB, id string) bool {

	return mutateState(db, id, false)

}
