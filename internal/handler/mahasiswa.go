package handler

import (
	"sort"
	"strconv"
	"strings"

	"latihan-fiber-api/internal/model"

	"github.com/gofiber/fiber/v3"
)

// GetMahasiswa menampilkan daftar mahasiswa.
func GetMahasiswa(c fiber.Ctx) error {
	data := make([]model.Mahasiswa, 0, len(model.DataMahasiswa))
	data = append(data, model.DataMahasiswa...)

	// Filter berdasarkan angkatan.
	angkatan := c.Query("angkatan")
	if angkatan != "" {
		tahun, err := strconv.Atoi(angkatan)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"pesan": "Angkatan harus berupa angka",
			})
		}

		hasil := make([]model.Mahasiswa, 0)
		for _, m := range data {
			if m.Angkatan == tahun {
				hasil = append(hasil, m)
			}
		}
		data = hasil
	}

	// Urutkan berdasarkan nama atau IPK.
	urut := strings.ToLower(c.Query("urut"))

	switch urut {
	case "nama":
		sort.Slice(data, func(i, j int) bool {
			return data[i].Nama < data[j].Nama
		})
	case "ipk":
		sort.Slice(data, func(i, j int) bool {
			return data[i].IPK > data[j].IPK
		})
	case "":
		// Tanpa parameter urut, gunakan urutan data awal.
	default:
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Parameter urut hanya boleh nama atau ipk",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Data mahasiswa berhasil diambil",
		"data":  data,
	})
}

// GetMahasiswaByID mencari mahasiswa berdasarkan ID.
func GetMahasiswaByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "ID harus berupa angka",
		})
	}

	for _, m := range model.DataMahasiswa {
		if m.ID == id {
			return c.JSON(fiber.Map{
				"pesan": "Data mahasiswa ditemukan",
				"data":  m,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"pesan": "Mahasiswa tidak ditemukan",
	})
}

// GetMahasiswaByNIM mencari mahasiswa berdasarkan NIM.
func GetMahasiswaByNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	for _, m := range model.DataMahasiswa {
		if strings.EqualFold(m.NIM, nim) {
			return c.JSON(fiber.Map{
				"pesan": "Data mahasiswa ditemukan",
				"data":  m,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"pesan": "Mahasiswa tidak ditemukan",
	})
}
