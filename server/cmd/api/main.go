package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/acemouty/go-movie/internal/database"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const port = 8080

type application struct {
  mux *chi.Mux
  db *database.Queries
  domain string
  auth Auth
  JwtSecret string
  JwtIssueer string
  JwtAudience string
  CookieDomain string
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

  flag.StringVar(&app.JwtSecret, "jwt-secret" ,"keep-it-secret-keep-it-safe", "signing secret")
  flag.StringVar(&app.JwtIssueer, "jwt-issuer" ,app.domain, "signing issuer")
  flag.StringVar(&app.JwtAudience, "jwt-audience" ,app.domain, "signing audience")
  flag.StringVar(&app.CookieDomain, "cookie-domain" ,"localhost", "cookie domain")
  flag.Parse()

  // compose app config

  // connect to the db
  db := dbConnect(dbConn)
  dbQueries := database.New(db)
  app.db = dbQueries

  app.auth = Auth {
    Issuer: app.JwtIssueer,
    Audiance: app.JwtAudience,
    Secret: app.JwtSecret,
    TokenExpiry: time.Minute * 15, // 15 minute life (access token)
    RefreshExpiry: time.Hour * 24, // 1 day life (refresh token)
    CookiePath: "/",
    CookieName: "__Host-refresh_token",
    CookieDomain: app.CookieDomain,
  }

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
