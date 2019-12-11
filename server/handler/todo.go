package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	"encoding/json"
	"net/http"
	"context"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/funabashi800/todoapp/server/model"
)

type Handler struct {
	Response http.ResponseWriter
	Request *http.Request
}

func (h *Handler) CreateTodo(collection *mongo.Collection) model.Todo {
	var todo model.Todo
	_ = json.NewDecoder(h.Request.Body).Decode(&todo)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		log.Fatal(err)
	}
	return todo
}

// GetTodo function Done
func (h *Handler) GetTodo(collection *mongo.Collection) model.Todo {
	h.Response.Header().Set("content-type", "application/json")
	ids, _ := h.Request.URL.Query()["id"]
	id := ids[0]
	var todo model.Todo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	document := collection.FindOne(ctx, model.Todo{ID: id})
	err := document.Decode(&todo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return todo
}

func (h *Handler) GetAllTodo(collection *mongo.Collection) []*model.Todo {
	h.Response.Header().Set("content-type", "application/json")
	var todos []*model.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var todo model.Todo
		cursor.Decode(&todo)
		todos = append(todos, &todo)
	}
	return todos
}