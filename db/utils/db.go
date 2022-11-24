package utils

import (
	"database/sql"
	"net/url"
  _ "github.com/jackc/pgx/v5/stdlib"
)
func InitDB() (*sql.DB, error) {
  dsn := url.URL{
    Scheme: "postgres",
    User: url.UserPassword("Eazy", "iprevail"),
    Host: "localhost:5432",
    Path: "go-moviesdb",
  }
  dbUrl := dsn.Query()
  dbUrl.Add("sslmode", "disable")
  dsn.RawQuery = dbUrl.Encode()

  db, err := sql.Open("pgx", dsn.String())
  if err != nil {
    return nil, err
  }
  return db, nil
}
