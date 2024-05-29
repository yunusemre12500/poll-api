package v1

import (
	"github.com/google/uuid"
	domain "github.com/yunusemre12500/poll-api/internal/domain/poll/v1"
)

type InMemoryPollRepository struct {
	polls []*domain.Poll
}

func NewInMemoryPollRepository() InMemoryPollRepository {
	return InMemoryPollRepository{
		polls: make([]*domain.Poll, 0),
	}
}

func (repository *InMemoryPollRepository) Create(poll *domain.Poll) error {
	repository.polls = append(repository.polls, poll)

	return nil
}

func (repository *InMemoryPollRepository) GetByID(id uuid.UUID) (*domain.Poll, error) {
	for _, poll := range repository.polls {
		if poll.ID == id {
			return poll, nil
		}
	}

	return nil, domain.ErrNotFound
}

func (repository *InMemoryPollRepository) List(limit, offset uint) ([]*domain.Poll, error) {
	var polls []*domain.Poll

	for index, poll := range repository.polls {
		if index == int(limit) {
			break
		}

		polls = append(polls, poll)
	}

	return polls, nil
}
