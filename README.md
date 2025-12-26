# Procurement App

Aplikasi manajemen procurement berbasis **Go + Fiber + GORM** dengan fitur CRUD master data, dan purchasing.

---

## Dependencies Utama

Jalankan perintah berikut sebelum mulai:

go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv

---

## Setup (All-in-One)

1. **Clone repository**

git clone https://github.com/arisandika/backend-procurement-app.git
cd backend-procurement-app

2. **Rename file `.env.example`** menjadi `.env` di root folder:

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=procurement_db
JWT_SECRET=supersecretkeyarisandika
WEBHOOK_URL=https://webhook.site/ad8a98bb-dbc4-4a58-8826-a14e2750571c
WEBHOOK_TIMEOUT=5

> Sesuaikan `DB_USER`, `DB_PASS`, `DB_NAME` dengan konfigurasi MySQL kamu.

3. **Jalankan aplikasi**

go run cmd/server/main.go

- Database akan otomatis **terkoneksi**, **migrasi tabel** berdasarkan package `models`, dan **seed data awal** (admin, supplier, item).  
- Server Fiber berjalan di **http://localhost:3000**.  

4. **Login pertama**

- Username: `admin`  
- Password: `admin123`  

> Setelah login, semua fitur CRUD dan purchasing sudah bisa digunakan.

---

## Fitur

- CRUD Master Data:
  - Supplier
  - Item
  - User
- Purchasing:
  - Create / Update / Delete
  - Keranjang (cart) dengan localStorage
  - Multi-tab sync
- Dashboard / Summary:
  - Total per supplier, item, dan transaksi
- Modal / SweetAlert untuk interaksi user-friendly
- Auto-migrate & seed data saat `go run main.go`

---

## Notes

- **Migrasi**: otomatis dari `main.go` via `db.AutoMigrate(...)`  
- **Seed data**: hanya dijalankan jika tabel kosong, sehingga aman dijalankan berulang.  
- **Reset database**: hapus database atau drop tabel, kemudian `go run cmd/server/main.go` lagi.
