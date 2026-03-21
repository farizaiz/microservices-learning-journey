// File: models/dokumen.go
package models

// Perhatikan nama package-nya sekarang "models", bukan "main" lagi.
// Ini karena file ini berada di dalam folder models.

type Dokumen struct {
	ID         string `gorm:"primaryKey" json:"id_dokumen"`
	Perusahaan string `json:"nama_perusahaan"`
	Status     string `json:"status_lingkungan"`
	Skor       int    `json:"skor_evaluasi"`
}
