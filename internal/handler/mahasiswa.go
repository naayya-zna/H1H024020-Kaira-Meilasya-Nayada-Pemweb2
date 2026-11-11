package handler

import (
	"errors"
	"strconv"
	"strings"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/dto"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
	"latihan-fiber-api/internal/service"
	"latihan-fiber-api/internal/validasi"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// GetMahasiswa menampilkan seluruh data mahasiswa.

func GetMahasiswa(c fiber.Ctx) error {
	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)
	// Mengambil nomor halaman, default halaman 1.
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Parameter page harus berupa angka positif",
		})
	}

	// Mengambil jumlah data per halaman, default 10.
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Parameter limit harus antara 1 sampai 100",
		})
	}

	// Mengambil filter angkatan.
	angkatan := 0
	if c.Query("angkatan") != "" {
		angkatan, err = strconv.Atoi(c.Query("angkatan"))
		if err != nil || angkatan < 1 {
			return c.Status(400).JSON(fiber.Map{
				"pesan": "Parameter angkatan tidak valid",
			})
		}
	}

	// Mengambil parameter pengurutan.
	urut := strings.ToLower(c.Query("urut"))
	if urut != "" && urut != "nama" && urut != "ipk" {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Parameter urut hanya boleh nama atau ipk",
		})
	}

	// Mengambil kata kunci pencarian.
	cari := strings.TrimSpace(c.Query("cari"))

	// Memanggil fungsi pagination pada repository.
	data, total, totalHalaman, err := svc.GetAll(
		page,
		limit,
		angkatan,
		urut,
		cari,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengambil data mahasiswa",
		})
	}

	// Mengembalikan hasil dalam format JSON.
	return c.JSON(fiber.Map{
		"pesan":         "Data mahasiswa berhasil diambil",
		"data":          data,
		"halaman":       page,
		"batas_data":    limit,
		"total_data":    total,
		"total_halaman": totalHalaman,
	})
}

// GetMahasiswaByID mencari mahasiswa berdasarkan ID.
func GetMahasiswaByID(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "ID harus berupa angka",
		})
	}

	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)

	mahasiswa, err := svc.GetByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"pesan": "Mahasiswa tidak ditemukan",
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengambil data mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Data mahasiswa ditemukan",
		"data":  mahasiswa,
	})
}

// GetMahasiswaByNIM mencari mahasiswa berdasarkan NIM.
func GetMahasiswaByNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)

	mahasiswa, err := svc.GetByNIM(nim)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"pesan": "Mahasiswa tidak ditemukan",
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengambil data mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Data mahasiswa ditemukan",
		"data":  mahasiswa,
	})
}

// CreateMahasiswa menambahkan data mahasiswa baru.

func CreateMahasiswa(c fiber.Ctx) error {
	var request dto.CreateMahasiswaRequest

	// Membaca data JSON dari request.
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	// Memeriksa aturan validasi pada DTO.
	galat := validasi.Periksa(request)

	// Mengembalikan HTTP 422 jika terdapat kesalahan validasi.
	if len(galat) > 0 {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat":  galat,
		})
	}

	mahasiswa := model.Mahasiswa{
		ProgramStudiID: request.ProgramStudiID,
		Nim:            request.Nim,
		Nama:           request.Nama,
		Email:          request.Email,
		Angkatan:       request.Angkatan,
		Ipk:            request.Ipk,
		Aktif:          request.Aktif,
	}

	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)

	if err := svc.Create(&mahasiswa); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menambahkan mahasiswa. Periksa NIM, email, dan program studi",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mahasiswa berhasil ditambahkan",
		"data":   mahasiswa,
	})
}

// UpdateMahasiswa memperbarui data mahasiswa.
func UpdateMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "ID harus berupa angka",
		})
	}

	var request dto.UpdateMahasiswaRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Format data tidak valid",
		})
	}

	if request.ProgramStudiID == 0 ||
		request.Nim == "" ||
		request.Nama == "" ||
		request.Email == "" ||
		request.Angkatan <= 0 ||
		request.Ipk < 0 ||
		request.Ipk > 4 {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Data mahasiswa tidak lengkap atau tidak valid",
		})
	}

	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)

	// Pastikan mahasiswa yang akan diubah tersedia.
	_, err = svc.GetByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"pesan": "Mahasiswa tidak ditemukan",
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal memeriksa data mahasiswa",
		})
	}

	// Gunakan map agar nilai false pada field Aktif tetap bisa diperbarui.
	data := map[string]interface{}{
		"program_studi_id": request.ProgramStudiID,
		"nim":              request.Nim,
		"nama":             request.Nama,
		"email":            request.Email,
		"angkatan":         request.Angkatan,
		"ipk":              request.Ipk,
		"aktif":            request.Aktif,
	}

	if err := repo.Update(uint(id), data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Gagal memperbarui data mahasiswa",
		})
	}

	mahasiswa, err := repo.GetByID(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Data berhasil diubah, tetapi gagal mengambil hasil perubahan",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Mahasiswa berhasil diperbarui",
		"data":  mahasiswa,
	})
}

// DeleteMahasiswa menghapus data mahasiswa.
func DeleteMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "ID harus berupa angka",
		})
	}

	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)
	// Pastikan data mahasiswa tersedia.
	_, err = svc.GetByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"pesan": "Mahasiswa tidak ditemukan",
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal memeriksa data mahasiswa",
		})
	}

	if err := svc.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal menghapus mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Mahasiswa berhasil dihapus",
	})
}
func RestoreMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "ID tidak valid",
		})
	}

	repo := repository.NewMahasiswaRepository(config.DB)
	svc := service.NewMahasiswaService(repo)

	mahasiswa, err := svc.GetByIDIncludingDeleted(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"pesan": "Mahasiswa tidak ditemukan",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengambil data mahasiswa",
		})
	}

	if !mahasiswa.DeletedAt.Valid {
		return c.Status(400).JSON(fiber.Map{
			"pesan": "Mahasiswa belum dihapus",
		})
	}

	if err := svc.Restore(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"pesan": "Gagal mengembalikan mahasiswa",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "Mahasiswa berhasil dikembalikan",
	})
}
