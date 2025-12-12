package repository

import (
	"context"
	"time"
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDbBrandRepository struct {
	client *mongo.Client
}

func NewMongoDbBrandRepository(client *mongo.Client) brand.Repository {
	return &MongoDbBrandRepository{
		client: client,
	}
}

func (m *MongoDbBrandRepository) FindTop10() ([]model.Brand, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoDbBrandRepository) FindAll() ([]model.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.client.Database("Cluster0").Collection("brands")

	projection := bson.D{
		{"brand_name", 1},
		{"devices", 0},
	}

	opts := options.Find().SetProjection(projection)

	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var brands []model.Brand
	if err := cursor.All(ctx, &brands); err != nil {
		return nil, err
	}

	return brands, nil
}
