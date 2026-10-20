package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/teremuhamblin/nonProfit-APP/internal/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID uint `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    token, err := security.GenerateToken(req.UserID)
    if err != nil {
        http.Error(w, "failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}
