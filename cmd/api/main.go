package main

import (
	"fmt"

	"github.com/teremuhamblin/nonProfit-APP/internal/config"
	"github.com/teremuhamblin/nonProfit-APP/internal/http"
	"github.com/teremuhamblin/nonProfit-APP/internal/storage"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("config error", zap.Error(err))
	}

	db, err := storage.InitSQLite(cfg.Database.Path)
	if err != nil {
		logger.Fatal("db error", zap.Error(err))
	}

	api := http.NewRouter(db, cfg, logger)
	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	logger.Info("NonProfit-APP running", zap.String("addr", addr))
	if err := api.Listen(addr); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}
