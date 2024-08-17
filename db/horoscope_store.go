package db

import (
	"context"
	"os"

	"github.com/sanLimbu/horroscope/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HoroscopeStore interface {
	InsertHoroscope(context.Context, *types.Horoscope, error)
	GetHoroscopes(context.Context, bson.M) ([]*types.Horoscope, error)
}

type MongoHoroscopeStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	ZodiacSignStore
}

func NewMongoHoroscopeStore(client *mongo.Client, zodiacStore ZodiacSignStore) *MongoHoroscopeStore {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoHoroscopeStore{
		client:          client,
		coll:            client.Database(dbname).Collection("horoscopes"),
		ZodiacSignStore: zodiacStore,
	}
}

func (s *MongoHoroscopeStore) GetHoroscopes(ctx context.Context, filter bson.M) ([]*types.Horoscope, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var horoscopes []*types.Horoscope
	if err := resp.All(ctx, &horoscopes); err != nil {
		return nil, err
	}

	return horoscopes, nil

}

func (s *MongoHoroscopeStore) InsertHoroscope(ctx context.Context, horoscope *types.Horoscope) (*types.Horoscope, error) {
	resp, err := s.coll.InsertOne(ctx, horoscope)
	if err != nil {
		return nil, err
	}

	horoscope.ID = resp.InsertedID.(primitive.ObjectID)

	//update zodiac with this horoscope id
	filter := Map{"_id": horoscope.ZodiacSignID}
	update := Map{"$push": bson.M{"horoscope": horoscope.ID}}

	if err := s.ZodiacSignStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}
	return horoscope, nil
}
