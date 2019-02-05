package main

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	config      *Config
	userService *UserService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", s.users)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) users(w http.ResponseWriter, r *http.Request) {
	users := s.userService.FindAll()
	bytes, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *Config, service *UserService) *Server {
	return &Server{
		config:      config,
		userService: service,
	}
}
