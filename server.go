package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	config      *Config
	userService *UserService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.users)
	mux.HandleFunc("/user", s.user)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}
func (s *Server) user(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	id, err := strconv.Atoi(q)
	if err != nil {
		log.Fatal(err)
	}
	user := s.userService.FindById(id)
	bytes, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *Server) users(w http.ResponseWriter, r *http.Request) {
	users := s.userService.FindAll()
	bytes, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(service *UserService) *Server {
	return &Server{
		userService: service,
	}
}
