package config

import (
	"latihan-fiber-api/internal/model"
	"log"
)

// Migrasi membuat tabel berdasarkan model GORM.
func Migrasi() {
	err := DB.AutoMigrate(
		&model.ProgramStudi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
		&model.Pengguna{},
		&model.TokenBlacklist{},
	)

	if err != nil {
		log.Fatal("Gagal melakukan migrasi database: ", err)
	}

	log.Println("Migrasi database berhasil")
}

// DataAwal mengisi tabel program studi.
func DataAwal() {
	var jumlah int64

	// Periksa jumlah data program studi.
	DB.Model(&model.ProgramStudi{}).Count(&jumlah)

	// Jika sudah ada data, proses seeding dihentikan.
	if jumlah > 0 {
		return
	}

	// Data awal program studi.
	daftar := []model.ProgramStudi{
		{Kode: "TK", Nama: "Teknik Komputer", Jenjang: "S1"},
		{Kode: "IF", Nama: "Informatika", Jenjang: "S1"},
		{Kode: "TE", Nama: "Teknik Elektro", Jenjang: "S1"},
	}

	// Simpan data ke database.
	if err := DB.Create(&daftar).Error; err != nil {
		log.Fatal("Gagal membuat data awal: ", err)
	}

	log.Println("Data awal program studi dibuat")
}
