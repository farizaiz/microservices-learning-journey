# 🚀 Golang & Microservices Learning Journey

Selamat datang di repositori perjalanan saya dalam belajar **Backend Engineering**, dengan fokus utama pada **Go (Golang)** dan arsitektur **Microservices**. 

Repositori ini berfungsi sebagai arsip pembelajaran terstruktur sekaligus portofolio profesional yang mendokumentasikan progres saya—dari pemahaman fundamental Go hingga perancangan dan implementasi aplikasi *distributed system* skala *enterprise*.

---

## 📂 Struktur Repositori

Pembelajaran dalam repositori ini dibagi menjadi 4 tahapan utama yang disusun secara inkremental:

### [📁 01-Golang-Fundamental](./01-Golang-Fundamental)
Tahap penguasaan dasar-dasar bahasa pemrograman Go.
* **Topik Utama:** Sintaks dasar, *Tipe Data*, *Structs*, *Interfaces*, dan implementasi *Concurrency* menggunakan *Goroutines* & *Channels*.

### [📁 02-REST-API](./02-REST-API)
Tahap transisi ke pengembangan web (*Web Development*).
* **Topik Utama:** Membangun RESTful API standar menggunakan *framework* **Gin-Gonic**. Fokus pada sistem *Routing*, pembuatan *Middleware* sederhana, penanganan *Request/Response* JSON, dan operasi CRUD dasar.

### [📁 03-Microservices-Fundamental](./03-Microservices-Fundamental)
Tahap pengenalan arsitektur terdistribusi.
* **Topik Utama:** Memecah aplikasi *Monolith* menjadi servis-servis kecil yang independen. Mempelajari konsep *Service-to-Service Communication*, pemisahan *Database* per *Service*, dan manajemen *Port*.

### [📁 04-Project-SMIPI](./04-Project-SMIPI) 🌟 *Capstone Project*
Tahap implementasi akhir (Proyek Utama). Membangun **SMIPI (Sistem Manajemen Informasi Pelaporan Investigasi)**, sebuah portal web operasional terpadu yang menggabungkan seluruh konsep dari tahap 01 hingga 03.

---

## 🛡️ Tahap 04: Capstone Project - SMIPI

**SMIPI (Sistem Manajemen Informasi Pelaporan Investigasi)** adalah portal aplikasi web skala *enterprise* yang dirancang untuk mengelola, melacak, dan memantau pelaporan kasus investigasi secara *real-time*. Sistem ini dibangun menggunakan arsitektur **Microservices** untuk memastikan skalabilitas, keamanan, dan pemisahan domain kerja yang jelas.

### 🚀 Arsitektur & Teknologi

Aplikasi ini dipisahkan menjadi beberapa layanan independen (*services*) yang berkomunikasi melalui REST API, menggunakan tumpukan teknologi modern:

**Backend (Microservices):**
* **Bahasa Pemrograman:** Golang (Go)
* **Web Framework:** Gin-Gonic
* **Database:** PostgreSQL (menggunakan GORM)
* **Keamanan:** Stateless JWT (JSON Web Tokens) & Bcrypt (Password Hashing)

**Frontend:**
* **Library:** React.js (di-build dengan Vite)
* **Routing:** React Router DOM (Nested Layouting)
* **HTTP Client:** Axios

### ✨ Fitur Unggulan (*Key Features*)

1. **Microservices Domain Separation:**
   * **User Service (Port 8081):** Mengelola otentikasi, otorisasi, dan manajemen identitas penyidik.
   * **Case Service (Port 8080):** Mengelola logika operasional pelaporan kasus, status, dan riwayat investigasi.

2. **Role-Based Access Control (RBAC):**
   * Diimplementasikan di level *Middleware* Golang dan *Frontend Routing*.
   * **Admin:** Memiliki akses penuh, termasuk registrasi dan manajemen pengguna (*User Management*).
   * **Investigator:** Dapat membuat, mengubah status, dan memantau kasus.
   * **Viewer:** Akses *Read-Only* untuk pemantauan dasbor.

3. **Keamanan & Standar Audit:**
   * **Stateless Auth:** Tidak menyimpan sesi di memori server, sepenuhnya mengandalkan JWT.
   * **Soft Deletes:** Penghapusan data tidak menghilangkan rekam jejak di *database* (krusial untuk data investigasi/hukum).
   * **Audit Trail Logic:** Setiap tindakan di-*stempel* dengan identitas pengguna yang melakukan perubahan.

4. **Dynamic Operational Dashboard:**
   * Tampilan *Summary Card* interaktif yang terhubung langsung dengan filter tabel.
   * Fitur pencarian, filter berdasarkan prioritas, dan filter berdasarkan status kasus.

---

## 📸 Antarmuka Pengguna (Screenshots)

| Halaman Login |
| :---: |
| ![Login Page](./frontend-app/src/assets/screenshot-login.png) |

| Papan Kendali Operasional | Manajemen Pengguna (Admin Only) |
| :---: | :---: |
| ![Dashboard](./frontend-app/src/assets/screenshot-dashboard.png) | ![User Management](./frontend-app/src/assets/screenshot-users.png) |

---

## 🛠️ Panduan Instalasi & Menjalankan Lokal

Pastikan Anda telah menginstal **Go**, **Node.js**, dan **PostgreSQL** di mesin Anda.

### 1. Setup Database
Buat dua database di PostgreSQL Anda:
* `smipi_users_db` (Untuk User Service)
* `smipi_cases_db` (Untuk Case Service)

### 2. Menjalankan User Service
Buka terminal dan arahkan ke folder `user-service`:
```bash
cd user-service
go mod tidy
go run main.go
```
Service ini akan berjalan di http://localhost:8081

### 3. Menjalankan Case Service
Buka terminal baru dan arahkan ke folder case-service:
```bash
cd case-service
go mod tidy
go run main.go
```
Service ini akan berjalan di http://localhost:8080

### 4. Menjalankan Frontend (React UI)
Buka terminal baru dan arahkan ke folder frontend-app:
```bash
cd frontend-app
npm install
npm run dev
```
Aplikasi web dapat diakses melalui browser di http://localhost:5173
