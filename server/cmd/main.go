package main

import (
	"log"

	"github.com/alonsofritz/simple-chatapp-go/server/db"
	"github.com/alonsofritz/simple-chatapp-go/server/internal/user"
	"github.com/alonsofritz/simple-chatapp-go/server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Could not initialize database connection: %s", err)
	}

	userRep := user.NewUserRepository(dbConn.GetDB())
	userSvc := user.NewUserService(userRep)
	userHandler := user.NewUserHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
