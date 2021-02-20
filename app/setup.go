package app

import (
	"log"
	"path/filepath"

	"github.com/itzmeanjan/f2d/app/config"
	"github.com/itzmeanjan/f2d/app/db"
)

// SetUp - Do basic ground set up work, required for
// running `f2d` on this machine
func SetUp() bool {

	path, err := filepath.Abs("./.env")
	if err != nil {

		log.Printf("[❗️] Failed to find `.env` : %s\n", err.Error())
		return false

	}

	if err := config.Read(path); err != nil {

		log.Printf("[❗️] Failed to read `.env` : %s\n", err.Error())
		return false

	}

	db.Connect()

	return true

}
