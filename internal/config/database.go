package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB menyimpan koneksi database secara global.
var DB *gorm.DB

// HubungkanBasisData menghubungkan aplikasi ke MySQL.
func HubungkanBasisData(konf Config) {

	// Membuat format koneksi database (DSN).
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		konf.DbUser,
		konf.DbSandi,
		konf.DbHost,
		konf.DbPort,
		konf.DbNama,
	)

	// Membuka koneksi menggunakan GORM.
	koneksi, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Memeriksa apakah koneksi berhasil.
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}

	// Menyimpan koneksi ke variabel global.
	DB = koneksi

	log.Println("Koneksi database berhasil")
}
