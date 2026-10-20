package handlers

import (
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/domain"
	"gorm.io/gorm"
)

func GetDonors(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var donors []domain.Donor
		db.Find(&donors)
		WriteJSON(w, 200, donors)
	}
}

type DonorRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateDonor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DonorRequest
		if err := ParseJSON(r, &req); err != nil {
			WriteError(w, 400, "invalid json")
			return
		}

		donor := domain.Donor{Name: req.Name, Email: req.Email}
		db.Create(&donor)

		WriteJSON(w, 201, donor)
	}
}
