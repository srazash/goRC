package message

import (
	"gorc/gorc/user"
	"time"
)

type Message struct {
	Id      int
	From    user.User
	Posted  time.Time
	Message string
}
