package jwthelpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
)

func GenerateToken(userID, email, role, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)

	claims := commondtos.Claims{
		UserID: userID,
		Email:  email,
		Role:   role,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidToken(t *jwt.Token, id string) bool {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	uid, ok := claims["user_id"].(string)
	if !ok {
		return false
	}

	return uid == id
}
