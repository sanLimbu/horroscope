package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sanLimbu/horroscope/api"
	"github.com/sanLimbu/horroscope/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi          = "mongodb://localhost:27017"
	dbname         = "horroscope"
	userCollection = "users"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userCollection)

	user := types.User{
		FirstName: "santos",
		LastName:  "lim",
	}

	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result", res)

	var u types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&u); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user find", u)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API Server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleUser)
	app.Listen(*listenAddr)

}
