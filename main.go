/*

	Main file. This is the application launch point
	The first is the installing configs
	Next comes the initialization of the database
	And the web server starts

*/

package main

import (
	"fmt"
	"log"
	"main.go/config"
	"main.go/routes"
	"main.go/storage"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	cfg := config.Load()

	db, err := storage.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.Setup(db)

	log.Printf("Server running on %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
