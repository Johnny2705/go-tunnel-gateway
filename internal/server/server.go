package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Johnny2705/go-tunnel-gateway/internal/config"
	"github.com/Johnny2705/go-tunnel-gateway/internal/httpapi"
)

type Server struct {
	cfg        *config.Config
	httpServer *http.Server
}

func NewServer(cfg *config.Config, healthHandler *httpapi.HealthHandler) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/live", healthHandler.Liveness)
	mux.HandleFunc("/health/ready", healthHandler.Readiness)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	return &Server{
		cfg:        cfg,
		httpServer: &server,
	}
}

func (s *Server) Start() error {
	slog.Info("HTTP server listening", slog.Int("port", s.cfg.Port))
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("HTTP server shutting down...")
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("could not shutdown server gracefully: %w", err)
	}
	slog.Info("HTTP server stopped")

	return nil
}
