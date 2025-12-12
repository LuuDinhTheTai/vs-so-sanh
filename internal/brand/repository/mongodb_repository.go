package repository

import (
	"context"
	"time"
	"vs-so-sanh/infrastructures/configuration"
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDbBrandRepository struct {
	cfg    configuration.Config
	client *mongo.Client
}

func NewMongoDbBrandRepository(cfg configuration.Config, client *mongo.Client) brand.Repository {
	return &MongoDbBrandRepository{
		cfg:    cfg,
		client: client,
	}
}

func (m *MongoDbBrandRepository) FindTop20() ([]model.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.client.Database(m.cfg.Database.DBName).Collection(m.cfg.Database.CollectionName)

	opts := options.Find().
		SetLimit(20)

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

func (m *MongoDbBrandRepository) FindAll() ([]model.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.client.Database(m.cfg.Database.DBName).Collection(m.cfg.Database.CollectionName)

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
