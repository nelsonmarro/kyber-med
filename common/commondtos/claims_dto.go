package commondtos

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
