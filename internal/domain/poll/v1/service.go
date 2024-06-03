package v1

import (
	"context"

	"github.com/google/uuid"
)

type PollService interface {
	Create(ctx context.Context, poll *Poll) error
	GetByID(ctx context.Context, id *uuid.UUID) (*Poll, error)
	List(ctx context.Context, limit, offset uint) ([]*Poll, error)
}
