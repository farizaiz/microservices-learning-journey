// File: config/database.go
package config

import (
	"fmt"
	"log"
	"os"

	"eval-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Merakit DSN (Alamat Database) dari kepingan-kepingan file .env
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal terhubung ke Database: ", err)
	}
	fmt.Println("✅ BINGO! Berhasil masuk ke Database PostgreSQL!")

	DB.AutoMigrate(&models.Dokumen{})
	fmt.Println("✅ Tabel 'dokumens' berhasil disinkronisasi ke Database!")
}
