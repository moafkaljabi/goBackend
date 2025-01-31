package main

import (
	"goBackend/internal/database"
	"goBackend/internal/server"
	"log"
)

func main() {

	store, err := database.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.NewAPIServer(":3000", store)
	server.Run()
}
