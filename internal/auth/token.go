package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

// Klaim menyimpan informasi yang terdapat di dalam JWT.
type Klaim struct {
	Peran string `json:"peran"`
	jwt.RegisteredClaims
}

// TerbitkanToken membuat JWT untuk pengguna yang berhasil login.
func TerbitkanToken(pengguna model.Pengguna, konf config.Config) (string, error) {
	waktuSekarang := time.Now()

	claims := Klaim{
		Peran: pengguna.Peran,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:  strconv.FormatUint(uint64(pengguna.ID), 10),
			Issuer:   "pemweb2-fiber-api",
			IssuedAt: jwt.NewNumericDate(waktuSekarang),
			ExpiresAt: jwt.NewNumericDate(waktuSekarang.Add(
				time.Duration(konf.JwtTtlMenit) * time.Minute,
			)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(konf.JwtSecret))
}

// PeriksaToken memeriksa keaslian dan masa berlaku JWT.
func PeriksaToken(tokenString string, konf config.Config) (*Klaim, error) {
	claims := &Klaim{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("metode tanda tangan JWT tidak valid")
			}

			return []byte(konf.JwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}

// HashToken mengubah token JWT menjadi hash SHA-256.
func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}
