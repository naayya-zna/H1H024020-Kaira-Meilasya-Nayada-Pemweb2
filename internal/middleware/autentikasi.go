package middleware

import (
	"strconv"
	"strings"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"

	"github.com/gofiber/fiber/v3"
)

// Autentikasi memeriksa token JWT dari request.
func Autentikasi() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Mengambil Authorization Header.
		authHeader := c.Get("Authorization")

		// Memastikan header menggunakan format Bearer Token.
		parts := strings.Fields(authHeader)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {
			return c.Status(401).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Token autentikasi tidak ditemukan",
			})
		}

		// Membaca konfigurasi aplikasi.
		konf := config.Load()

		// Memeriksa keabsahan token JWT.
		claims, err := auth.PeriksaToken(parts[1], konf)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Token tidak valid atau sudah kedaluwarsa",
			})
		}

		// Mengambil ID pengguna dari Subject JWT.
		idPengguna, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil || idPengguna == 0 {
			return c.Status(401).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Identitas pengguna dalam token tidak valid",
			})
		}

		// Menyimpan identitas pengguna ke context Fiber.
		c.Locals("id_pengguna", uint(idPengguna))
		c.Locals("peran", claims.Peran)

		// Melanjutkan request ke handler berikutnya.
		return c.Next()
	}
}

// HanyaPeran membatasi akses berdasarkan peran pengguna.
func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Mengambil peran pengguna dari context.
		peran, ok := c.Locals("peran").(string)

		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"sukses": false,
				"pesan":  "Pengguna belum terautentikasi",
			})
		}

		// Memeriksa apakah peran pengguna diizinkan.
		for _, peranDiizinkan := range peranDiizinkan {
			if peran == peranDiizinkan {
				return c.Next()
			}
		}

		// Menolak akses jika peran tidak sesuai.
		return c.Status(403).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Anda tidak memiliki hak akses",
		})
	}
}
