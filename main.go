package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().String()
		_, _ = w.Write([]byte(fmt.Sprintf("the time is %s", t)))
	})

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
