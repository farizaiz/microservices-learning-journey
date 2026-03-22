package main

import (
	"fmt"
	"user-service/config"
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// ==========================================
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Izinkan React masuk
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// ==========================================

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "SMIPI User Service Siap Melayani!"})
	})
	// RUTE PUBLIK (Siapa saja boleh masuk)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// ==========================================
	// RUTE RAHASIA (Dijaga ketat oleh Middleware)
	// ==========================================
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware()) // Pasang satpam di sini
	{
		// Artinya rute ini menjadi: GET http://localhost:8081/api/profil
		protected.GET("/profil", controllers.GetProfil)
	}

	fmt.Println("🚀 User Service menyala di PORT 8081...")
	r.Run(":8081")
}
