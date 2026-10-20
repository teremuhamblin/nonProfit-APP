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
├── Makefile
├── Dockerfile
├── README.md
└── .github/
    └── workflows/
        └── ci.yml
```

---

