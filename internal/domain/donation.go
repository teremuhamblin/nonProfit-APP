package domain

import "time"

type Donation struct {
	ID        uint      `gorm:"primaryKey"`
	DonorID   uint      `gorm:"index"`
	Amount    float64   `gorm:"not null"`
	CreatedAt time.Time
}
