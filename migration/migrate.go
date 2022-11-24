package main

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
)


func main(){
  dsn := url.URL{
    Scheme: "postgres",
    User: url.UserPassword("Eazy", "iprevail"),
    Host: "localhost:5432",
    Path: "go-moviesdb",
  }

  dbUrl := dsn.Query()
  dbUrl.Add("sslmode", "disable")
  dsn.RawQuery = dbUrl.Encode()

  cmd := fmt.Sprintf("migrate -path db/migrations -database \"postgres://Eazy:iprevail@localhost:5432/go-moviesdb?sslmode=disable\" up")
  if err := exec.Command(cmd).Run(); err != nil {
    log.Fatal("Migrate error:",err)
  }
}
