package message

import (
	"gorc/gorc/user"
	"time"

	"github.com/beevik/guid"
)

type Message struct {
	Id      guid.Guid
	From    user.User
	Posted  time.Time
	Message string
}
