package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct untuk menerima data tugas (task) dari user
type TaskInput struct {
	UserID    string `json:"user_id"`
	NamaTugas string `json:"nama_tugas"`
}

// Struct untuk membaca balasan dari User Service
type UserValidationResponse struct {
	UserID string `json:"user_id"`
	Valid  bool   `json:"valid"`
}

func main() {
	r := gin.Default()

	// Endpoint API POST /tasks
	r.POST("/tasks", func(c *gin.Context) {
		var input TaskInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
			return
		}

		fmt.Println("⏳ Mengecek ke User Service untuk ID:", input.UserID)

		// =========================================================
		// INILAH INTI MICROSERVICES: KOMUNIKASI ANTAR SERVICE!
		// Task Service bertindak sebagai "Klien" yang menelepon API User Service
		// =========================================================
		urlUserService := "http://localhost:8081/users/validate/" + input.UserID

		// Melakukan HTTP GET Request ke service lain
		resp, err := http.Get(urlUserService)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghubungi User Service"})
			return
		}
		defer resp.Body.Close()

		// Membaca balasan dari User Service
		var userResp UserValidationResponse
		if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca balasan User Service"})
			return
		}

		// =========================================================
		// LOGIKA PENENTUAN BERDASARKAN JAWABAN USER SERVICE
		// =========================================================
		if !userResp.Valid {
			// Jika User Service bilang FALSE (Tidak Valid)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "DITOLAK",
				"pesan":  "User ID " + input.UserID + " tidak terdaftar di sistem pusat!",
			})
			return
		}

		// Jika User Service bilang TRUE (Valid)
		c.JSON(http.StatusCreated, gin.H{
			"status": "DITERIMA",
			"pesan":  "Tugas '" + input.NamaTugas + "' berhasil disimpan untuk User " + input.UserID,
		})
	})

	fmt.Println("📝 Task Service berjalan di PORT 8082...")
	// Kita jalankan di port 8082, berdampingan dengan user-service di 8081
	r.Run(":8082")
}
