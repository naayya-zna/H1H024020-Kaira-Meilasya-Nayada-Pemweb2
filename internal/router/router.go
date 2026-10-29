package router

import (
	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

// Setup mengatur endpoint API.
func Setup(app *fiber.App) {
	api := app.Group("/api")

	// Endpoint status tidak membutuhkan API Key.
	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "aktif",
			"pesan":  "API berjalan dengan baik",
		})
	})

	// Semua endpoint mahasiswa membutuhkan API Key.
	mahasiswa := api.Group("/mahasiswa", middleware.KunciApi)
	mahasiswa.Get("/", handler.GetMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.GetMahasiswaByNIM)
	mahasiswa.Get("/:id", handler.GetMahasiswaByID)
}
