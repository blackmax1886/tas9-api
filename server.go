package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/blackmax1886/tas9-api/graph"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	dsn := "localuser:localpass@tcp(127.0.0.1:3306)/localdb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Raw("show tables").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err)
		}
		fmt.Println(table)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
