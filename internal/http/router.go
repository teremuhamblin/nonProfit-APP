package http

import (
	"net/http"

	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/donors", DonorsHandler(db))
	mux.HandleFunc("/donations", DonationsHandler(db))
	mux.HandleFunc("/campaigns", CampaignsHandler(db))

	return mux
}
