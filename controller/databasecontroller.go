package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Database struct {
	Url string
}

type Response struct {
	correct string
	weight float32
}

type Alternative struct {
	string
}

type Template struct {
	Name         string
	alternatives []Alternative
	Responses    []Response
}

func (i Database) InsertData () bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		i.Url,
	))
	if err != nil { log.Fatal(err) 
		return false }
	collection := client.Database("server").Collection("templates")
	cursor, err := collection.InsertOne(ctx, bson.M{})
	_ = cursor
	if err != nil { return false }
	return true
}

func (i Database) GetTemplates () []bson.M  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		i.Url,
	))
	if err != nil { log.Fatal(err) }
	collection := client.Database("server").Collection("templates")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func (i Database) GetTemplate (id int) bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		i.Url,
	))
	if err != nil { log.Fatal(err) }
	collection := client.Database("server").Collection("templates")
	var result bson.M
	if err = collection.FindOne(ctx, bson.M{"id":id}).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result
}