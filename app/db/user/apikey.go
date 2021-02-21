package user

import (
	"github.com/ethereum/go-ethereum/common"
	schema "github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// GetNonce - Retrieves how many APIKey(s) are created
// by this user address
func GetNonce(db *gorm.DB, address common.Address) uint64 {

	var count int64

	if err := db.Model(&schema.Users{}).Where("address = ?", address.Hex()).Count(&count).Error; err != nil {

		// This is due to no APIKey(s) have ever been
		// created by this user address on `f2d`
		return 0

	}

	return uint64(count)

}
