package database

import (
	"context"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Url string
}

func (i Database) GetTemplates () []bson.D  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://system:WV4fNKP2axPC5eUv@cluster0.agvxp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	if err != nil { log.Fatal(err) }
	collection := client.Database("server").Collection("test")
	cur, err := collection.Find(ctx, bson.D{})
	var results []bson.D
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		results =append(results, result)
	}
	return results
}