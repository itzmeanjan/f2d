package task

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - ...
func Create(db *gorm.DB, user common.Hash, startBlock uint64, contract common.Address, topics []common.Hash) bool {

	var task *schema.Tasks

	switch len(topics) {
	case 0:

		task = &schema.Tasks{
			Client:     user.Hex(),
			StartBlock: startBlock,
			Contract:   contract.Hex(),
			TimeStamp:  time.Now().UTC(),
		}

	case 1:

		task = &schema.Tasks{
			Client:     user.Hex(),
			StartBlock: startBlock,
			Contract:   contract.Hex(),
			Topic0:     topics[0].Hex(),
			TimeStamp:  time.Now().UTC(),
		}

	case 2:

		task = &schema.Tasks{
			Client:     user.Hex(),
			StartBlock: startBlock,
			Contract:   contract.Hex(),
			Topic0:     topics[0].Hex(),
			Topic1:     topics[1].Hex(),
			TimeStamp:  time.Now().UTC(),
		}

	case 3:

		task = &schema.Tasks{
			Client:     user.Hex(),
			StartBlock: startBlock,
			Contract:   contract.Hex(),
			Topic0:     topics[0].Hex(),
			Topic1:     topics[1].Hex(),
			Topic2:     topics[2].Hex(),
			TimeStamp:  time.Now().UTC(),
		}

	case 4:

		task = &schema.Tasks{
			Client:     user.Hex(),
			StartBlock: startBlock,
			Contract:   contract.Hex(),
			Topic0:     topics[0].Hex(),
			Topic1:     topics[1].Hex(),
			Topic2:     topics[2].Hex(),
			Topic3:     topics[3].Hex(),
			TimeStamp:  time.Now().UTC(),
		}

	}

	if err := db.Create(task).Error; err != nil {

		log.Printf("[❗️] Failed to create task : %s\n", err.Error())
		return false

	}

	return true

}
