package routes

import (
	"github.com/gorilla/mux"
	"github.com/prakhar-8999/todo/pkg/controllers"
)

var RegisterTodo = func(router *mux.Router) {
	router.HandleFunc("/todo/", controllers.CreateTodos).Methods("POST")
	router.HandleFunc("/todo/", controllers.GetTodos).Methods("GET")
	// router.HandleFunc("/todo/{todo_id}", controllers).Methods("GET")
	router.HandleFunc("/todo/{todo_id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{todo_id}", controllers.DeleteTodo).Methods("DELETE")
}
