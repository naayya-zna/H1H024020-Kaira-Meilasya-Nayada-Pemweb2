package handler

import (
	"strconv"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/dto"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
	"latihan-fiber-api/internal/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetMataKuliah(c fiber.Ctx) error {
	repo := repository.NewMataKuliahRepository(config.DB)
	svc := service.NewMataKuliahService(repo)

	data, err := svc.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Data mata kuliah berhasil diambil",
		"data":  data,
	})
}

func GetMataKuliahByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "ID tidak valid",
		})
	}

	repo := repository.NewMataKuliahRepository(config.DB)
	svc := service.NewMataKuliahService(repo)

	data, err := svc.GetByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"pesan": "Mata kuliah tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal mengambil data"})
	}

	return c.JSON(fiber.Map{
		"pesan": "Data mata kuliah ditemukan",
		"data":  data,
	})
}

func CreateMataKuliah(c fiber.Ctx) error {
	var req dto.CreateMataKuliahRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"pesan": "Format JSON tidak valid"})
	}

	data := model.MataKuliah{
		ProgramStudiID: req.ProgramStudiID,
		Kode:           req.Kode,
		Nama:           req.Nama,
		SKS:            req.SKS,
		Semester:       req.Semester,
	}

	repo := repository.NewMataKuliahRepository(config.DB)
	svc := service.NewMataKuliahService(repo)

	if err := svc.Create(&data); err != nil {
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal menambahkan mata kuliah"})
	}

	return c.Status(201).JSON(fiber.Map{
		"pesan": "Mata kuliah berhasil ditambahkan",
		"data":  data,
	})
}

func UpdateMataKuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return c.Status(400).JSON(fiber.Map{"pesan": "ID tidak valid"})
	}

	var req dto.UpdateMataKuliahRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"pesan": "Format JSON tidak valid"})
	}

	repo := repository.NewMataKuliahRepository(config.DB)
	svc := service.NewMataKuliahService(repo)

	_, err = svc.GetByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"pesan": "Mata kuliah tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal mengambil data"})
	}

	data := map[string]interface{}{
		"program_studi_id": req.ProgramStudiID,
		"kode":             req.Kode,
		"nama":             req.Nama,
		"sks":              req.SKS,
		"semester":         req.Semester,
	}

	if err := svc.Update(uint(id), data); err != nil {
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal memperbarui mata kuliah"})
	}

	hasil, err := svc.GetByID(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal mengambil data terbaru"})
	}

	return c.JSON(fiber.Map{
		"pesan": "Mata kuliah berhasil diperbarui",
		"data":  hasil,
	})
}

func DeleteMataKuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return c.Status(400).JSON(fiber.Map{"pesan": "ID tidak valid"})
	}

	repo := repository.NewMataKuliahRepository(config.DB)
	svc := service.NewMataKuliahService(repo)

	_, err = svc.GetByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"pesan": "Mata kuliah tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal mengambil data"})
	}

	if err := svc.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"pesan": "Gagal menghapus mata kuliah"})
	}

	return c.JSON(fiber.Map{"pesan": "Mata kuliah berhasil dihapus"})
}
