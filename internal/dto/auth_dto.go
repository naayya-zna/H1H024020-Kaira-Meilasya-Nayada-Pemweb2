package dto

type PermintaanRegistrasi struct {
	Nama      string `json:"nama" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required,min=8"`
}

type PermintaanLogin struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}
