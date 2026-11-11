package router

import (
	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

// Setup mengatur seluruh endpoint API.
func Setup(app *fiber.App) {
	api := app.Group("/api")

	// =====================================
	// ENDPOINT STATUS
	// =====================================

	// Tidak membutuhkan API Key.
	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "aktif",
			"pesan":  "API berjalan dengan baik",
		})
	})

	// =====================================
	// ENDPOINT AUTENTIKASI
	// =====================================

	// Route autentikasi.
	auth := api.Group("/auth")
	auth.Post("/registrasi", handler.Registrasi)
	auth.Post("/login", handler.Login)
	auth.Get("/profil", middleware.Autentikasi(), handler.Profil)
	auth.Post("/segarkan", middleware.Autentikasi(), handler.Segarkan)
	auth.Post("/logout", middleware.Autentikasi(), handler.Logout)
	// =====================================
	// ENDPOINT MAHASISWA
	// =====================================

	// Semua endpoint mahasiswa membutuhkan API Key.
	mahasiswa := api.Group("/mahasiswa", middleware.KunciApi)

	mahasiswa.Get("/", handler.GetMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.GetMahasiswaByNIM)
	mahasiswa.Get("/:id", handler.GetMahasiswaByID)
	mahasiswa.Post("/", handler.CreateMahasiswa)
	mahasiswa.Put("/:id", handler.UpdateMahasiswa)
	mahasiswa.Delete("/:id", handler.DeleteMahasiswa)
	mahasiswa.Post("/:id/restore", handler.RestoreMahasiswa)

	// =====================================
	// ENDPOINT MATA KULIAH
	// =====================================

	matakuliah := api.Group("/matakuliah", middleware.KunciApi)

	matakuliah.Get("/", handler.GetMataKuliah)
	matakuliah.Get("/:id", handler.GetMataKuliahByID)
	matakuliah.Post("/", handler.CreateMataKuliah)
	matakuliah.Put("/:id", handler.UpdateMataKuliah)
	matakuliah.Delete("/:id", handler.DeleteMataKuliah)

	// =====================================
	// ENDPOINT PROGRAM STUDI
	// =====================================

	programStudi := api.Group("/program-studi", middleware.KunciApi)

	programStudi.Get("/", handler.GetProgramStudi)
}
