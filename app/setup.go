package app

import (
	"log"
	"path/filepath"

	"github.com/itzmeanjan/f2d/app/config"
	"github.com/itzmeanjan/f2d/app/data"
	"github.com/itzmeanjan/f2d/app/db"
)

// SetUp - Do basic ground set up work, required for
// running `f2d` on this machine
func SetUp() *data.Resources {

	path, err := filepath.Abs("./.env")
	if err != nil {

		log.Printf("[❗️] Failed to find `.env` : %s\n", err.Error())
		return nil

	}

	if err := config.Read(path); err != nil {

		log.Printf("[❗️] Failed to read `.env` : %s\n", err.Error())
		return nil

	}

	_db := db.Connect()
	if _db == nil {
		return nil
	}

	resources := data.Resources{
		DB: _db,
	}

	return &resources

}
