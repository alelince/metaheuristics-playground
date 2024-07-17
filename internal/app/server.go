package app

import (
	"context"
	"net/http"
	"time"

	handlers "github.com/alelince/metaheuristics-playground/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Address    string
	HttpServer *http.Server
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
		HttpServer: &http.Server{
			Addr:    address,
			Handler: newRouter(),
		},
	}
}

func (s *Server) Start() error {
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func (s *Server) IsRunning() bool {
	return s.HttpServer != nil
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/problems", handlers.GetProblems)

	return r
}
