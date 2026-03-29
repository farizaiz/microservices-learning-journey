package controllers

import (
	"fmt"
	"net/http"

	"case-service/config"
	"case-service/models"

	"github.com/gin-gonic/gin"
)

// ==============================================================================
// STRUKTUR DATA (Payload / DTO)
// ==============================================================================

// UpdateStatusInput adalah format JSON dari React
type UpdateStatusInput struct {
	StatusID  string `json:"status_id" binding:"required"` // Misalnya: "VALIDASI"
	Prioritas string `json:"prioritas"`                    // Misalnya: "URGENT" (Opsional)
}

// ==============================================================================
// HANDLER FUNGSI-FUNGSI KASUS (CONTROLLERS)
// ==============================================================================

// CreateKasus menerima laporan kasus baru
func CreateKasus(c *gin.Context) {
	var kasusBaru models.Kasus

	if err := c.ShouldBindJSON(&kasusBaru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "GAGAL", "pesan": err.Error()})
		return
	}

	// CATAT SIAPA YANG MEMBUAT KASUS INI
	userID, exists := c.Get("user_id")
	if exists {
		if userIDStr, ok := userID.(string); ok {
			kasusBaru.CreatedBy = userIDStr

			// Jika PelaporID kosong dari frontend, otomatis isi dengan user yang sedang login
			if kasusBaru.PelaporID == "" {
				kasusBaru.PelaporID = userIDStr
			}
		}
	}

	if err := config.DB.Create(&kasusBaru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "GAGAL", "pesan": "Kesalahan database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "BERHASIL", "data": kasusBaru})
}

// GetAllKasus mengambil semua daftar laporan (Papan Kendali)
func GetAllKasus(c *gin.Context) {
	var daftarKasus []models.Kasus

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
	id := c.Param("id")

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

// UpdateKasus memperbarui data laporan secara menyeluruh (Edit Kasus)
func UpdateKasus(c *gin.Context) {
	var kasus models.Kasus
	id := c.Param("id")

	// Ambil data siapa yang mengedit
	userID, _ := c.Get("user_id")

	if err := config.DB.First(&kasus, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "GAGAL",
			"pesan":  "Laporan dengan ID tersebut tidak ditemukan!",
		})
		return
	}

	if err := c.ShouldBindJSON(&kasus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "GAGAL",
			"pesan":  "Format data tidak valid: " + err.Error(),
		})
		return
	}

	config.DB.Save(&kasus)
	fmt.Printf("Kasus %v diedit oleh User ID: %v\n", id, userID)

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Laporan Kasus Investigasi berhasil diperbarui!",
		"data":   kasus,
	})
}

func UpdateStatusKasus(c *gin.Context) {
	var kasus models.Kasus
	id := c.Param("id")

	userID, _ := c.Get("user_id")

	// Cari data kasusnya
	if err := config.DB.First(&kasus, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "GAGAL", "pesan": "Laporan tidak ditemukan!"})
		return
	}

	// Tangkap data JSON dari React
	var input UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "GAGAL", "pesan": "Format status tidak valid!"})
		return
	}

	// Gunakan StatusID, bukan TindakanOperasional
	kasus.StatusID = input.StatusID

	// Jika dari frontend juga mengirim prioritas baru, kita update juga
	if input.Prioritas != "" {
		kasus.Prioritas = input.Prioritas
	}

	// Catat siapa yang melakukan update (dari token JWT)
	if userIDStr, ok := userID.(string); ok {
		kasus.UpdatedBy = userIDStr
	}

	// Simpan perubahan ke database
	config.DB.Save(&kasus)

	// Kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Status kasus berhasil diperbarui menjadi " + kasus.StatusID,
		"data": gin.H{
			"kasus_id":   id,
			"status":     kasus.StatusID,
			"prioritas":  kasus.Prioritas,
			"updated_by": kasus.UpdatedBy,
		},
	})
}

// DeleteKasus menghapus laporan menggunakan mekanisme Soft Delete GORM
func DeleteKasus(c *gin.Context) {
	var kasus models.Kasus
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	if err := config.DB.First(&kasus, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "GAGAL",
			"pesan":  "Laporan dengan ID tersebut tidak ditemukan!",
		})
		return
	}

	config.DB.Delete(&kasus)
	fmt.Printf("Kasus %v diarsipkan (Soft Delete) oleh User ID: %v\n", id, userID)

	c.JSON(http.StatusOK, gin.H{
		"status": "BERHASIL",
		"pesan":  "Laporan Kasus berhasil dihapus (disembunyikan dari sistem)!",
	})
}
