package users

// User role names
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Role string
type Roles []Role
