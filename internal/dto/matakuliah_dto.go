package dto

type CreateMataKuliahRequest struct {
	ProgramStudiID uint   `json:"program_studi_id"`
	Kode           string `json:"kode"`
	Nama           string `json:"nama"`
	SKS            int    `json:"sks"`
	Semester       int    `json:"semester"`
}

type UpdateMataKuliahRequest struct {
	ProgramStudiID uint   `json:"program_studi_id"`
	Kode           string `json:"kode"`
	Nama           string `json:"nama"`
	SKS            int    `json:"sks"`
	Semester       int    `json:"semester"`
}
