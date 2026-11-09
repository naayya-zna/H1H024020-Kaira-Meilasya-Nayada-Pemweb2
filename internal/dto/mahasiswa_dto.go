package dto

// CreateMahasiswaRequest digunakan untuk menerima
// data ketika menambahkan mahasiswa baru.
type CreateMahasiswaRequest struct {
	ProgramStudiID uint    `json:"program_studi_id"`
	Nim            string  `json:"nim"`
	Nama           string  `json:"nama"`
	Email          string  `json:"email"`
	Angkatan       int     `json:"angkatan"`
	Ipk            float64 `json:"ipk"`
	Aktif          bool    `json:"aktif"`
}

// UpdateMahasiswaRequest digunakan untuk menerima
// data ketika memperbarui mahasiswa.
type UpdateMahasiswaRequest struct {
	ProgramStudiID uint    `json:"program_studi_id"`
	Nim            string  `json:"nim"`
	Nama           string  `json:"nama"`
	Email          string  `json:"email"`
	Angkatan       int     `json:"angkatan"`
	Ipk            float64 `json:"ipk"`
	Aktif          bool    `json:"aktif"`
}
