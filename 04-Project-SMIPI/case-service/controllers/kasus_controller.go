package controllers

import (
	"net/http"

	"case-service/config"
	"case-service/models"

	"github.com/gin-gonic/gin"
)

// CreateKasus adalah fungsi untuk menerima laporan kasus baru
func CreateKasus(c *gin.Context) {
	var kasusBaru models.Kasus

	// 1. Menangkap data JSON yang dikirim oleh user (misal dari Insomnia/Frontend)
	if err := c.ShouldBindJSON(&kasusBaru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "GAGAL",
			"pesan":  "Format data tidak valid: " + err.Error(),
		})
		return
	}

	// 2. Menyimpan data tersebut ke PostgreSQL menggunakan GORM
	if err := config.DB.Create(&kasusBaru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "GAGAL",
			"pesan":  "Terjadi kesalahan saat menyimpan ke database",
		})
		return
	}

	// 3. Mengembalikan jawaban sukses beserta data yang berhasil disimpan
	c.JSON(http.StatusCreated, gin.H{
		"status": "BERHASIL",
		"pesan":  "Laporan Kasus Investigasi berhasil dicatat!",
		"data":   kasusBaru,
	})
}

// GetAllKasus mengambil semua daftar laporan dari database
func GetAllKasus(c *gin.Context) {
	// Karena datanya banyak, kita gunakan Slice (Array) dari model Kasus
	var daftarKasus []models.Kasus

	// GORM akan otomatis melakukan: SELECT * FROM kasus;
	if err := config.DB.Find(&daftarKasus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "GAGAL",
			"pesan":  "Terjadi kesalahan saat mengambil data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "BERHASIL",
		"total_data": len(daftarKasus),
		"data":       daftarKasus,
	})
}

// GetKasusByID mengambil satu detail laporan spesifik
func GetKasusByID(c *gin.Context) {
	var kasus models.Kasus

	// Menangkap ID dari URL (misal: /kasus/510d53dc-...)
	id := c.Param("id")

	// GORM akan otomatis melakukan: SELECT * FROM kasus WHERE id = '...';
	if err := config.DB.First(&kasus, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "GAGAL",
			"pesan":  "Laporan dengan ID tersebut tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"data":   kasus,
	})
}

// UpdateKasus memperbarui data laporan yang sudah ada
func UpdateKasus(c *gin.Context) {
	var kasus models.Kasus
	id := c.Param("id")

	// 1. Cek dulu apakah kasus dengan ID tersebut ada di database?
	if err := config.DB.First(&kasus, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "GAGAL",
			"pesan":  "Laporan dengan ID tersebut tidak ditemukan!",
		})
		return
	}

	// 2. Tangkap data perubahan dari user
	if err := c.ShouldBindJSON(&kasus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "GAGAL",
			"pesan":  "Format data tidak valid: " + err.Error(),
		})
		return
	}

	// 3. Simpan perubahannya ke database
	config.DB.Save(&kasus)

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Laporan Kasus Investigasi berhasil diperbarui!",
		"data":   kasus,
	})
}

// DeleteKasus menghapus laporan (Soft Delete)
func DeleteKasus(c *gin.Context) {
	var kasus models.Kasus
	id := c.Param("id")

	// 1. Cek keberadaan data
	if err := config.DB.First(&kasus, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "GAGAL",
			"pesan":  "Laporan dengan ID tersebut tidak ditemukan!",
		})
		return
	}

	// 2. Lakukan penghapusan
	// KARENA KITA PAKAI gorm.DeletedAt, GORM otomatis hanya akan mengisi kolom deleted_at (Soft Delete)
	// Datanya TETAP ADA di database demi keamanan audit!
	config.DB.Delete(&kasus)

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Laporan Kasus berhasil dihapus (disembunyikan dari sistem)!",
	})
}
