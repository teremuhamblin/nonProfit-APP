package handlers

import (
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/domain"
	"gorm.io/gorm"
)

func GetDonations(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var donations []domain.Donation
		db.Find(&donations)
		WriteJSON(w, 200, donations)
	}
}

type DonationRequest struct {
	DonorID uint    `json:"donor_id"`
	Amount  float64 `json:"amount"`
}

func CreateDonation(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DonationRequest
		if err := ParseJSON(r, &req); err != nil {
			WriteError(w, 400, "invalid json")
			return
		}

		donation := domain.Donation{
			DonorID: req.DonorID,
			Amount:  req.Amount,
		}

		db.Create(&donation)
		WriteJSON(w, 201, donation)
	}
}
