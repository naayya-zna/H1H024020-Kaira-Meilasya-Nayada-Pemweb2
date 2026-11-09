package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config menyimpan konfigurasi aplikasi.
type Config struct {
	AppName string
	AppPort string
	AppEnv  string

	// Konfigurasi database MySQL
	DbHost  string
	DbPort  string
	DbNama  string
	DbUser  string
	DbSandi string
}

// Load membaca konfigurasi dari file .env.
func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment yang tersedia")
	}

	return Config{
		AppName: getEnv("APP_NAME", "Pemweb2 Fiber API"),
		AppPort: getEnv("APP_PORT", "3000"),
		AppEnv:  getEnv("APP_ENV", "local"),

		// Membaca konfigurasi database MySQL
		DbHost:  getEnv("DB_HOST", "127.0.0.1"),
		DbPort:  getEnv("DB_PORT", "3306"),
		DbNama:  getEnv("DB_NAME", "pemweb2_fiber"),
		DbUser:  getEnv("DB_USER", "root"),
		DbSandi: getEnv("DB_PASS", ""),
	}
}

// getEnv mengambil nilai environment atau nilai default.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
