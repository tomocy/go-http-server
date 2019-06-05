package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func New(root string) *Server {
	return &Server{
		root:   root,
		router: chi.NewRouter(),
	}
}

type Server struct {
	root   string
	router chi.Router
}

func (s *Server) ListenAndServe(addr string) error {
	s.registerRoutes()
	log.Printf("listen and serve on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) registerRoutes() {
	s.router.Get("/*", http.FileServer(http.Dir(s.root)).ServeHTTP)
}
