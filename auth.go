package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

// secrect key
var signKey = []byte("test_secret_key")

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

func (u *User) valid() bool {
	if u.Username == "hit" && u.Password == "123" {
		return true
	} else {
		return false
	}
}

type JWTResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh_token"`
}

func siginToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return tokenString
}

func refreshToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "bar",
		"exp":     time.Now().Add(time.Minute * 1).Unix(),
		"refresh": true,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return tokenString
}

func getTokens() *JWTResponse {
	token := siginToken()
	refresh := refreshToken()

	return &JWTResponse{
		Token:   token,
		Refresh: refresh,
	}
}
