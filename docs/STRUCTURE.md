###### STRUCTURE.md >> markdown 

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

