package jwthelpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
)

func GenerateToken(userID, role, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)

	claims := commondtos.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidToken(t *jwt.Token, id string) bool {
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["user_id"]

	return uid == id
}
