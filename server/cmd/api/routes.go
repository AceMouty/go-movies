package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *application) mapRoutes() http.Handler {
  mux := chi.NewRouter()

  // middleware...
  mux.Use(middleware.Recoverer) // recover from a panic and return a 500 internal server error

  // map routes... 
  mux.Get("/", a.handlerLiveCheck)

  return mux
}
