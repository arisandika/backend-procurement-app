#!/bin/bash

# =====================================================
# Procurement App (Backend)
# Go + Fiber + GORM
# =====================================================

# -----------------------------
# 1. Clone Repository
# -----------------------------
git clone https://github.com/arisandika/backend-procurement-app.git
cd backend-procurement-app


# -----------------------------
# 2. Install Dependencies
# -----------------------------
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv


# -----------------------------
# 3. Setup Environment File
# -----------------------------
# Rename .env.example to .env
cp .env.example .env


# Isi file .env dengan konfigurasi berikut
cat <<EOF > .env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=procurement_db

JWT_SECRET=supersecretkeyarisandika

WEBHOOK_URL=https://webhook.site/ad8a98bb-dbc4-4a58-8826-a14e2750571c
WEBHOOK_TIMEOUT=5
EOF

# NOTE:
# - Sesuaikan DB_USER, DB_PASS, dan DB_NAME dengan MySQL kamu
# - Pastikan database sudah dibuat di MySQL


# -----------------------------
# 4. Run Application
# -----------------------------
go run cmd/server/main.go

# Saat server pertama kali dijalankan:
# - Database akan terkoneksi otomatis
# - Auto migrate tabel dari models
# - Seed data awal (admin, supplier, item)
# - Server berjalan di http://localhost:3000


# -----------------------------
# 5. Test Login (Postman / Curl)
# -----------------------------
# Endpoint: POST /login
# Body:
# {
#   "username": "admin",
#   "password": "password"
# }

# Contoh curl:
# curl -X POST http://localhost:3000/login \
#   -H "Content-Type: application/json" \
#   -d '{"username":"admin","password":"password"}'


# -----------------------------
# 6. Reset Database (Optional)
# -----------------------------
# - Drop database atau semua tabel
# - Jalankan ulang:
# go run cmd/server/main.go
