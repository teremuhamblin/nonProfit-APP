###### STRUCTURE.md >> markdown
# nonProfit-APP
---

>> Une base propre, scalable, orientée ***API REST sécurisée avec JWT + SQLite + CI/CD + Swagger***.

---

📁 Arborescence du projet

```text
nonProfit-APP/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── http/
│   │   ├── router.go
│   │   ├── middleware.go
│   │   └── handlers/
│   │       ├── donor_handler.go
│   │       ├── donation_handler.go
│   │       └── auth_handler.go
│   │
│   ├── domain/
│   │   ├── donor.go
│   │   ├── donation.go
│   │   └── user.go
│   │
│   ├── storage/
│   │   ├── sqlite.go
│   │   └── migrations.go
│   │
│   ├── security/
│   │   └── jwt.go
│   │
│   └── docs/
│       └── swagger.go
│
├── config.yaml
├── go.mod
├── MAKEFILE
├── DOCKERFILE
├── README.md
└── .github/
    └── workflows/
        └── ci.yml
```

---

### 🧱 Rôles
- Principaux modules
```text
- cmd/api/main.go  
  Point d’entrée de l’API : charge la config, initialise la DB, configure le router, lance le serveur.

- internal/config/  
  Gestion de la configuration (Viper) : port HTTP, chemin SQLite, secrets JWT, mode debug.

- internal/storage/  
  Accès aux données (SQLite via GORM ou sqlx) + migrations auto (création des tables Donor, Donation, User).

- internal/domain/  
  Modèles métier (Donor, Donation, User) + éventuellement interfaces de repository.

- internal/security/jwt.go  
  Génération/validation des tokens JWT, gestion des claims (user ID, roles).

- internal/http/router.go  
  Définition des routes, montage des handlers, injection des middlewares (auth, logging, CORS).

- internal/http/middleware.go  
  Auth JWT, CORS, logging, rate limiting.

- internal/http/handlers/  
  Handlers REST :
  - auth_handler : POST /auth/login, POST /auth/register
  - donor_handler : GET/POST /donors
  - donation_handler : GET/POST /donations

- internal/docs/swagger.go  
  Déclaration des routes pour Swagger (/swagger/index.html) + génération via swag.

- .github/workflows/ci.yml  
  Build, tests, lint, éventuellement build Docker.
```

---

### 🔐 Flux typique
- JWT + SQLite
```text
1. Register : POST /auth/register → création user en DB → retour 201.
2. Login : POST /auth/login → vérif credentials → génération JWT → retour token.
3. Routes protégées : Authorization: Bearer <token>  
   - GET /donors
   - POST /donations
4. Middleware auth vérifie le token, injecte userID dans le contexte.
```

---

### 🧪 CI/CD minimal
- GitHub Actions
```text
- ci.yml :
  - go test ./...
  - go vet ./...
  - golangci-lint run
  - build binaire nonprofit-api
  - optionnel : build Docker
```

---
