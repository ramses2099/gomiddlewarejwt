package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type webServer struct {
	s    *http.Server
	port string
}

// when with try to  running multiple file package main
// use go run . command
func main() {
	server := newWebSever("3001")
	router()
	server.start()
}

func router() {
	http.HandleFunc("/sigin", userHandler())
	http.HandleFunc("/secure", middlewareValidaToken(secureHandler()))

}

func newWebSever(port string) *webServer {
	s := &http.Server{
		Addr: ":" + port,
	}
	return &webServer{
		s:    s,
		port: port,
	}
}

// start the server
func (w *webServer) start() {
	fmt.Printf("Server running in http://localhost:%s", w.port)
	log.Fatal((w.s.ListenAndServe()))

}

// middleware function
func middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := r.Method
		log.Println(m)

		//call fuction next
		next.ServeHTTP(w, r)
	}

}

//valida token
func middlewareValidaToken(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if validaToken(bearerToken) {
			//call fuction next
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "invalid token", http.StatusUnauthorized)
	}

}

func validaToken(s string) bool {
	tokenString := strings.Replace(s, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return signKey, nil
	})

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return token.Valid
}

//genera token
func userHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" {
			body := r.Body
			defer body.Close()
			var user User
			err := json.NewDecoder(body).Decode(&user)
			if err != nil {
				http.Error(w, "cannot decode body to json", http.StatusMethodNotAllowed)
				return
			}

			if user.valid() {
				// return jwt
				res := getTokens()
				_ = json.NewEncoder(w).Encode(&res)

			} else {
				http.Error(w, "bad credentials", http.StatusUnauthorized)
				return
			}

		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

//genera token
func secureHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("hello secure"))
	}
}
