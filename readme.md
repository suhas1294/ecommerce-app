1. Install golang
2. Export its path in zshrc / bashrc
2. Check `go -version`
4. Install sqlite3 : brew install sqlite, confirm after installing `sqlite3 --version`
5. Testing if sqlite 3 is installed properly or not
   1. one more way to test if sqlite is installed correctly by creating a db  : `sqlite3 ecommerce.db` : This opens a DB file ecommerce.db in the current folder. Exit with `.exit`
6. mkdir ecommerce-app
7. cd ecommerce-app
8. go mod init ecommerce-app
9. Installing Add GORM + SQLite driver
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```
GORM: ORM itself\
SQLite Driver: for connecting to SQLite from Go

#### Redefining project structure : 

```
ecommerce-app/
├── cmd/                # Entrypoints (main.go)
├── internal/
│   ├── db/             # DB connection logic
│   ├── models/         # GORM models
│   ├── handlers/       # HTTP handlers
│   ├── services/       # Business logic
│   ├── middleware/     # Auth, logging
│   ├── auth/           # JWT + OAuth
├── config/             # App config
├── go.mod
├── go.sum
```