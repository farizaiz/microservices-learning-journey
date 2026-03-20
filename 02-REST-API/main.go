package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dokumen struct {
	ID         string `json:"id_dokumen"`
	Perusahaan string `json:"nama_perusahaan"`
	Status     string `json:"status_lingkungan"`
	Skor       int    `json:"skor_evaluasi"`
}

// SIMULASI DATABASE: Kita letakkan di luar (Global) agar datanya tidak ter-reset
var daftarAntrean = []Dokumen{
	{ID: "DOC-001", Perusahaan: "PT. Tambang Emas", Status: "Aman", Skor: 85},
	{ID: "DOC-002", Perusahaan: "PT. Asap Hitam", Status: "Kritis", Skor: 40},
}

func main() {
	r := gin.Default()

	// 1. Endpoint GET Status
	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "🟢 ONLINE"})
	})

	// 2. Endpoint GET Dokumen (Membaca Data)
	r.GET("/api/dokumen", func(c *gin.Context) {
		c.JSON(http.StatusOK, daftarAntrean)
	})

	// 3. ENDPOINT BARU: POST Dokumen (Menerima Data)
	r.POST("/api/dokumen", func(c *gin.Context) {
		var dokumenBaru Dokumen

		// Ajaibnya Gin: Menangkap teks JSON dari Insomnia,
		// lalu memasukkannya (Bind) ke dalam Struct Golang kita secara otomatis!
		if err := c.ShouldBindJSON(&dokumenBaru); err != nil {
			// Jika format JSON-nya salah/cacat, tolak mentah-mentah!
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Jika sukses, masukkan data baru itu ke dalam "Database" (Slice) kita
		daftarAntrean = append(daftarAntrean, dokumenBaru)

		// Berikan balasan sukses ke pengirim
		c.JSON(http.StatusCreated, gin.H{
			"pesan": "Dokumen berhasil disubmit ke sistem SPEED!",
			"data":  dokumenBaru,
		})
	})
	// 4. ENDPOINT BARU: DELETE Dokumen (Menghapus Data)
	// Tanda titik dua (:id) berarti itu adalah variabel dinamis yang bisa berubah-ubah
	r.DELETE("/api/dokumen/:id", func(c *gin.Context) {

		// 1. Tangkap ID yang diketik user di URL
		idTarget := c.Param("id")

		// 2. Cari dokumen tersebut di dalam "Database" (Slice) kita
		for i, dok := range daftarAntrean {
			if dok.ID == idTarget {

				// JURUS RAHASIA GOLANG: Menghapus elemen dari tengah Slice
				// Kita menyambungkan elemen "sebelum" target, dengan elemen "sesudah" target.
				// Elemen targetnya otomatis tertimpa dan lenyap!
				daftarAntrean = append(daftarAntrean[:i], daftarAntrean[i+1:]...)

				// Balas dengan pesan sukses
				c.JSON(http.StatusOK, gin.H{
					"pesan": "Dokumen " + idTarget + " berhasil dibakar dari sistem!",
				})
				return // Hentikan pencarian karena dokumen sudah ketemu dan dihapus
			}
		}

		// 3. Jika perulangan (looping) selesai tapi ID tidak ditemukan
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Gagal menghapus. Dokumen dengan ID " + idTarget + " tidak ditemukan.",
		})
	})

	r.Run(":8080")
}
