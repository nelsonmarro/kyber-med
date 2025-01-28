package commondtos

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
