package dtos

type UserLoginDTO struct {
	Email    string `json:"email"`
	IDCard   string `json:"idCard"`
	Password string `json:"password"`
}
