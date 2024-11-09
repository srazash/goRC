package user

import (
	"github.com/beevik/guid"
)

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
	Id     guid.Guid
	Type   UserType
	Name   string
	Status UserStatus
}
