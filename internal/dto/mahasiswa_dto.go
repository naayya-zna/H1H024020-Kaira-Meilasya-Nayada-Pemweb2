package dto

// CreateMahasiswaRequest digunakan untuk menerima
// data ketika menambahkan mahasiswa baru.
type CreateMahasiswaRequest struct {
	ProgramStudiID uint    `json:"program_studi_id" validate:"required,gt=0"`
	Nim            string  `json:"nim" validate:"required,nim"`
	Nama           string  `json:"nama" validate:"required,min=3,max=100"`
	Email          string  `json:"email" validate:"required,email"`
	Angkatan       int     `json:"angkatan" validate:"required,gt=0"`
	Ipk            float64 `json:"ipk" validate:"gte=0,lte=4"`
	Aktif          bool    `json:"aktif"`
}

// UpdateMahasiswaRequest digunakan untuk menerima
// data ketika memperbarui mahasiswa.
type UpdateMahasiswaRequest struct {
	ProgramStudiID uint    `json:"program_studi_id"`
	Nim            string  `json:"nim" validate:"required,nim"`
	Nama           string  `json:"nama" validate:"required,min=3,max=100"`
	Email          string  `json:"email"`
	Angkatan       int     `json:"angkatan"`
	Ipk            float64 `json:"ipk"`
	Aktif          bool    `json:"aktif"`
}
