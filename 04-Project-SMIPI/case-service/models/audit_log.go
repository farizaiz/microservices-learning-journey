package models

import (
	"time"

	"github.com/google/uuid"
)

// AuditLog mencatat setiap jejak rekam (history) perubahan pada kasus
type AuditLog struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	KasusID    uuid.UUID `gorm:"type:uuid;not null;index" json:"kasus_id"`
	UserID     string    `gorm:"type:varchar(50);not null;index" json:"user_id"` // Disamakan dengan PenyidikID (varchar 50)
	Action     string    `gorm:"type:varchar(50);not null" json:"action"`        // CREATE, UPDATE_STATUS, DELETE
	Keterangan string    `gorm:"type:text;not null" json:"keterangan"`
	CreatedAt  time.Time `json:"created_at"`
}
