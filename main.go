/*

	Main file. This is the application launch point
	The first is the installing configs
	Next comes the initialization of the database
	And the web server starts

*/

package main

import (
	"context"
	"fmt"
	"log"
	"main.go/config"
	"main.go/routes"
	"main.go/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting server...")

	cfg := config.Load()

	db, err := storage.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.Setup(db)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	go func() {
		log.Printf("Server running on %s", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server stopped gracefully")
}
