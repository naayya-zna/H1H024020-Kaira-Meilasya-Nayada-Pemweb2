
package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

// WaktuProses menghitung durasi pemrosesan request.
func WaktuProses(c fiber.Ctx) error {
	mulai := time.Now()

	err := c.Next()

	durasi := time.Since(mulai)
	log.Printf("%s %s - %v", c.Method(), c.Path(), durasi)

	return err
}
