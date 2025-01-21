package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
  Issuer string
  Audiance string
  Secret string
  TokenExpiry time.Duration
  RefreshExpiry time.Duration
  CookieDomain string
  CookiePath string
  CookieName string
}

type jwtUser struct {
  ID int `json:"id"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

type TokenPairs struct {
  Token string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
}

type Claims struct {
  jwt.RegisteredClaims 
}

// Genreate JWT and Refresh Tokens
func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
  // Create JWT Token 
  token := jwt.New(jwt.SigningMethodHS256)

  // Set the claims of the token
  claims := token.Claims.(jwt.MapClaims)
  claims["name"] = fmt.Sprintf("%v %v", user.FirstName, user.LastName)
  // JWT spec based claims 
  claims["sub"] = fmt.Sprint(user.ID)
  claims["aud"] = j.Audiance
  claims["iss"] = j.Issuer
  claims["iat"] = time.Now().UTC().Unix() // required format for go-jwt
  claims["typ"] = "JWT"
  claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

  // Create a signed token
  signedAccesstoken, err := token.SignedString([]byte(j.Secret))
  if err != nil {
    return TokenPairs{}, fmt.Errorf("unable to sign access token: %v", err)
  }

  // Create refresh Token and set claims
  refreshToken := jwt.New(jwt.SigningMethodHS256)

  refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
  refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
  refreshTokenClaims["iat"] = time.Now().UTC().Unix()
  refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

  // Create signed refresh token
  signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
  if err != nil {
    return TokenPairs{}, fmt.Errorf("unable to sign refresh token: %v", err)
  }

  // Create Token Pairs with signed tokens 
  tokenPairs := TokenPairs{ Token: signedAccesstoken, RefreshToken: signedRefreshToken}
  return tokenPairs, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie{
  return &http.Cookie{
    Name: j.CookieName,
    Path: j.CookiePath,
    Value: refreshToken,
    Expires: time.Now().Add(j.RefreshExpiry),
    MaxAge: int(j.RefreshExpiry.Seconds()),
    SameSite: http.SameSiteStrictMode,
    Domain: j.CookieDomain,
    HttpOnly: true,
    Secure: true,
  }
}

func (j *Auth) GetExpiredRefreshCookie(refreshToken string) *http.Cookie{
  return &http.Cookie{
    Name: j.CookieName,
    Path: j.CookiePath,
    Value: "",
    Expires: time.Unix(0,0),
    MaxAge: -1,
    SameSite: http.SameSiteStrictMode,
    Domain: j.CookieDomain,
    HttpOnly: true,
    Secure: true,
  }
}
