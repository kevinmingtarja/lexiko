package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router *mux.Router
	db     *redis.Client
}

func (s *server) routes() {
	// Handler functions don’t actually handle the requests, they return a function that does.
	// This gives us a closure environment in which our handler can operate.
	// If a particular handler has a dependency, take it as an argument.
	// Reference: https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html
	//s.router.HandleFunc("/chat", s.handleChatSetID()).Methods("POST")

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
}

func newHTTPServer(db *redis.Client) *server {
	srv := &server{
		router: mux.NewRouter(),
		db: db,
	}
	srv.routes()
	return srv
}

// Implementing ServeHTTP turns the server type into a http.Handler.
// Hence, server can be used wherever http.Handler can (e.g. http.ListenAndServe).
// Inside, we simply pass the execution to the router.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}