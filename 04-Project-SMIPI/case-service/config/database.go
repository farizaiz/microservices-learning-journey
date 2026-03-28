package config

import (
	"fmt"
	"log"
	"os" // Tambahan untuk membaca environment variable

	"case-service/models"

	"github.com/joho/godotenv" // Tambahan library dotenv
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// 1. Muat (Load) file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Peringatan: File .env tidak ditemukan, menggunakan variabel sistem.")
	}

	// 2. Ambil rahasia dari brankas .env
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// 3. Rakit DSN (Data Source Name) secara dinamis
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	// 4. Buka Koneksi
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal terhubung ke database:", err)
	}

	fmt.Println("✅ Berhasil terhubung ke PostgreSQL via .env!")

	// 5. Auto Migrate
	// 🚨 TAMBAHKAN &models.AuditLog{} DI SINI 🚨
	err = database.AutoMigrate(&models.Kasus{}, &models.AuditLog{})
	if err != nil {
		log.Fatal("❌ Gagal melakukan migrasi tabel:", err)
	}

	DB = database
}
