package repository

import (
	"latihan-fiber-api/internal/model"

	"gorm.io/gorm"
)

type ProgramStudiRepository struct {
	DB *gorm.DB
}

func NewProgramStudiRepository(db *gorm.DB) *ProgramStudiRepository {
	return &ProgramStudiRepository{DB: db}
}

func (r *ProgramStudiRepository) GetAllWithJumlahMahasiswa() ([]model.ProgramStudiDenganJumlahMahasiswa, error) {
	var data []model.ProgramStudiDenganJumlahMahasiswa

	err := r.DB.Model(&model.ProgramStudi{}).
		Select(`
			program_studis.id,
			program_studis.kode,
			program_studis.nama,
			program_studis.jenjang,
			COUNT(mahasiswas.id) AS jumlah_mahasiswa
		`).
		Joins("LEFT JOIN mahasiswas ON mahasiswas.program_studi_id = program_studis.id").
		Group("program_studis.id").
		Scan(&data).Error

	return data, err
}
