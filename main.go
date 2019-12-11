package main

import (
	"log"
	"net/http"
	"github.com/funabashi800/todoapp/server/handler"
	"github.com/funabashi800/todoapp/server"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	h := handler.Handler{Response: w, Request: r}
	collection := db.Collection("todo")
	switch r.Method {
		case "GET": {
			h.GetTodo(collection)
		}
		case "POST": {
			h.CreateTodo(collection)
		}
	}
}

func main()  {
	const port string = "3000"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.
		Fatal(err)
	}
	log.Printf("Golagn server is listening on Port: %s", port)
	db := server.NewConnection()
	http.HandleFunc("/todos", handler.TodoHandler(database))
}