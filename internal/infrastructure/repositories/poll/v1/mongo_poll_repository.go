package v1

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"
	domain "github.com/yunusemre12500/poll-api/internal/domain/poll/v1"
)

type MongoPollRepository struct {
	collection mongo.Collection
}

func NewMongoPollRepository(collection mongo.Collection) MongoPollRepository {
	return MongoPollRepository{
		collection: collection,
	}
}

func (repository *MongoPollRepository) Create(poll *domain.Poll) error {
	if _, err := repository.collection.InsertOne(context.TODO(), poll, nil); err != nil {
		return err
	}

	return nil
}

func (repository *MongoPollRepository) GetByID(id uuid.UUID) (*domain.Poll, error) {
	var poll domain.Poll

	if err := repository.collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&poll); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &poll, nil
}

func (repository *MongoPollRepository) List(limit, offset uint) ([]*domain.Poll, error) {
	var polls []*domain.Poll

	options := options.Find().SetSkip(int64(limit) * int64(offset)).SetLimit(int64(limit))

	cursor, err := repository.collection.Find(context.TODO(), bson.D{}, options)

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &polls)

	if err != nil {
		return nil, err
	}

	return polls, nil
}
