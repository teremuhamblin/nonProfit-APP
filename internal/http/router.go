package http

import (
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/storage"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	store := storage.NewMemoryStore()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/donors", donorsHandler(store))

	return mux
}
