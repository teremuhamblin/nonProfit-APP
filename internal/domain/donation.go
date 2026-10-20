package domain

type Donation struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	DonorID uint    `json:"donor_id"`
	Amount  float64 `json:"amount"`
}
