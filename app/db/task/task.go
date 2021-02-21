package task

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - Attempts to create one task entry in DB
func Create(db *gorm.DB, user common.Hash, startBlock uint64, contract string, topics map[uint8]string) bool {

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

	if err := db.Create(task).Error; err != nil {

		log.Printf("[❗️] Failed to create task : %s\n", err.Error())
		return false

	}

	return true

}
