package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	PortString := os.Getenv("PORT")

	if PortString == "" {
		log.Fatal("PORT environment variable is not set")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Route := chi.NewRouter()
	v1Route.Get("/health", handlerReadiness)
	v1Route.Get("/err", handlerErr)

	router.Mount("/v1", v1Route)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PortString,
	}

	log.Printf("Server starting on %v", PortString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
