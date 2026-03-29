package main

import (
	"fmt"

	"case-service/config"
	"case-service/controllers"
	"case-service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Jalankan koneksi database
	config.ConnectDatabase()

	// 2. Siapkan Gin Router
	r := gin.Default()

	// ==========================================
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// ==========================================

	// Rute percobaan
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "SMIPI Case Service Berjalan Mulus!"})
	})

	// ==========================================
	// RUTE RAHASIA (Dijaga Satpam)
	// ==========================================
	protected := r.Group("/api") // Praktik terbaik: gunakan prefix /api

	// SATPAM LAPIS 1: Cek apakah bawa Token JWT yang valid
	protected.Use(middlewares.AuthMiddleware())
	{
		// 1. CREATE Kasus (Hanya Investigator dan Admin yang boleh buat laporan)
		protected.POST("/kasus", middlewares.RequireRoles("Investigator", "Admin"), controllers.CreateKasus)

		// 2. READ Kasus (Semua role boleh melihat Papan Kendali)
		protected.GET("/kasus", middlewares.RequireRoles("Investigator", "Supervisor", "Admin", "Viewer"), controllers.GetAllKasus)

		// 3. UPDATE Seluruh Kasus (Edit detail laporan)
		protected.PUT("/kasus/:id", middlewares.RequireRoles("Investigator", "Supervisor", "Admin"), controllers.UpdateKasus)

		// ---------------------------------------------------------
		// RUTE BARU: Khusus untuk Modal "Update Status Kasus" di React
		// ---------------------------------------------------------
		protected.PUT("/kasus/:id/status", middlewares.RequireRoles("Investigator", "Supervisor", "Admin"), controllers.UpdateStatusKasus)

		// 4. DELETE Kasus (Soft Delete, sangat ketat, hanya atasan yang boleh)
		protected.DELETE("/kasus/:id", middlewares.RequireRoles("Supervisor", "Admin"), controllers.DeleteKasus)
	}
	// 4. Jalankan server
	fmt.Println("🚀 Case Service menyala di PORT 8080...")
	r.Run(":8080")
}
