package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/acemouty/go-movie/internal/database"
	_ "github.com/lib/pq"
)

const port = 8080

type application struct {
  mux *chi.Mux
  db *database.Queries
  domain string
}

func main() {
  // create application config
  app := application{ 
    domain: "example.com", 
    mux: chi.NewMux(),
  }

  // read application flags

  // take in db conn string
  var dbConn string
  flag.StringVar(&dbConn, "dbConn", "postgres://postgres:postgres@localhost:5432/gomovies?sslmode=disable", "Postgres Conection String")
  db := dbConnect(dbConn)
  dbQueries := database.New(db)
  app.db = dbQueries

  // connect to the db

  // startup application
  log.Println("starting application on port:", port)
  app.mapRoutes()

  // listen and serve
  err := http.ListenAndServe(fmt.Sprintf(":%v", port), app.mux)
  if err != nil {
    log.Fatal(fmt.Errorf("unable to run server: %v", err))
  }
}

func dbConnect(dbConn string) *sql.DB {
  db, err := sql.Open("postgres", dbConn)
  if err != nil {
    log.Fatalf("main: Issue connecting to database: %v", err)
  }

  return db
}
