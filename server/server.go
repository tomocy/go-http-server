package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	router chi.Router
}

func New() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}

func (s *Server) ListenAndServe(addr string) error {
	s.registerRoutes()
	log.Printf("listen and serve on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) registerRoutes() {
	s.router.Get("/*", http.FileServer(http.Dir("public")).ServeHTTP)
}
