package message

import (
	"context"
	"order/internal/domain/message"
	"time"
)

type Service interface {
	CreateMessage(ctx context.Context, message *message.Message) (err error)
	DeleteOldMessages(ctx context.Context, timeStamp time.Time) (err error)
}
