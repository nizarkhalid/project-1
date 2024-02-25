package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/nizarkhalid/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	// our server port
	portSting := os.Getenv("PORT")
	if portSting == "" {
		log.Fatal("port was not found in env")
	}
	// our database url connection string
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("Database was not found in env")
	}

	// we make the connection
	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("can't connnet to the Database")
	}

	// HERE
	api_confi := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/ready", handlerReadiness)
	v1router.Get("/err", handlerErr)
	v1router.Post("/user", api_confi.handlerCreateUser)
	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portSting,
	}

	fmt.Printf("the server is running in port %v", portSting)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
