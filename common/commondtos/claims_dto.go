package commondtos

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	IDCard string `json:"id_card"`
	jwt.RegisteredClaims
}
