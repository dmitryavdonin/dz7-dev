package message

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	id        uuid.UUID
	createdAt time.Time
}

func NewMessage(msg_id uuid.UUID) (*Message, error) {

	return &Message{
		id:        msg_id,
		createdAt: time.Now(),
	}, nil
}

func NewMessageWithId(
	id uuid.UUID,
	createdAt time.Time,
) *Message {
	return &Message{
		id:        id,
		createdAt: createdAt,
	}
}

func (u Message) Id() uuid.UUID {
	return u.id
}

func (u Message) CreatedAt() time.Time {
	return u.createdAt
}
