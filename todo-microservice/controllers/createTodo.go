package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/todo-microservice/model"
	"todo/todo-microservice/utils"
)

// CreateTodo Create todos and responds with the status
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// title, _ := r.Body.Read([]byte("title"))
	var body input
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println(err)
	}

	if err := body.validate(); err != nil {
		w.WriteHeader(err.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	todo := model.Todo{
		Title:       body.Title,
		IsCompleted: false,
	}

	_ = json.NewEncoder(w).Encode(todo)

}

// Request body structure for the current Route
type input struct {
	Title string `json:"title"`
}

// validate will return errors if request body validation fails
func (r input) validate() *utils.ApiErrors {
	if r.Title == "" {
		return &utils.ApiErrors{
			Errors:     []utils.FieldError{{Field: "title", Error: "title is required field"}},
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
		}
	}

	return nil
}
