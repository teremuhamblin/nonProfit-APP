###### 🅰️ docs/ROADMAP.md
# 🧭 nonProfit-APP
- nonProfit-APP est une API moderne en Go 1.23, conçue pour être scalable, sécurisée et extensible.  
   - Cette roadmap présente les objectifs techniques, les évolutions prévues et les futures fonctionnalités.

---

### 🗺️ Vision globale

```mermaid
mindmap
  root((nonProfit-APP))
    Backend
      Go 1.23
      Chi Router
      GORM + SQLite
      JWT Auth
      Zap Logging
    DevOps
      GitHub Actions
      Linting
      Tests
      Releases SemVer
      Docker
    Features
      Donors
      Donations
      Campaigns
      Reporting
    Future
      Monitoring
      Metrics Prometheus
      OpenAPI 3.0
      Postgres migration
```

---

### 📌 Objectifs par version
#### ✅ v1.x
- Version actuelle

| Domaine | Objectif | Statut |
|--------|----------|--------|
| Auth | JWT complet | ✔️ |
| Donors | CRUD minimal | ✔️ |
| Donations | CRUD minimal | ✔️ |
| DB | SQLite + migrations | ✔️ |
| Logs | Zap | ✔️ |
| Config | Viper | ✔️ |
| CI/CD | Build + tests | 🔄 En cours |
| Docs | README + API | ✔️ |

---

#### 🚀 v2.0
>> Version majeure
```text
- 🎯 Objectifs
   - Ajout du module Campaigns
   - Ajout du module Reporting
   - Ajout d’un système de permissions (RBAC)
   - Migration optionnelle vers PostgreSQL
   - Ajout de Swagger complet
   - Ajout de tests unitaires + intégration
   - Ajout d’un dashboard admin (future extension)
```

>>📐 Diagramme des modules v2.0

```mermaid
flowchart TD
    API[nonProfit-APP API] --> AUTH[Auth]
    API --> DNR[Donors]
    API --> DON[Donations]
    API --> CMP[Campaigns]
    API --> RPT[Reporting]

    CMP --> DB[(SQLite/PostgreSQL)]
    RPT --> DB
```

---

#### 🧬 v3.0
>>Scalabilité & Observabilité
```text
- 🎯 Objectifs
   - Export métriques Prometheus
   - Intégration Grafana
   - Ajout d’un système de webhooks
   - Ajout d’un cache Redis
   - Ajout d’un rate limiter
   - Ajout d’un système d’alerting
```

---

#### 🛰️ v4.0
>>Enterprise Edition
```text
- 🎯 Objectifs
   - Multi-tenancy
   - Multi‑DB (MySQL, PostgreSQL, SQLite)
   - Système de plugins
   - API GraphQL
   - Intégration Keycloak / OAuth2
```

---

### 🧱 Roadmap détaillée (Mermaid Gantt)

```mermaid
gantt
    dateFormat  YYYY-MM-DD
    title Roadmap nonProfit-APP

    section Version 1.x
    Auth JWT                :done,    a1, 2024-01-01, 7d
    Donors CRUD             :done,    a2, 2024-01-08, 7d
    Donations CRUD          :done,    a3, 2024-01-15, 7d
    Docs API                :done,    a4, 2024-01-22, 5d

    section Version 2.0
    Campaigns module        :active,  b1, 2024-02-01, 14d
    Reporting module        :         b2, 2024-02-15, 14d
    RBAC permissions        :         b3, 2024-03-01, 14d
    Swagger full            :         b4, 2024-03-15, 10d

    section Version 3.0
    Prometheus metrics      :         c1, 2024-04-01, 14d
    Grafana dashboards      :         c2, 2024-04-15, 14d
    Redis cache             :         c3, 2024-05-01, 14d

    section Version 4.0
    Multi-tenancy           :         d1, 2024-06-01, 30d
    GraphQL API             :         d2, 2024-07-01, 20d
```

---

### 🏁 Conclusion

Cette roadmap guide l’évolution du projet vers une API robuste, scalable, sécurisée et adaptée aux besoins modernes.

---
