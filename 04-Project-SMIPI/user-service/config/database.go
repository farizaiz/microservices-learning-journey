package config

import (
	"fmt"
	"log"
	"os"

	"user-service/models" // Import struct User Anda

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ File .env tidak ditemukan")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal terhubung ke database:", err)
	}

	fmt.Println("✅ Berhasil terhubung ke PostgreSQL (smipi_user_db)!")

	// Auto-Migrate tabel User
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Gagal melakukan migrasi tabel User:", err)
	}

	fmt.Println("✅ Tabel User berhasil diciptakan!")

	DB = database
}
