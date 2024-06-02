package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ArrayOfLilly/blog-aggregator/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	filepathRoot := "."

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  }
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	dbURL := os.Getenv("CONNECTION")
	if dbURL == "" {
		log.Fatal("CONNECTION environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Couldn't establish the database connection")
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", handlerReady)
	mux.HandleFunc("GET /v1/err", handlerErr)

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUserCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handlerFeedsGet)

	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsDelete))


	svr := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	fmt.Printf("Server is started. Serving files on %s and listening on %s port\n", filepathRoot, port)
	log.Fatal(svr.ListenAndServe())
}