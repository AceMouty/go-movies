package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (a *application) handlerLiveCheck(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
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

func (a *application) handlerLogin(w http.ResponseWriter, r *http.Request) {
	// read JSON page
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"pssword"`
	}

	if err := a.readJson(w, r, &loginRequest); err != nil {
		a.errorJson(w, err)
		return
	}

	// validate user against DB
	dbUser, err := a.db.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil {
		log.Println("GetUserByEmail: unable to get user from the database:", err)
		a.errorJson(w, errors.New("invalid credentials"))
		return
	}

	// validate password provided
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginRequest.Password)); err != nil {
		log.Printf("GetUserByEmail: unable to verify user: %v\n", err)
		a.errorJson(w, errors.New("invalid credentials"))
		return
	}

	// create JWT user
	u := jwtUser{ID: int(dbUser.ID), FirstName: dbUser.FirstName, LastName: dbUser.LastName}

	// gen token pair
	tokens, err := a.auth.GenerateTokenPair(&u)
	if err != nil {
		a.errorJson(w, err)
		return
	}

	refreshCookie := a.auth.GetRefreshCookie(tokens.RefreshToken)

	http.SetCookie(w, refreshCookie)
	// w.Write([]byte(tokens.Token))
	a.writeJson(w, http.StatusAccepted, tokens)
}
