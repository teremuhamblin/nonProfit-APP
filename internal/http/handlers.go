package http

import (
	"encoding/json"
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/domain"
	"github.com/teremuhamblin/nonProfit-APP/internal/storage"
)

func donorsHandler(store *storage.MemoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(store.Donors)

		case http.MethodPost:
			var d domain.Donor
			if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
				http.Error(w, "invalid body", http.StatusBadRequest)
				return
			}
			d.ID = int64(len(store.Donors) + 1)
			store.Donors = append(store.Donors, d)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(d)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
