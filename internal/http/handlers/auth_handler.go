package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/teremuhamblin/nonProfit-APP/internal/config"
	"github.com/teremuhamblin/nonProfit-APP/internal/domain"
	"github.com/teremuhamblin/nonProfit-APP/internal/security"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *gorm.DB, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := ParseJSON(r, &req); err != nil {
			WriteError(w, 400, "invalid json")
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		user := domain.User{
			Email:    req.Email,
			Password: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			WriteError(w, 400, "email already exists")
			return
		}

		WriteJSON(w, 201, "user created")
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(db *gorm.DB, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := ParseJSON(r, &req); err != nil {
			WriteError(w, 400, "invalid json")
			return
		}

		var user domain.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			WriteError(w, 401, "invalid credentials")
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
			WriteError(w, 401, "invalid credentials")
			return
		}

		token, _ := security.GenerateToken(cfg.Security.JWTSecret, user.ID)
		WriteJSON(w, 200, map[string]string{"token": token})
	}
}
