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
	// KONFIGURASI CORS
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

	// ---------------------------------------------------------
	// RUTE PUBLIK (Siapa saja boleh masuk)
	// ---------------------------------------------------------
	// Rute register dihapus dari sini agar publik tidak bisa mendaftar sendiri
	r.POST("/login", controllers.LoginUser)

	// ==========================================
	// RUTE RAHASIA (Dijaga ketat oleh Middleware)
	// ==========================================
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		// 1. Rute Profil Umum
		protected.GET("/profil", controllers.GetProfil)

		// ---------------------------------------------------------
		// IMPLEMENTASI RBAC (Role-Based Access Control)
		// ---------------------------------------------------------

		// 2. Rute Khusus Admin (Sangat Ketat)

		// 🔒 FITUR BARU: Hanya Admin yang bisa mendaftarkan penyidik/user baru
		protected.POST("/register", middlewares.RequireRoles("Admin"), controllers.RegisterUser)

		// Melihat daftar seluruh pengguna di sistem SMIPI
		protected.GET("/users", middlewares.RequireRoles("Admin"), controllers.GetAllUsers)

		// 3. Rute Khusus Admin & Supervisor
		protected.PUT("/users/:id/role", middlewares.RequireRoles("Admin", "Supervisor"), controllers.UpdateUserRole)

		// 4. Rute Hapus User (Hanya Admin)
		protected.DELETE("/users/:id", middlewares.RequireRoles("Admin"), controllers.NonaktifkanUser)
	}

	fmt.Println("🚀 User Service menyala di PORT 8081...")
	r.Run(":8081")
}
