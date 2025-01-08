package main

import (
	httpHandler "example/hello/internal/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	log.Println("Start proccess for Create Http Server")
	router := chi.NewRouter()
	log.Println("New Chi Router Created .Ok")
	//router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET" /*"POST", "PUT", "DELETE", "OPTIONS"*/},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	log.Println("Routes for http request start")
	router.Get("/emails", httpHandler.GetEmails)
	log.Println("Routes for http request Finish. Ok")
	log.Println("Server Http listen in 8080 port")
	http.ListenAndServe(":8080", router)
}
