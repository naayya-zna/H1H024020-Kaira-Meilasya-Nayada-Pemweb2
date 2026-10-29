package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

// KunciApi memeriksa API Key pada setiap request.
func KunciApi(c fiber.Ctx) error {
	kunciBenar := os.Getenv("API_KEY")
	kunciRequest := c.Get("X-Api-Key")

	// Tolak request jika API Key belum dikonfigurasi
	// atau kunci yang dikirim tidak sesuai.
	if kunciBenar == "" || kunciRequest != kunciBenar {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"pesan": "API Key tidak valid atau tidak tersedia",
		})
	}

	// Lanjutkan request jika kunci benar.
	return c.Next()
}
