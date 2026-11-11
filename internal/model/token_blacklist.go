package model

import "time"

// TokenBlacklist menyimpan token JWT yang sudah dicabut.
type TokenBlacklist struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Hash token digunakan agar token asli tidak disimpan di database.
	TokenHash string `gorm:"size:64;uniqueIndex;not null" json:"-"`

	// Waktu kedaluwarsa token.
	KedaluwarsaPada time.Time `gorm:"not null" json:"kedaluwarsa_pada"`

	CreatedAt time.Time `json:"dibuat_pada"`
}

// TableName menentukan nama tabel di database.
func (TokenBlacklist) TableName() string {
	return "token_blacklists"
}
