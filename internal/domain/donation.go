package domain

type Donation struct {
	ID      int64   `json:"id"`
	DonorID int64   `json:"donor_id"`
	Amount  float64 `json:"amount"`
}
