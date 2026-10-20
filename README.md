###### README.md >> markdown 
# nonProfit-APP

### Structure proposée du projet Go

---

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
