# 📝 Blog API

Blog API adalah sebuah RESTful API sederhana yang dibuat menggunakan bahasa pemrograman Go (Golang). Proyek ini bertujuan untuk menyediakan antarmuka backend untuk membuat, membaca, memperbarui, dan menghapus artikel blog (CRUD).

## 📁 Struktur Folder

├── main.go
├── internal/
│ ├── handlers/ # Berisi logic handler (controller)
│ ├── models/ # Berisi struct dan model data
│ └── repository/ # Berisi akses ke database (query)
├── configs/ # Konfigurasi aplikasi (YAML, env, dll)
└── go.mod / go.sum


## 🚀 Fitur

- ✍️ Create Post
- 📖 Read All Posts
- 🔍 Read Post by ID
- ✏️ Update Post
- ❌ Delete Post
- 📦 Modular architecture (handler, service, repository)

## 🔧 Teknologi

- ⚙️ Golang
- 🐘 MySQL
- 🧱 REST API
- 📦 Go Modules (gin)

## ⚙️ Instalasi

```bash
# 1. Clone repository
git clone https://github.com/DediMurphy/blog-api.git
cd blog-api

# 2. Inisialisasi module & install dependencies
go mod tidy

# 3. Atur konfigurasi
# Edit file configs/config.yml sesuai dengan konfigurasi database lokal kamu
# Contoh:
# 
# app:
#   port: 8080
# database:
#   host: localhost
#   port: 3306
#   user: root
#   password: yourpassword
#   name: blog_db

# 4. Jalankan aplikasi
go run main.go

👤 Author
Dedi Murphy – @DediMurphy
