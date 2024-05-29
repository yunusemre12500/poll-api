package v1

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("not found")
)

type PollRepository interface {
	Create(poll *Poll) error
	GetByID(id uuid.UUID) (*Poll, error)
	List(limit, offset uint) ([]*Poll, error)
}
