package service

import (
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

type MataKuliahService struct {
	Repo *repository.MataKuliahRepository
}

func NewMataKuliahService(repo *repository.MataKuliahRepository) *MataKuliahService {
	return &MataKuliahService{Repo: repo}
}

func (s *MataKuliahService) GetAll() ([]model.MataKuliah, error) {
	return s.Repo.GetAll()
}

func (s *MataKuliahService) GetByID(id uint) (model.MataKuliah, error) {
	return s.Repo.GetByID(id)
}

func (s *MataKuliahService) Create(data *model.MataKuliah) error {
	return s.Repo.Create(data)
}

func (s *MataKuliahService) Update(id uint, data map[string]interface{}) error {
	return s.Repo.Update(id, data)
}

func (s *MataKuliahService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
