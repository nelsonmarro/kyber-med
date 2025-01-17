package dtos

type UserLoginDTO struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
