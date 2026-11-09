package handler

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/repository"
	"latihan-fiber-api/internal/service"

	"github.com/gofiber/fiber/v3"
)

func GetProgramStudi(c fiber.Ctx) error {
	repo := repository.NewProgramStudiRepository(config.DB)
	svc := service.NewProgramStudiService(repo)

	data, err := svc.GetAllWithJumlahMahasiswa()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengambil data program studi",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Data program studi berhasil diambil",
		"data":  data,
	})
}
