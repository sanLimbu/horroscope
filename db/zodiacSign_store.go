package db

import (
	"context"
	"os"

	"github.com/sanLimbu/horroscope/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ZodiacSignStore interface {
	InsertZodiac(context.Context, *types.ZodiaSign, error)
	Update(context.Context, Map, Map) error
	GetZodiacs(context.Context, Map, *Pagination) ([]*types.ZodiaSign, error)
	GetZodiacByID(context.Context, string) (*types.ZodiaSign, error)
}

type MongoZodiacStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoZodiacStore(client *mongo.Client) *MongoZodiacStore {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoZodiacStore{
		client: client,
		coll:   client.Database(dbname).Collection("zodiac"),
	}
}

func (s *MongoZodiacStore) GetZodiacByID(ctx context.Context, id string) (*types.ZodiaSign, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var zodiac types.ZodiaSign
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&zodiac); err != nil {
		return nil, err
	}

	return &zodiac, nil

}
func (s *MongoZodiacStore) GetZodiacs(ctx context.Context, filter Map, page *Pagination) ([]*types.ZodiaSign, error) {
	opts := options.FindOptions{}
	opts.SetSkip((page.Page - 1) * page.Limit)
	opts.SetLimit(page.Limit)
	resp, err := s.coll.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	var zodiacs []*types.ZodiaSign
	if err := resp.All(ctx, &zodiacs); err != nil {
		return nil, err
	}

	return zodiacs, nil
}

func (s *MongoZodiacStore) Update(ctx context.Context, filter Map, update Map) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoZodiacStore) InsertZodiac(ctx context.Context, zodiac *types.ZodiaSign) (*types.ZodiaSign, error) {
	resp, err := s.coll.InsertOne(ctx, zodiac)
	if err != nil {
		return nil, err
	}

	zodiac.ID = resp.InsertedID.(primitive.ObjectID)
	return zodiac, nil
}
