// File: controllers/dokumencontroller.go
package controllers

import (
	"net/http"

	"eval-service/config"
	"eval-service/models"

	"github.com/gin-gonic/gin"
)

// 1. Fungsi GET
func TampilSemuaDokumen(c *gin.Context) {
	var daftarDokumen []models.Dokumen

	// Membaca dari Brankas lewat package config
	config.DB.Find(&daftarDokumen)
	c.JSON(http.StatusOK, daftarDokumen)
}

// 2. Fungsi POST
func TambahDokumen(c *gin.Context) {
	var dokumenBaru models.Dokumen

	if err := c.ShouldBindJSON(&dokumenBaru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah!"})
		return
	}

	// Menyimpan ke Brankas
	config.DB.Create(&dokumenBaru)
	c.JSON(http.StatusCreated, gin.H{
		"pesan": "Berhasil dikunci di dalam Brankas Baja PostgreSQL!",
		"data":  dokumenBaru,
	})
}

// 3. Fungsi DELETE
func HapusDokumen(c *gin.Context) {
	idTarget := c.Param("id")

	// Menghapus dari Brankas
	config.DB.Delete(&models.Dokumen{}, "id = ?", idTarget)
	c.JSON(http.StatusOK, gin.H{
		"pesan": "Dokumen " + idTarget + " berhasil dibumihanguskan!",
	})
}
