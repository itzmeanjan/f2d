package result

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// Create - Attempt to push job result entry(ies), when
// new block gets mined & association of event log in block
// and user submitted job(s) is found
func Create(db *gorm.DB, blockHash common.Hash, index uint, taskIds []string) bool {

	if len(taskIds) == 0 {

		log.Printf("[❗️] No task id(s) given\n")
		return false

	}

	// Wrap insertion op inside tx
	if err := db.Transaction(func(tx *gorm.DB) error {

		// Allocating buffer in a single go
		results := make([]*schema.TaskResults, 0, len(taskIds))

		for _, v := range taskIds {

			results = append(results, &schema.TaskResults{
				BlockHash: blockHash.Hex(),
				Index:     index,
				ID:        v,
			})

		}

		// Performing batch insert operation
		return tx.Create(results).Error

	}); err != nil {

		log.Printf("[❗️] Failed to insert task result(s) : %s\n", err.Error())
		return false

	}

	return true

}
