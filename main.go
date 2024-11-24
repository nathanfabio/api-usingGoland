package main

import (
	"encoding/json"
	"github.com/nathanfabio/api-usingGoland/internal/todo"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {

	svc := todo.NewService()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(svc)
		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		svc.Add(t.Item)
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("DELETE /", func(w http.ResponseWriter, r *http.Request) {

	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
