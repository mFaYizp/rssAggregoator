package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mFaYizp/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	PortString := os.Getenv("PORT")
	if PortString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbUrl := os.Getenv("DE_URL")
	if dbUrl == "" {
		log.Fatal("DB environment variable is not set")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Route.Post("/users", apiCfg.handlerCreateUser)

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
