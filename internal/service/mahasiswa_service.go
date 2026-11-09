package service

import (
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

// MahasiswaService menangani logika bisnis mahasiswa.
type MahasiswaService struct {
	Repo *repository.MahasiswaRepository
}

// NewMahasiswaService membuat service mahasiswa baru.
func NewMahasiswaService(repo *repository.MahasiswaRepository) *MahasiswaService {
	return &MahasiswaService{
		Repo: repo,
	}
}

// GetAll mengambil seluruh data mahasiswa dengan pagination.
func (s *MahasiswaService) GetAll(
	page int,
	limit int,
	angkatan int,
	urut string,
	cari string,
) ([]model.Mahasiswa, int64, int, error) {
	return s.Repo.GetPaginated(page, limit, angkatan, urut, cari)
}

// GetByID mengambil mahasiswa berdasarkan ID.
func (s *MahasiswaService) GetByID(id uint) (model.Mahasiswa, error) {
	return s.Repo.GetByID(id)
}

// GetByNIM mengambil mahasiswa berdasarkan NIM.
func (s *MahasiswaService) GetByNIM(nim string) (model.Mahasiswa, error) {
	return s.Repo.GetByNIM(nim)
}

// Create menambahkan data mahasiswa.
func (s *MahasiswaService) Create(mahasiswa *model.Mahasiswa) error {
	return s.Repo.Create(mahasiswa)
}

// Update memperbarui data mahasiswa.
func (s *MahasiswaService) Update(id uint, data map[string]interface{}) error {
	return s.Repo.Update(id, data)
}

// Delete menghapus data mahasiswa.
func (s *MahasiswaService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
func (s *MahasiswaService) GetByIDIncludingDeleted(id uint) (model.Mahasiswa, error) {
	return s.Repo.GetByIDIncludingDeleted(id)
}

func (s *MahasiswaService) Restore(id uint) error {
	return s.Repo.Restore(id)
}
