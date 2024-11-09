package user

type UserType int

const (
	Visitor UserType = iota
	Standard
	Admin
	Superadmin
)

type UserStatus int

const (
	Offline UserStatus = iota
	Online
)

type User struct {
	Id     int
	Type   UserType
	Name   string
	Status UserStatus
}
