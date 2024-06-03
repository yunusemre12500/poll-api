package v1

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/yunusemre12500/poll-api/internal/domain/poll/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPollRepository struct {
	collection mongo.Collection
}

func NewMongoPollRepository(collection *mongo.Collection) *MongoPollRepository {
	return &MongoPollRepository{
		collection: *collection,
	}
}

func (repository *MongoPollRepository) Create(ctx context.Context, poll *domain.Poll) error {
	_, err := repository.collection.InsertOne(ctx, &poll, nil)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.ErrPollExists
		}

		return err
	}

	return nil
}

func (repository *MongoPollRepository) GetByID(ctx context.Context, id *uuid.UUID) (*domain.Poll, error) {
	var poll *domain.Poll

	filter := bson.D{{Key: "_id", Value: id}}

	result := repository.collection.FindOne(ctx, filter, nil)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrPollNotFound
		}

		return nil, err
	}

	if err := result.Decode(&poll); err != nil {
		return nil, err
	}

	return poll, nil
}

func (repository *MongoPollRepository) List(ctx context.Context, limit, offset uint) ([]*domain.Poll, error) {
	var polls []*domain.Poll

	opts := options.Find().
		SetSkip(int64(limit) * int64(offset)).
		SetLimit(int64(limit))

	cursor, err := repository.collection.Find(ctx, bson.D{}, opts)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNoPollsFound
		}

		return nil, err
	}

	if err := cursor.All(ctx, &polls); err != nil {
		return nil, err
	}

	return polls, nil
}
