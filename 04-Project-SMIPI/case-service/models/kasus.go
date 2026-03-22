package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Kasus merepresentasikan entitas utama dalam Case Service
type Kasus struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	NomorLP         string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"nomor_lp"`
	DeskripsiKasus  string    `gorm:"type:text;not null" json:"deskripsi_kasus"`
	TanggalKejadian time.Time `gorm:"not null" json:"tanggal_kejadian"`

	// Lokasi (Sangat krusial untuk pemetaan/GIS)
	LokasiKejadian string  `gorm:"type:varchar(255);not null" json:"lokasi_kejadian"`
	Latitude       float64 `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude      float64 `gorm:"type:decimal(11,8)" json:"longitude"`

	// Relasi Lintas Service (Disimpan sebagai ID saja karena profil ada di User Service)
	PelaporID      string `gorm:"type:varchar(50);index;not null" json:"pelapor_id"`
	PenyidikID     string `gorm:"type:varchar(50);index" json:"penyidik_id"`
	UnitPenanganan string `gorm:"type:varchar(100)" json:"unit_penanganan"`

	// Klasifikasi & Status (Bisa menggunakan tipe Enum di database atau referensi tabel lain)
	StatusID        string `gorm:"type:varchar(50);index;default:'DRAFT'" json:"status_id"`
	TingkatBahayaID string `gorm:"type:varchar(50);index" json:"tingkat_bahaya_id"`
	KategoriKasusID string `gorm:"type:varchar(50);index" json:"kategori_kasus_id"`
	Prioritas       string `gorm:"type:varchar(20);index;default:'MEDIUM'" json:"prioritas"`

	// Tracking Waktu Investigasi
	TglMulaiLidik   *time.Time `json:"tgl_mulai_lidik"` // Menggunakan pointer (*) agar bisa bernilai NULL di awal
	TglSelesaiLidik *time.Time `json:"tgl_selesai_lidik"`

	// Metadata & Audit Trail (Wajib untuk Enterprise)
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft Delete GORM
	CreatedBy string         `gorm:"type:varchar(50)" json:"created_by"`
	UpdatedBy string         `gorm:"type:varchar(50)" json:"updated_by"`
}
