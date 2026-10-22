###### ARCHITECTURE.md >> markdown 
# nonProfit-APP
###### docs/ARCHITECTURE.md

### 🧬 Architecture technique
#### 🏷️ Vue d’ensemble
- nonProfit-APP est une API REST sécurisée en Go 1.23, basée sur Chi, avec :
```text
- Authentification JWT
- SQLite via GORM
- Configuration dynamique via Viper
- Logging structuré via Zap
- CI/CD GitHub Actions
- Conteneurisation Docker
```

---

### 🧱 Architecture globale

```mermaid
flowchart TD
    Client[Client HTTP / Frontend] --> API[nonProfit-APP API (Go + Chi)]
    API --> MW[Middlewares (JWT, Logging, Recovery)]
    API --> H[Handlers REST]
    H --> S[Services métier]
    S --> DB[(SQLite via GORM)]
    S --> CFG[Config (Viper)]
    S --> LOG[Logger (Zap)]
```

---

### 🧩 Organisation du code

```mermaid
graph TD
    CMD[cmd/api] --> HTTP[internal/http]
    HTTP --> R[router.go]
    HTTP --> M[middleware.go]
    HTTP --> HND[handlers/*]

    CMD --> DOM[internal/domain]
    DOM --> U[user.go]
    DOM --> DNR[donor.go]
    DOM --> DON[donation.go]

    CMD --> STG[internal/storage]
    STG --> SQL[sqlite.go]

    CMD --> SEC[internal/security]
    SEC --> JWT[jwt.go]

    CMD --> CFG[internal/config]
    CFG --> CFGF[config.go]

    CMD --> DOC[internal/docs]
    DOC --> SWAG[swagger.go]
```

---

### 🛡️ Sécurité — Flux JWT

```mermaid
sequenceDiagram
    participant C as Client
    participant A as API nonProfit-APP
    participant D as DB (SQLite)
    participant J as JWT Engine

    C->>A: POST /auth/login (email, password)
    A->>D: Vérifier credentials
    D-->>A: OK / KO
    A->>J: Générer JWT (user_id, roles, exp)
    J-->>A: Token signé (HS256)
    A-->>C: 200 OK + token

    C->>A: GET /donations (Authorization: Bearer <token>)
    A->>J: Vérifier signature + expiration
    J-->>A: Token valide
    A->>D: Récupérer donations(user_id)
    D-->>A: Liste des dons
    A-->>C: 200 OK + JSON
```

---

### ⚙️ Pipeline CI/CD

```mermaid
flowchart LR
    Dev[Developer Push] --> GH[GitHub Repo]
    GH --> CI[Workflow ci.yml]
    CI --> B[go build ./...]
    CI --> T[go test ./...]
    CI --> L[golangci-lint run]
    CI --> SEC[go vet / security checks]
    CI --> REL[release.yml (tag + changelog)]
    REL --> IMG[Docker build & push]
```

---

### 🧱 Dataflow Donors / Donations

```mermaid
flowchart TD
    U[User (JWT)] --> API[nonProfit-APP API]
    API --> DNR[Donor Service]
    API --> DON[Donation Service]

    DNR --> DB[(SQLite)]
    DON --> DB

    DB --> RPT[Reporting / Stats (future)]
```

---

### 📊 Stack technique

| Composant       | Technologie      | Rôle                          |
|-----------------|------------------|-------------------------------|
| Langage         | Go 1.23          | Backend API                   |
| Router          | Chi              | Routing HTTP                  |
| ORM             | GORM             | Accès DB SQLite               |
| Base de données | SQLite           | Stockage persistant           |
| Config          | Viper            | Configuration dynamique       |
| Logging         | Zap              | Logs structurés               |
| Auth            | golang-jwt v5    | JWT HS256                     |
| CI/CD           | GitHub Actions   | Build, tests, lint, release   |
| Conteneur       | Docker           | Déploiement containerisé      |

---

### 📌 Endpoints principaux

| Domaine    | Méthode | Endpoint       | Auth JWT | Description                    |
|-----------|---------|----------------|----------|--------------------------------|
| Auth      | POST    | /auth/register | ❌       | Créer un utilisateur           |
| Auth      | POST    | /auth/login    | ❌       | Obtenir un token JWT           |
| Donors    | GET     | /donors        | ✅       | Liste des donateurs            |
| Donors    | POST    | /donors        | ✅       | Créer un donateur              |
| Donations | GET     | /donations     | ✅       | Liste des dons                 |
| Donations | POST    | /donations     | ✅       | Créer un don                   |

---

### Badges pour *README.md*

```md
<p align="center">
  <img src="https://img.shields.io/badge/Backend-Go%201.23-00ADD8?style=for-the-badge&logo=go"/>
  <img src="https://img.shields.io/badge/Framework-Chi-3D7EBB?style=for-the-badge&logo=go"/>
  <img src="https://img.shields.io/badge/Auth-JWT-8B0000?style=for-the-badge&logo=jsonwebtokens"/>
  <img src="https://img.shields.io/badge/DB-SQLite-003B57?style=for-the-badge&logo=sqlite"/>
  <img src="https://img.shields.io/badge/Config-Viper-4B8BBE?style=for-the-badge&logo=go"/>
  <img src="https://img.shields.io/badge/Logging-Zap-6A5ACD?style=for-the-badge&logo=go"/>
</p>

<p align="center">
  <img src="https://img.shields.io/github/actions/workflow/status/teremuhamblin/nonProfit-APP/ci.yml?style=for-the-badge&logo=github&label=CI%2FCD"/>
  <img src="https://img.shields.io/github/last-commit/teremuhamblin/nonProfit-APP?style=for-the-badge&color=8B0000&label=LAST%20COMMIT"/>
  <img src="https://img.shields.io/github/license/teremuhamblin/nonProfit-APP?style=for-the-badge&color=2F4F4F"/>
  <img src="https://img.shields.io/github/stars/teremuhamblin/nonProfit-APP?style=for-the-badge&color=FFD700&label=STARS"/>
</p>
```

---
