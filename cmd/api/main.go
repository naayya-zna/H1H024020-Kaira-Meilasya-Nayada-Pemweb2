package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/middleware"
	"latihan-fiber-api/internal/router"
)

func main() {
	cfg := config.Load()

	// Hubungkan ke database MySQL.
	config.HubungkanBasisData(cfg)

	// Buat tabel berdasarkan model GORM.
	config.Migrasi()

	// Masukkan data awal program studi.
	config.DataAwal()

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,

		// Menangani error secara terstruktur.
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			pesan := "Terjadi kesalahan pada server"

			// Periksa apakah error berasal dari Fiber.
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code

				// Gunakan pesan error untuk status selain 500.
				if code < 500 {
					pesan = e.Message
				}
			}

			// Catat detail error di terminal server.
			log.Printf("Error [%d]: %v", code, err)

			return c.Status(code).JSON(fiber.Map{
				"sukses": false,
				"pesan":  pesan,
			})
		},
	})

	// Middleware bawaan Fiber.
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// Middleware buatan sendiri.
	app.Use(middleware.WaktuProses)

	// Daftarkan endpoint API.
	router.Setup(app)

	// Respons untuk endpoint yang tidak ditemukan.
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Endpoint tidak ditemukan",
		})
	})

	log.Printf("%s berjalan pada port %s", cfg.AppName, cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
