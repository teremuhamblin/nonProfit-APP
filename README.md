###### README.md >> markdown 
# nonProfit-APP
### Language : GOLang

---

- Une petite ***API en Go*** pour gérer :
```markdown
- les donateurs
- les dons
- les campagnes
```

>>Lancer le projet :
```bash
go run ./cmd/nonProfit-APP
```

>>Endpoints :
```markdown
- GET /health → status
- GET /donors → liste des donateurs
- POST /donors → créer un donateur
```

---

### Structure du projet Go
```text
nonprofitapp-go/
├── go.mod
├── go.sum
├── cmd/
│   └── nonprofitapp/
│       └── main.go
├── internal/
│   ├── http/
│   │   ├── router.go
│   │   └── handlers.go
│   ├── domain/
│   │   ├── donor.go
│   │   ├── donation.go
│   │   └── campaign.go
│   └── storage/
│       └── memory.go
└── README.md
```

---
