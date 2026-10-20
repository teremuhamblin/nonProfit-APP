###### nonProfit-APP/README.md >> markdown 
>>API moderne : Go 1.23, Chi, JWT, SQLite, Viper, Zap, CI/CD, Swagger.

# 🟦 nonProfit-APP
- API REST sécurisée en Go
>Une API moderne, rapide et sécurisée écrite en Golang, conçue pour gérer :
```text
- les utilisateurs (authentification JWT)
- les donateurs
- les dons
- les campagnes (extension possible)
```
- Cette base est pensée pour être scalable, propre, modulaire, et prête pour la production.

---

### 🚀 Fonctionnalités principales
#### 🔐 Authentification JWT
```text
- POST /auth/register — création d’un utilisateur  
- POST /auth/login — génération d’un token JWT  
- Middleware de protection des routes  
- Extraction automatique du user_id depuis le token
```
#### 👥 Donateurs
```text
- GET /donors — liste des donateurs  
- POST /donors — création d’un donateur  
```
#### 💸 Dons
```text
- GET /donations — liste des dons  
- POST /donations — création d’un don  
```
#### 🗄️ Base de données SQLite
```text
- Auto-migrations via GORM  
- Modèles : User, Donor, Donation
```
#### ⚙️ Configuration (Viper)
```text
- Port serveur  
- Chemin SQLite  
- Secret JWT  
- Mode debug  
```
#### 📦 CI/CD GitHub Actions
```text
- Build  
- Tests  
- Lint  
- Go 1.23  
```
#### 📘 Swagger (prêt à être activé)

#### 📁 Structure du projet
```markdown
nonProfit-APP/docs/STRUCTURE.md
```

⚙️ Installation & Lancement
1️⃣ Cloner le projet
```bash
git clone https://github.com/teremuhamblin/nonProfit-APP
cd nonProfit-APP
```

2️⃣ Installer les dépendances
```bash
go mod tidy
```

3️⃣ Lancer l’API
```bash
go run ./cmd/api
```

4️⃣ Build
```bash
make build
```

5️⃣ Docker
```bash
docker build -t nonprofit-app .
docker run -p 8080:8080 nonprofit-app
```

### 🔧 Configuration (config.yaml)
```yaml
server:
  port: 8080

database:
  path: "nonProfit-APP.db"

security:
  jwtsecret: "CHANGEMETOALONGRANDOM_SECRET"

app:
  debug: true
```

##$ 🔐 Authentification JWT
>Register
```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secret"
}
```

>Login
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secret"
}
```

>Réponse :
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR..."
  }
}
```

>Utilisation du token
```text
Authorization: Bearer <token>
```

### 📘 Endpoints REST
#### 👥 Donors

| Méthode | Endpoint     | Description |
|---------|--------------|-------------|
| GET     | /donors      | Liste des donateurs |
| POST    | /donors      | Créer un donateur |

#### 💸 Donations

| Méthode | Endpoint     | Description |
|---------|--------------|-------------|
| GET     | /donations   | Liste des dons |
| POST    | /donations   | Créer un don |

### 🧪 CI/CD GitHub Actions
>>Le workflow ci.yml :
```text
- installe Go 1.23  
- build le projet  
- exécute les tests  
- vérifie les dépendances  
```

### 🧱 Technologies utilisées
```text
- Go 1.23
- Chi Router
- GORM + SQLite
- Viper
- Zap Logger
- JWT (golang-jwt v5)
- GitHub Actions
- Docker
```

### 📌 Roadmap
```text
VOIR nonProfit-APP/docs/ROADMAP.md
```

### 🏁 Conclusion

- Ce projet est une base solide, moderne et sécurisée pour construire une API Go professionnelle.  
- Il est conçu pour être simple à comprendre, facile à étendre, et prêt pour la production.

---
