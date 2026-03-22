// File: main.go
package main

import (
	"fmt"
	"log"

	"eval-service/config"
	"eval-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. BACA FILE .env TERLEBIH DAHULU!
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error memuat file .env")
	}

	fmt.Println("🚀 Memulai Service Evaluasi SPEED KLHK...")

	// 2. Nyalakan Mesin Database
	config.ConnectDatabase()

	// 3. Siapkan Pelayan
	r := gin.Default()

	// 4. Serahkan Buku Menu
	routes.SetupRoutes(r)

	fmt.Println("🟢 Server Microservices berjalan di Port 8080")
	r.Run(":8080")
}
