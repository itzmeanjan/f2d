package user

import (
	"encoding/json"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

// GenerateNewAPIKey - Given user address, attempts to generate a new APIKey
// for that user, by applying keccak256 hash on top of a JSON serialized
// message
//
// In that message of certain structure, current UNIX timestamp with
// nano-second level precision is also used, to reduce chance
// of collision
//
// What it denotes, is no repeated APIKey will ( in sometime soon ) ever be generated
func GenerateNewAPIKey(db *gorm.DB, address common.Address) []byte {

	msg := &struct {
		Address common.Address `json:"address"`
		Nonce   uint64         `json:"nonce"`
		Time    int64          `json:"time"`
	}{
		Address: address,
		Nonce:   GetNonce(db, address) + 1,
		Time:    time.Now().UTC().UnixNano(),
	}

	data, err := json.Marshal(msg)
	if err != nil {

		log.Printf("[❗️] Failed to serialize APIKey generation message : %s\n", err.Error())
		return nil

	}

	return crypto.Keccak256(data)

}
