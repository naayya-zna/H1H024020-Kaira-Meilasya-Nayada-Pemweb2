package validasi

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Membuat objek validator.
var validate = validator.New()

// Pola NIM: huruf H kapital diikuti angka.
var polaNIM = regexp.MustCompile(`^H[0-9]+$`)

// Mendaftarkan custom validation.
func init() {
	err := validate.RegisterValidation("nim", func(fl validator.FieldLevel) bool {
		return polaNIM.MatchString(fl.Field().String())
	})

	if err != nil {
		panic(err)
	}
}

// PesanGalat menyimpan informasi kesalahan validasi.
type PesanGalat struct {
	Field string `json:"field"`
	Pesan string `json:"pesan"`
}

// Periksa memeriksa apakah data sudah sesuai aturan validasi.
func Periksa(data any) []PesanGalat {
	var galat []PesanGalat

	err := validate.Struct(data)
	if err == nil {
		return galat
	}

	// Mengambil setiap kesalahan validasi.
	for _, e := range err.(validator.ValidationErrors) {
		pesan := susunPesan(e)

		galat = append(galat, PesanGalat{
			Field: strings.ToLower(e.Field()),
			Pesan: pesan,
		})
	}

	return galat
}

// susunPesan mengubah kesalahan validasi menjadi pesan bahasa Indonesia.
func susunPesan(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "Field wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Jumlah karakter terlalu sedikit"
	case "max":
		return "Jumlah karakter terlalu banyak"
	case "gte":
		return "Nilai terlalu kecil"
	case "lte":
		return "Nilai terlalu besar"
	case "nim":
		return "Format NIM harus diawali huruf H kapital dan diikuti angka"
	default:
		return "Data tidak valid"
	}
}
