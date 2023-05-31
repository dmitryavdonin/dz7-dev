package message

import (
	"context"
	"order/internal/domain/message"

	"time"
)

func (s *service) CreateMessage(ctx context.Context, message *message.Message) (err error) {
	return s.repository.CreateMessage(ctx, message)
}

func (s *service) DeleteOldMessages(ctx context.Context, timeStamp time.Time) (err error) {
	return s.repository.DeleteOldMessages(ctx, timeStamp)
}
