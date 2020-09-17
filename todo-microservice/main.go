package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo/todo-microservice/controllers"
)

func main() {
	router := mux.NewRouter()

	// routes function is managing all the routes
	routes(router)

	fmt.Println("Application Starting at PORT: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func routes(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/", controllers.CreateTodo).Methods("POST")
}
