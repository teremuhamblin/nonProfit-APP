package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"

	"github.com/teremuhamblin/nonProfit-APP/internal/config"
	"github.com/teremuhamblin/nonProfit-APP/internal/http/handlers"
	"go.uber.org/zap"
)

type APIServer struct {
	router *chi.Mux
}

func NewRouter(db *gorm.DB, cfg *config.Config, logger *zap.Logger) *APIServer {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	}))

	for _, mw := range CommonMiddlewares(logger, cfg) {
		r.Use(mw)
	}

	r.Post("/auth/register", handlers.Register(db, cfg))
	r.Post("/auth/login", handlers.Login(db, cfg))

	r.Group(func(pr chi.Router) {
		pr.Use(JWTAuth(cfg.Security.JWTSecret))

		pr.Get("/donors", handlers.GetDonors(db))
		pr.Post("/donors", handlers.CreateDonor(db))

		pr.Get("/donations", handlers.GetDonations(db))
		pr.Post("/donations", handlers.CreateDonation(db))
	})

	return &APIServer{router: r}
}

func (s *APIServer) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
