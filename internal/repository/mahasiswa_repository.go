package repository

import (
	"latihan-fiber-api/internal/model"
	"math"

	"gorm.io/gorm"
)

// MahasiswaRepository menyimpan koneksi database.
type MahasiswaRepository struct {
	DB *gorm.DB
}

// NewMahasiswaRepository membuat repository mahasiswa baru.
func NewMahasiswaRepository(db *gorm.DB) *MahasiswaRepository {
	return &MahasiswaRepository{
		DB: db,
	}
}

// GetAll mengambil seluruh data mahasiswa.
func (r *MahasiswaRepository) GetAll() ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa

	err := r.DB.Preload("ProgramStudi").
		Order("id ASC").
		Find(&mahasiswa).Error

	return mahasiswa, err
}

// GetByID mencari mahasiswa berdasarkan ID.
func (r *MahasiswaRepository) GetByID(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := r.DB.Preload("ProgramStudi").
		First(&mahasiswa, id).Error

	return mahasiswa, err
}

// GetByNIM mencari mahasiswa berdasarkan NIM.
func (r *MahasiswaRepository) GetByNIM(nim string) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := r.DB.Preload("ProgramStudi").
		Where("nim = ?", nim).
		First(&mahasiswa).Error

	return mahasiswa, err
}

// Create menyimpan data mahasiswa baru.
func (r *MahasiswaRepository) Create(mahasiswa *model.Mahasiswa) error {
	return r.DB.Create(mahasiswa).Error
}

// Update memperbarui data mahasiswa.
func (r *MahasiswaRepository) Update(id uint, data map[string]interface{}) error {
	return r.DB.Model(&model.Mahasiswa{}).
		Where("id = ?", id).
		Updates(data).Error
}

// Delete menghapus data mahasiswa.
func (r *MahasiswaRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Mahasiswa{}, id).Error
}

func (r *MahasiswaRepository) GetPaginated(
	page int,
	limit int,
	angkatan int,
	urut string,
	cari string,
) ([]model.Mahasiswa, int64, int, error) {

	var mahasiswa []model.Mahasiswa
	var total int64

	// Membuat query dasar.
	query := r.DB.Model(&model.Mahasiswa{})

	// Filter berdasarkan angkatan.
	if angkatan > 0 {
		query = query.Where("angkatan = ?", angkatan)
	}

	// Pencarian berdasarkan nama atau NIM.
	if cari != "" {
		pola := "%" + cari + "%"
		query = query.Where("nama LIKE ? OR nim LIKE ?", pola, pola)
	}

	// Menghitung jumlah data setelah filter.
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	// Menghitung jumlah halaman.
	totalHalaman := int(math.Ceil(float64(total) / float64(limit)))

	// Menentukan posisi awal data.
	offset := (page - 1) * limit

	// Mengatur pengurutan data.
	switch urut {
	case "nama":
		query = query.Order("nama ASC")
	case "ipk":
		query = query.Order("ipk DESC")
	default:
		query = query.Order("id ASC")
	}

	// Mengambil data sesuai halaman.
	err := query.Preload("ProgramStudi").
		Limit(limit).
		Offset(offset).
		Find(&mahasiswa).Error

	return mahasiswa, total, totalHalaman, err
}

// GetByIDIncludingDeleted mengambil data termasuk yang sudah dihapus.
func (r *MahasiswaRepository) GetByIDIncludingDeleted(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := r.DB.Unscoped().
		Preload("ProgramStudi").
		First(&mahasiswa, id).Error

	return mahasiswa, err
}

// Restore mengembalikan mahasiswa yang sudah di-soft delete.
func (r *MahasiswaRepository) Restore(id uint) error {
	return r.DB.Unscoped().
		Model(&model.Mahasiswa{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}
