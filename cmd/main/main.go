package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prakhar-8999/todo/pkg/routes"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTodo(r)
	http.Handle("/", r)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
