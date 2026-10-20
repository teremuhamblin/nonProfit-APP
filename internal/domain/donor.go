package domain

import "time"

type Donor struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
