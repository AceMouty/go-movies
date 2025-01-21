package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer        string
	Audiance      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
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
	tokenPairs := TokenPairs{Token: signedAccesstoken, RefreshToken: signedRefreshToken}
	return tokenPairs, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	// add to header
	w.Header().Add("Vary", "Authorization")

	// get auth header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil, errors.New("no auth header present")
	}

	// split header parts
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", nil, errors.New("invalid auth header")
	}

	// check for Bearer
	if headerParts[0] != "Bearer" {
		return "", nil, errors.New("invalid auth header")
	}

	// validate token
	token := headerParts[1]

	claims := &Claims{}

	// parse token and grab claims
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})

	isExpiredToken := strings.HasPrefix(err.Error(), "token is expired by")
	if err != nil && isExpiredToken {
		return "", nil, errors.New("expired token")
	} else if err != nil {
		return "", nil, err
	}

	// did we issue the token
	if claims.Issuer != j.Issuer {
		return "", nil, errors.New("invalid issuer")
	}

	return token, claims, nil
}
