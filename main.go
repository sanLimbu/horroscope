package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sanLimbu/hotel-reservation/api"
	"github.com/sanLimbu/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi          = "mongodb://localhost:27017"
	dbname         = "hotel-reservation"
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

	fmt.Println(res)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API Server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleUser)
	app.Listen(*listenAddr)

}
