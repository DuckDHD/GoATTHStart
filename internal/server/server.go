package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"

	"GoATTHStart/internal/config"
	"GoATTHStart/internal/handlers"
)

type Server struct {
	router     *chi.Mux
	logger     *slog.Logger
	config     *config.Config
	handlers   *Handlers
	httpServer *http.Server
}

type Handlers struct {
	Health *handlers.HealthHander
}

func New(cfg *config.Config, logger *slog.Logger, handlers *Handlers) *Server {
	router := chi.NewRouter()

	s := &Server{
		config:   cfg,
		logger:   logger,
		router:   router,
		handlers: handlers,
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
	}

	s.SetupRoutes()

	return s
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) GetHTTPServer() *http.Server {
	return s.httpServer
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Server is Quacking at port " + s.config.Port)
	return s.httpServer.ListenAndServe()
}
