package main

import (
	"fmt"
	"log"

	"github.com/teremuhamblin/nonProfit-APP/internal/config"
	"github.com/teremuhamblin/nonProfit-APP/internal/http"
	"github.com/teremuhamblin/nonProfit-APP/internal/storage"
	"go.uber.org/zap"
)

func main() {
	// Logger professionnel
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Starting NonProfit-APP API...")

	// Charger la configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Initialiser la base SQLite
	db, err := storage.InitSQLite(cfg.Database.Path)
	if err != nil {
		logger.Fatal("Failed to initialize SQLite", zap.Error(err))
	}

	// Initialiser le router HTTP (Chi)
	api := http.NewRouter(db, cfg, logger)

	// Démarrer le serveur
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("NonProfit-APP running", zap.String("address", addr))

	if err := api.Listen(addr); err != nil {
		logger.Fatal("Server stopped with error", zap.Error(err))
	}
}
