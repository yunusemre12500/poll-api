package v1

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPollExists   = errors.New("poll exists")
	ErrPollNotFound = errors.New("poll not found")
	ErrNoPollsFound = errors.New("no polls found")
)

type PollRepository interface {
	Create(ctx context.Context, poll *Poll) error
	GetByID(ctx context.Context, id *uuid.UUID) (*Poll, error)
	List(ctx context.Context, limit, offset uint) ([]*Poll, error)
}
