package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"context"

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
	server *http.Server
}

func (s *Server) ListenAndServe(addr string) error {
	s.registerRoutes()
	log.Println("registered routes")

	s.server = &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	errCh := s.listenAndServe(addr)
	log.Printf("listen and serve files in %s on %s\n", s.root, addr)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	sig := <-sigCh
	log.Printf("received signal: %s\n", sig)

	log.Println("shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}
	log.Println("shut down the server")

	return <-errCh
}

func (s *Server) registerRoutes() {
	s.router.Get("/*", http.FileServer(http.Dir(s.root)).ServeHTTP)
}

func (s *Server) listenAndServe(addr string) <-chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)

		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			ch <- err
			return
		}

		ch <- nil
	}()

	return ch
}
