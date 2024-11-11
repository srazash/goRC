package message

import (
	"fmt"
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

func New(from user.User, message string) *Message {
	return &Message{
		Id:      *guid.New(),
		From:    from,
		Posted:  time.Now(),
		Message: message,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("%s: %s [%s]", m.From.Name, m.Message, m.Posted.Format("15:04:05"))
}
