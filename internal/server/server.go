package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"qa-service/internal/config"
	"qa-service/internal/di"
	"qa-service/internal/httpapi"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	srv *http.Server

	stopC chan os.Signal
}

func New(c *di.Container, cfg *config.Config) *Server {
	mux := chi.NewRouter()

	httpapi.RegisterRoutes(mux, c)

	s := &http.Server{
		Addr:    cfg.HTTP.Address,
		Handler: mux,
	}

	return &Server{
		srv:   s,
		stopC: make(chan os.Signal, 1),
	}
}

func (s *Server) Run() {
	signal.Notify(s.stopC, os.Interrupt)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-s.stopC
	log.Println("Shutting down server...")

	s.stop()

	log.Println("Server stopped gracefully")
}

func (s *Server) stop() {
	s.srv.Shutdown(context.Background())
}
