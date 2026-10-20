package main

import (
	"log"
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/config"
	"github.com/teremuhamblin/nonProfit-APP/internal/http"
	"github.com/teremuhamblin/nonProfit-APP/internal/storage"
)

func main() {
	// Charger la configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erreur configuration: %v", err)
	}

	// Initialiser la base SQLite
	db, err := storage.InitSQLite(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Erreur SQLite: %v", err)
	}

	// Router HTTP
	router := http.NewRouter(db)

	log.Printf("NonProfitApp Go running on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatal(err)
	}
}
