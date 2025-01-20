package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
  domain string
}

func main() {
  // create application config
  app := application{ domain: "example.com" }

  // read application flags

  // connect to the db

  // startup application
  log.Println("starting application on port:", port)

  // listen and serve
  err := http.ListenAndServe(fmt.Sprintf(":%v", port), app.mapRoutes())
  if err != nil {
    log.Fatal(fmt.Errorf("unable to run server: %v", err))
  }
}
