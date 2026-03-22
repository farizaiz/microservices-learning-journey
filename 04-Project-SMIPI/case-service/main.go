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
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/kasus", controllers.CreateKasus)
		protected.GET("/kasus", controllers.GetAllKasus)
		protected.PUT("/kasus/:id", controllers.UpdateKasus)
		protected.DELETE("/kasus/:id", controllers.DeleteKasus)
	}

	// 4. Jalankan server
	fmt.Println("🚀 Case Service menyala di PORT 8080...")
	r.Run(":8080")
}
