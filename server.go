package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mceazy2700/movies-golang/graph"
	"github.com/mceazy2700/movies-golang/graph/generated"
	"github.com/mceazy2700/movies-golang/graph/utils"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  db, err := utils.InitDb()
  if err != nil {
    log.Fatal(err)
  }

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

  customContext := &utils.CustomContext{Database: db}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", utils.CreateCotext(customContext, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
