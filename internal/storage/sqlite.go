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

	// Auto migration
	db.AutoMigrate(
		&domain.Donor{},
		&domain.Donation{},
		&domain.Campaign{},
	)

	return db, nil
}
