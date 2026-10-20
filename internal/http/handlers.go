package http

import (
	"encoding/json"
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/domain"
	"gorm.io/gorm"
)

func DonorsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			var donors []domain.Donor
			db.Find(&donors)
			json.NewEncoder(w).Encode(donors)

		case http.MethodPost:
			var d domain.Donor
			json.NewDecoder(r.Body).Decode(&d)
			db.Create(&d)
			json.NewEncoder(w).Encode(d)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func DonationsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			var donations []domain.Donation
			db.Find(&donations)
			json.NewEncoder(w).Encode(donations)

		case http.MethodPost:
			var d domain.Donation
			json.NewDecoder(r.Body).Decode(&d)
			db.Create(&d)
			json.NewEncoder(w).Encode(d)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func CampaignsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			var campaigns []domain.Campaign
			db.Find(&campaigns)
			json.NewEncoder(w).Encode(campaigns)

		case http.MethodPost:
			var c domain.Campaign
			json.NewDecoder(r.Body).Decode(&c)
			db.Create(&c)
			json.NewEncoder(w).Encode(c)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
