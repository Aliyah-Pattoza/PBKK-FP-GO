package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"` // Tambahkan role
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key") // Kunci rahasia JWT

func GenerateJWT(email string, role string) (string, error) {
	claims := &Claims{
		Email: email,
		Role:  role, // Role ditambahkan di sini
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // Token berlaku selama 2 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT validates the JWT token and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	// Hapus prefix "Bearer " dari token jika ada
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validasi metode tanda tangan
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Ambil klaim dari token
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	// Validasi waktu kedaluwarsa
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
