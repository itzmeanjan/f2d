package main

import (
	"log"
	"os"

	"github.com/itzmeanjan/f2d/app"
)

func main() {
	log.Printf("Firebase for DApps 🔥")

	resources := app.SetUp()
	if resources == nil {

		log.Printf("[❗️] Shutting down\n")
		os.Exit(1)

	}
}
