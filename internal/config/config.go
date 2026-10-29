
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
