package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)
var books = map[string]Book{}
const PORT = "8080" 

func main() {
	router := chi.NewRouter()
	fmt.Printf("server stating on port %v\n", PORT)


	router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"}, 
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, 
    }))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handleReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router) 
	v1Router.Get("/books", handleGetBooks)
v1Router.Post("/books", handleCreateBook)
v1Router.Get("/books/{id}", handleGetBook)
v1Router.Put("/books/{id}", handleUpdateBook)
v1Router.Delete("/books/{id}", handleDeleteBook)






	srv:=&http.Server{Handler:router,Addr: ":"+PORT}


	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
	
}
