package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mceazy2700/movies-golang/db/utils"
	"github.com/mceazy2700/movies-golang/graph/generated"
	"github.com/mceazy2700/movies-golang/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  db, err := utils.InitDB()
  if err != nil {
    log.Fatal("Database Initialization Error:", err)
  }

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

  customContext := &utils.CustomContext{
    Database: db,
  }

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", utils.CreateContext(customContext, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
