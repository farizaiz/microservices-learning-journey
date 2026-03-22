package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Endpoint API GET /users/validate/:id
	r.GET("/users/validate/:id", func(c *gin.Context) {
		userID := c.Param("id")

		// Simulasi database: Anggap saja hanya ID "123" dan "456" yang terdaftar
		isValid := false
		if userID == "123" || userID == "456" {
			isValid = true
		}

		// Mengirimkan respon JSON status validasi
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"valid":   isValid,
		})
	})

	fmt.Println("👤 User Service berjalan di PORT 8081...")
	// Perhatikan! Kita jalankan di port 8081 agar tidak bentrok dengan service lain
	r.Run(":8081")
}
