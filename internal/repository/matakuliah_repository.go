package repository

import (
	"latihan-fiber-api/internal/model"

	"gorm.io/gorm"
)

type MataKuliahRepository struct {
	DB *gorm.DB
}

func NewMataKuliahRepository(db *gorm.DB) *MataKuliahRepository {
	return &MataKuliahRepository{DB: db}
}

func (r *MataKuliahRepository) GetAll() ([]model.MataKuliah, error) {
	var data []model.MataKuliah

	err := r.DB.Preload("ProgramStudi").Find(&data).Error

	return data, err
}

func (r *MataKuliahRepository) GetByID(id uint) (model.MataKuliah, error) {
	var data model.MataKuliah

	err := r.DB.Preload("ProgramStudi").First(&data, id).Error

	return data, err
}

func (r *MataKuliahRepository) Create(data *model.MataKuliah) error {
	return r.DB.Create(data).Error
}

func (r *MataKuliahRepository) Update(id uint, data map[string]interface{}) error {
	return r.DB.Model(&model.MataKuliah{}).
		Where("id = ?", id).
		Updates(data).Error
}

func (r *MataKuliahRepository) Delete(id uint) error {
	return r.DB.Delete(&model.MataKuliah{}, id).Error
}
