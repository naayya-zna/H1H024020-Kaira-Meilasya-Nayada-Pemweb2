package service

import (
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

type ProgramStudiService struct {
	Repo *repository.ProgramStudiRepository
}

func NewProgramStudiService(repo *repository.ProgramStudiRepository) *ProgramStudiService {
	return &ProgramStudiService{Repo: repo}
}

func (s *ProgramStudiService) GetAllWithJumlahMahasiswa() ([]model.ProgramStudiDenganJumlahMahasiswa, error) {
	return s.Repo.GetAllWithJumlahMahasiswa()
}
