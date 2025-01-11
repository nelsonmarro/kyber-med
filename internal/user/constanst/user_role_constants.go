package constanst

type UserRole string

const (
	RoleUser   UserRole = "user"
	RoleAdmin  UserRole = "admin"
	RoleDoctor UserRole = "doctor"
)
