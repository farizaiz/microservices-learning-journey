package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 1. Tambahkan Tag GORM ke Struct Anda
// gorm:"primaryKey" memberi tahu database bahwa ID ini adalah kunci utamanya
type Dokumen struct {
	ID         string `gorm:"primaryKey" json:"id_dokumen"`
	Perusahaan string `json:"nama_perusahaan"`
	Status     string `json:"status_lingkungan"`
	Skor       int    `json:"skor_evaluasi"`
}

// Variabel Global untuk menyimpan mesin Database kita
var DB *gorm.DB

func main() {
	// 2. DSN (Data Source Name) - Alamat Lengkap Menuju Gudang Anda
	dsn := "host=localhost user=dummy password=dummy dbname=speed_klhk port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// 3. Mengetuk Pintu Database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Jika gagal connect, program langsung dihentikan (Fatal)
		log.Fatal("❌ Gagal terhubung ke Database: ", err)
	}
	fmt.Println("✅ BINGO! Berhasil masuk ke Database PostgreSQL!")

	// 4. AUTO-MIGRATE (Keajaiban GORM)
	// GORM akan membaca Struct "Dokumen", lalu berlari ke PostgreSQL
	// dan otomatis membuatkan tabelnya untuk Anda!
	DB.AutoMigrate(&Dokumen{})
	fmt.Println("✅ Tabel 'dokumens' berhasil disinkronisasi ke Database!")

	// --- SETUP GIN ROUTER ---
	r := gin.Default()

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"aplikasi": "SPEED KLHK",
			"status":   "🟢 ONLINE & TERHUBUNG KE DATABASE",
		})
	})

	// ==========================================
	// 1. GET: Mengambil Semua Dokumen dari Database
	// ==========================================
	r.GET("/api/dokumen", func(c *gin.Context) {
		var daftarDokumen []Dokumen

		// KEAJAIBAN GORM: 1 baris ini setara dengan "SELECT * FROM dokumens"
		DB.Find(&daftarDokumen)

		c.JSON(http.StatusOK, daftarDokumen)
	})

	// ==========================================
	// 2. POST: Menyimpan Dokumen Baru ke Database
	// ==========================================
	r.POST("/api/dokumen", func(c *gin.Context) {
		var dokumenBaru Dokumen

		if err := c.ShouldBindJSON(&dokumenBaru); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah!"})
			return
		}

		// KEAJAIBAN GORM: 1 baris ini setara dengan "INSERT INTO dokumens ..."
		DB.Create(&dokumenBaru)

		c.JSON(http.StatusCreated, gin.H{
			"pesan": "Berhasil dikunci di dalam Brankas Baja PostgreSQL!",
			"data":  dokumenBaru,
		})
	})

	// ==========================================
	// 3. DELETE: Menghapus Dokumen dari Database
	// ==========================================
	r.DELETE("/api/dokumen/:id", func(c *gin.Context) {
		idTarget := c.Param("id")

		// KEAJAIBAN GORM: 1 baris ini setara dengan "DELETE FROM dokumens WHERE id = ?"
		DB.Delete(&Dokumen{}, "id = ?", idTarget)

		c.JSON(http.StatusOK, gin.H{
			"pesan": "Dokumen " + idTarget + " berhasil dibumihanguskan dari Database!",
		})
	})

	// Nyalakan Server
	r.Run(":8080")
}
