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

	// READ: Mengambil seluruh data mahasiswa.
	mahasiswa.Get("/", handler.GetMahasiswa)

	// READ: Mengambil mahasiswa berdasarkan NIM.
	mahasiswa.Get("/nim/:nim", handler.GetMahasiswaByNIM)

	// READ: Mengambil mahasiswa berdasarkan ID.
	mahasiswa.Get("/:id", handler.GetMahasiswaByID)

	// CREATE: Menambahkan mahasiswa baru.
	mahasiswa.Post("/", handler.CreateMahasiswa)

	// UPDATE: Memperbarui data mahasiswa.
	mahasiswa.Put("/:id", handler.UpdateMahasiswa)

	// DELETE: Menghapus data mahasiswa.
	mahasiswa.Delete("/:id", handler.DeleteMahasiswa)
	mahasiswa.Post("/:id/restore", handler.RestoreMahasiswa)

	matakuliah := api.Group("/matakuliah", middleware.KunciApi)
	matakuliah.Get("/", handler.GetMataKuliah)
	matakuliah.Get("/:id", handler.GetMataKuliahByID)
	matakuliah.Post("/", handler.CreateMataKuliah)
	matakuliah.Put("/:id", handler.UpdateMataKuliah)
	matakuliah.Delete("/:id", handler.DeleteMataKuliah)

	programStudi := api.Group("/program-studi", middleware.KunciApi)
	programStudi.Get("/", handler.GetProgramStudi)
}
