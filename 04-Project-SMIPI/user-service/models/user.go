package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User merepresentasikan entitas Penyidik, Pelapor, atau Admin
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	NamaLengkap  string    `gorm:"type:varchar(100);not null" json:"nama_lengkap"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"type:varchar(255);not null" json:"password"`
	Role         string    `gorm:"type:varchar(20);default:'PELAPOR'" json:"role"` // ADMIN, PENYIDIK, PELAPOR
	NomorTelepon string    `gorm:"type:varchar(20)" json:"nomor_telepon"`
	NomorInduk   string    `gorm:"type:varchar(50);uniqueIndex" json:"nomor_induk"` // NRP untuk Polisi / NIP untuk ASN KLHK

	// Audit Trail
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
