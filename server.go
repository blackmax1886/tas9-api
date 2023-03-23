package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/blackmax1886/tas9-api/db"
	"github.com/blackmax1886/tas9-api/graph"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}
}

func main() {
	// loadEnv()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	// if allowedOrigin == "" {
	// 	allowedOrigin = "http://localhost:3000"
	// }

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{allowedOrigin},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
