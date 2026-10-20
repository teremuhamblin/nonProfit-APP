package http

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
    "gorm.io/gorm"

    "github.com/teremuhamblin/nonProfit-APP/internal/config"
    "github.com/teremuhamblin/nonProfit-APP/internal/http/handlers"
    "github.com/teremuhamblin/nonProfit-APP/internal/http/middleware"
    "go.uber.org/zap"
)

type APIServer struct {
    router *chi.Mux
}

func NewRouter(db *gorm.DB, cfg *config.Config, logger *zap.Logger) *APIServer {
    r := chi.NewRouter()

    // --- CORS ---
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Authorization", "Content-Type"},
    }))

    // --- Middlewares communs ---
    for _, mw := range CommonMiddlewares(logger, cfg) {
        r.Use(mw)
    }

    // --- Auth publique ---
    r.Post("/auth/register", handlers.Register(db, cfg))
    r.Post("/auth/login", handlers.Login(db, cfg))

    // --- Routes protégées ---
    r.Group(func(protected chi.Router) {
        protected.Use(middleware.JWTAuth)

        protected.Get("/donors", handlers.GetDonors(db))
        protected.Post("/donors", handlers.CreateDonor(db))

        protected.Get("/donations", handlers.GetDonations(db))
        protected.Post("/donations", handlers.CreateDonation(db))
    })

    return &APIServer{router: r}
}

func (s *APIServer) Listen(addr string) error {
    return http.ListenAndServe(addr, s.router)
}
