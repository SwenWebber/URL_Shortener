package mongodb

import (
	"context"
	"time"
	"urlShortener/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, dbName string) *mongoRepository {
	collection := client.Database(dbName).Collection("urls")

	return &mongoRepository{
		client:     client,
		collection: collection,
	}
}

func (r *mongoRepository) SaveURL(url *models.URL) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := r.collection.InsertOne(ctx, url)
	return err
}

func (r *mongoRepository) FindByShortCode(code string) (*models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var url models.URL

	err := r.collection.FindOne(ctx, bson.M{"short_code": code}).Decode(&url)

	if err != nil {
		return nil, err
	}

	return &url, nil
}
