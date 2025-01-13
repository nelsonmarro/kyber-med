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
			Issuer:    "nutricional_app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString, jwtKey string) (bool, error) {
	claims := &commondtos.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
