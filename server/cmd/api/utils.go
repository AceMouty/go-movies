package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonResponse struct {
  Error bool `json:"error"`
  Message string `json:"message"`
  Data interface{} `json:"data,omitempty"`
}

func (a *application) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
  respData, err := json.Marshal(data)
  if err != nil {
    return err
  }

  if len(headers) > 0 {
    for key, val:= range headers[0]{
      w.Header()[key] = val
    }
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  if _, err = w.Write(respData); err != nil {
    return err
  }

  return nil
}

func (a *application) readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
  maxBytes := 1024 * 1024 // 1MB JSON page size

  r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

  decoder := json.NewDecoder(r.Body)
  decoder.DisallowUnknownFields() // dont map to data we dont expect

  err := decoder.Decode(data)
  if err != nil {
    return err
  }

  // check for more than one JSON value
  err = decoder.Decode(&struct{}{})
  if err != io.EOF {
    return errors.New("body can only be one JSON value")
  }

  return nil
}

func (a *application) errorJson(w http.ResponseWriter, err error, status ...int) error {
  statusCode := http.StatusBadRequest
  if len(status) > 0 {
    statusCode = status[0]
  }

  var resp JsonResponse
  resp.Error = true
  resp.Message = err.Error()
  
  return a.writeJson(w, statusCode, resp)
}
