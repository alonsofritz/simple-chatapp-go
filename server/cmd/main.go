package main

import (
	"log"

	"github.com/alonsofritz/simple-chatapp-go/server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Could not initialize database connection: %s", err)
	}
}
