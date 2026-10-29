
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

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"pesan": err.Error(),
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
		return c.Status(404).JSON(fiber.Map{
			"pesan": "Endpoint tidak ditemukan",
		})
	})

	log.Printf("%s berjalan pada port %s", cfg.AppName, cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
