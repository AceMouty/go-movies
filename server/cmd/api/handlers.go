package main

import (
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
  if err != nil {
    fmt.Println(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(data)
}
