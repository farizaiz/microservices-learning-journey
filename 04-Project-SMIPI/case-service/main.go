package main

import (
	"fmt"

	"case-service/config"
	"case-service/controllers" // Import folder controllers Anda

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Jalankan koneksi database
	config.ConnectDatabase()

	// 2. Siapkan Gin Router
	r := gin.Default()

	// Rute percobaan
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "SMIPI Case Service Berjalan Mulus!"})
	})

	// ==========================================
	// 3. DAFTARKAN RUTE API KASUS DI SINI
	// ==========================================
	r.POST("/kasus", controllers.CreateKasus)
	r.GET("/kasus", controllers.GetAllKasus)        // <-- Rute untuk mengambil semua data
	r.GET("/kasus/:id", controllers.GetKasusByID)   // <-- Rute untuk mengambil data by ID
	r.PUT("/kasus/:id", controllers.UpdateKasus)    // <-- Rute untuk Update
	r.DELETE("/kasus/:id", controllers.DeleteKasus) // <-- Rute untuk Delete

	// 4. Jalankan server
	fmt.Println("🚀 Case Service menyala di PORT 8080...")
	r.Run(":8080")
}
