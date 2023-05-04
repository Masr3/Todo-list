package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-webapp/Controllers"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))

	//Crud Operations
	r.Get("/tasks/{id}", Controllers.GetTask)
	r.Get("/tasks", Controllers.GetTasks)
	r.Post("/tasks", Controllers.PostTask)
	r.Delete("/tasks/{id}", Controllers.DeleteTask)
	r.Put("/tasks/{id}", Controllers.PutTask)
	http.ListenAndServe("", r)

	fmt.Println("Serving at port: localhost:80")

}
