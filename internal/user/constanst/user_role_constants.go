package user

type UserRole string

const (
	RoleUser   UserRole = "user"
	RoleAdmin  UserRole = "admin"
	RoleDoctor UserRole = "doctor"
)
