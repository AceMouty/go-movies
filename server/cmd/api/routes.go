package main

import (
	"github.com/go-chi/chi/v5/middleware"
)

func (a *application) mapRoutes() {

  // middleware...
  a.mux.Use(middleware.Recoverer) // recover from a panic and return a 500 internal server error
  a.mux.Use(enableCORS)

  // map routes... 
  a.mux.Get("/", a.handlerLiveCheck)
  a.mux.Get("/api/movies", a.handlerGetMovies)
}
