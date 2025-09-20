# Inventory App Backend

A lightweight backend for a home inventory application built in **Go** using **Chi** router and **SQLite**.  
Tracks items and their locations with JWT-based authentication.

---

## Features

- CRUD for items and locations
- SQLite database for lightweight storage
- JWT authentication for protected routes
- Middleware for logging, request ID, and panic recovery
- Easy-to-extend folder structure

---

## Tech Stack

- [Go](https://go.dev/)
- [Chi Router](https://github.com/go-chi/chi)
- [SQLite](https://www.sqlite.org/)
- [sqlx](https://jmoiron.github.io/sqlx/) for database handling
- [JWT](https://github.com/golang-jwt/jwt/v5) for authentication
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) for password hashing

---

## Folder Structure

inventory-app/
│── cmd/
│ └── server/
│ └── main.go # entrypoint
│
│── internal/
│ ├── api/ # route handlers
│ ├── auth/ # JWT auth + middleware
│ ├── db/ # database init + migrations
│ ├── models/ # structs for items, locations, users
│ └── services/ # business logic (optional)
│
│── go.mod
│── go.sum



## Setup & Run

### 1. Clone repo
```bash
git clone https://github.com/yourname/inventory-app.git
cd inventory-app
```

### 2. Initialize Go modules
```bash
go mod tidy
```

### 3. Install dependencies
```bash
go get github.com/go-chi/chi/v5
go get github.com/jmoiron/sqlx
go get github.com/mattn/go-sqlite3
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

### 4. Run the server
```bash
go run ./cmd/server
```