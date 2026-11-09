package model

import (
	"time"

	"gorm.io/gorm"
)

// Model Program Studi
type ProgramStudi struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Kode      string    `gorm:"size:10;uniqueIndex" json:"kode"`
	Nama      string    `gorm:"size:100" json:"nama"`
	Jenjang   string    `gorm:"size:10" json:"jenjang"`
	CreatedAt time.Time `json:"dibuat_pada"`
	UpdatedAt time.Time `json:"diubah_pada"`
}

// Model Mahasiswa
type Mahasiswa struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ProgramStudiID uint           `gorm:"index" json:"program_studi_id"`
	ProgramStudi   *ProgramStudi  `gorm:"foreignKey:ProgramStudiID" json:"program_studi,omitempty"`
	Nim            string         `gorm:"size:20;uniqueIndex" json:"nim"`
	Nama           string         `gorm:"size:100" json:"nama"`
	Email          string         `gorm:"size:100;uniqueIndex" json:"email"`
	Angkatan       int            `json:"angkatan"`
	Ipk            float64        `gorm:"type:decimal(4,2)" json:"ipk"`
	Aktif          bool           `gorm:"default:true" json:"aktif"`
	CreatedAt      time.Time      `json:"dibuat_pada"`
	UpdatedAt      time.Time      `json:"diubah_pada"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
type ProgramStudiDenganJumlahMahasiswa struct {
	ID              uint   `json:"id"`
	Kode            string `json:"kode"`
	Nama            string `json:"nama"`
	Jenjang         string `json:"jenjang"`
	JumlahMahasiswa int64  `json:"jumlah_mahasiswa"`
}
