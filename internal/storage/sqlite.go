package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/teremuhamblin/nonProfit-APP/internal/domain"
)

func InitSQLite(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Donor{}, &domain.Donation{}); err != nil {
		return nil, err
	}

	return db, nil
}
