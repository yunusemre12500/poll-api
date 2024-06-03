package v1

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/yunusemre12500/poll-api/internal/domain/poll/v1"
)

type HTTPPollService struct {
	repository domain.PollRepository
}

func NewHTTPPollService(repository domain.PollRepository) *HTTPPollService {
	return &HTTPPollService{
		repository: repository,
	}
}

func (service *HTTPPollService) Create(ctx context.Context, poll *domain.Poll) error {
	return service.repository.Create(ctx, poll)
}

func (service *HTTPPollService) GetByID(ctx context.Context, id *uuid.UUID) (*domain.Poll, error) {
	return service.repository.GetByID(ctx, id)
}

func (service *HTTPPollService) List(ctx context.Context, limit, offset uint) ([]*domain.Poll, error) {
	return service.repository.List(ctx, limit, offset)
}
