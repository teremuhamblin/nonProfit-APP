package handlers

import (
	"encoding/json"
	"net/http"

	hctx "github.com/teremuhamblin/nonProfit-APP/internal/http"
)

// Standard JSON response
type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Write JSON success
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(JSONResponse{
		Success: true,
		Data:    data,
	})
}

// Write JSON error
func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(JSONResponse{
		Success: false,
		Error:   msg,
	})
}

// Parse JSON body into struct
func ParseJSON(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// Extract user ID from context (JWT middleware)
func GetUserID(r *http.Request) (uint, bool) {
	return hctx.GetUserID(r.Context())
}
