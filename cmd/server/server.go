package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nironwp/graphql/graph"
	"github.com/nironwp/graphql/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	db, err := sql.Open("sqlite3", "db.sqlite")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if port == "" {
		port = defaultPort
	}
	CategoryDB := database.NewCategory(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: CategoryDB,
		CourseDB:   database.NewCourse(db),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
