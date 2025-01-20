package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *application) handlerLiveCheck(w http.ResponseWriter, r *http.Request) {
  resp := struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Version string `json:"version"`
  } {
    Status: "active",
    Message: "Server up and running",
    Version: "v1",
  }

  data, err := json.Marshal(resp)
  err = fmt.Errorf("Generic err")
  if err != nil {
    fmt.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(data)
}

func (a *application) handlerGetMovies(w http.ResponseWriter, r *http.Request) {
  movies, err := a.db.GetMovies(context.Background())
  if err != nil {
    a.errorJson(w, err)
    return
  }

  data, err := json.Marshal(movies)
  if err != nil {
    fmt.Println(err)
    return
  }

  err = a.writeJson(w, http.StatusOK, data)
  if err != nil {
    fmt.Println(err)
    return
  }
}
