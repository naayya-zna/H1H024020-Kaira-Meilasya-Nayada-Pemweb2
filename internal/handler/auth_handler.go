package handler

import (
	"errors"
	"strings"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/dto"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/validasi"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Registrasi mendaftarkan pengguna baru.
func Registrasi(c fiber.Ctx) error {
	var request dto.PermintaanRegistrasi

	// Membaca data JSON dari request.
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	// Memeriksa aturan validasi.
	galat := validasi.Periksa(request)
	if len(galat) > 0 {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data registrasi tidak valid",
			"galat":  galat,
		})
	}

	// Membersihkan spasi pada nama dan email.
	email := strings.TrimSpace(strings.ToLower(request.Email))
	nama := strings.TrimSpace(request.Nama)

	// Memeriksa apakah email sudah terdaftar.
	var penggunaLama model.Pengguna
	err := config.DB.Where("email = ?", email).First(&penggunaLama).Error

	if err == nil {
		return c.Status(409).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Email sudah terdaftar",
		})
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memeriksa email pengguna",
		})
	}

	// Mengubah password menjadi hash menggunakan bcrypt.
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.KataSandi),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengamankan kata sandi",
		})
	}

	// Membuat data pengguna baru.
	pengguna := model.Pengguna{
		Nama:      nama,
		Email:     email,
		KataSandi: string(hashPassword),
		Peran:     "mahasiswa",
	}

	if err := config.DB.Create(&pengguna).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mendaftarkan pengguna",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Registrasi berhasil",
		"data": fiber.Map{
			"id":    pengguna.ID,
			"nama":  pengguna.Nama,
			"email": pengguna.Email,
			"peran": pengguna.Peran,
		},
	})
}

// Login memeriksa identitas pengguna dan menerbitkan JWT.
func Login(c fiber.Ctx) error {
	var request dto.PermintaanLogin

	// Membaca data JSON dari request.
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	// Memeriksa aturan validasi.
	galat := validasi.Periksa(request)
	if len(galat) > 0 {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data login tidak valid",
			"galat":  galat,
		})
	}

	email := strings.TrimSpace(strings.ToLower(request.Email))

	// Mencari pengguna berdasarkan email.
	var pengguna model.Pengguna
	err := config.DB.Where("email = ?", email).First(&pengguna).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Email atau kata sandi salah",
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data pengguna",
		})
	}

	// Membandingkan password dengan hash di database.
	err = bcrypt.CompareHashAndPassword(
		[]byte(pengguna.KataSandi),
		[]byte(request.KataSandi),
	)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Email atau kata sandi salah",
		})
	}

	// Menerbitkan JWT setelah login berhasil.
	token, err := auth.TerbitkanToken(pengguna, config.Load())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal membuat token autentikasi",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"data": fiber.Map{
			"token": token,
			"tipe":  "Bearer",
			"pengguna": fiber.Map{
				"id":    pengguna.ID,
				"nama":  pengguna.Nama,
				"email": pengguna.Email,
				"peran": pengguna.Peran,
			},
		},
	})
}

// Profil menampilkan informasi pengguna yang sedang login.
func Profil(c fiber.Ctx) error {
	// ID pengguna akan disediakan oleh middleware autentikasi.
	idPengguna, ok := c.Locals("id_pengguna").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Pengguna belum terautentikasi",
		})
	}

	var pengguna model.Pengguna

	if err := config.DB.First(&pengguna, idPengguna).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Pengguna tidak ditemukan",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil profil pengguna",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Profil pengguna berhasil diambil",
		"data": fiber.Map{
			"id":    pengguna.ID,
			"nama":  pengguna.Nama,
			"email": pengguna.Email,
			"peran": pengguna.Peran,
		},
	})
}

// Segarkan menerbitkan token JWT baru berdasarkan token yang masih valid.
func Segarkan(c fiber.Ctx) error {
	// Ambil ID pengguna dari middleware autentikasi.
	idPengguna, ok := c.Locals("id_pengguna").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Pengguna tidak terautentikasi",
		})
	}

	// Ambil peran pengguna dari token lama.
	peran, ok := c.Locals("peran").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Informasi peran tidak ditemukan",
		})
	}

	// Bentuk data pengguna untuk menerbitkan token baru.
	pengguna := model.Pengguna{
		ID:    idPengguna,
		Peran: peran,
	}

	// Terbitkan token baru.
	tokenBaru, err := auth.TerbitkanToken(pengguna, config.Load())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menerbitkan token baru",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil diperbarui",
		"data": fiber.Map{
			"token": tokenBaru,
		},
	})
}

// Logout mencabut token JWT agar tidak dapat digunakan kembali.
func Logout(c fiber.Ctx) error {
	// Mengambil token dari header Authorization.
	parts := strings.Fields(c.Get("Authorization"))

	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Token autentikasi tidak ditemukan",
		})
	}

	tokenString := parts[1]

	// Memeriksa token dan mengambil klaimnya.
	claims, err := auth.PeriksaToken(tokenString, config.Load())
	if err != nil || claims.ExpiresAt == nil {
		return c.Status(401).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Token tidak valid atau sudah kedaluwarsa",
		})
	}

	// Membuat data blacklist menggunakan hash token.
	blacklist := model.TokenBlacklist{
		TokenHash:       auth.HashToken(tokenString),
		KedaluwarsaPada: claims.ExpiresAt.Time,
	}

	// Menyimpan token yang dicabut ke database.
	if err := config.DB.Create(&blacklist).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mencabut token",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil",
	})
}
