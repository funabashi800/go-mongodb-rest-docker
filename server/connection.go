package server

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctx.Done()
	clientOptions := options.Client().ApplyURI("mongodb:mongo")
	client, _ := mongo.Connect(ctx, clientOptions)
	database := client.Database("todoapp")
	return database
}