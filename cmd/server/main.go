package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Johnny2705/go-tunnel-gateway/internal/config"
	"github.com/Johnny2705/go-tunnel-gateway/internal/gateway"
	"github.com/Johnny2705/go-tunnel-gateway/internal/health"
	"github.com/Johnny2705/go-tunnel-gateway/internal/httpapi"
	"github.com/Johnny2705/go-tunnel-gateway/internal/server"
)

func setupLogger() {
	env := os.Getenv("ENV")

	var handler slog.Handler

	if env == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func main() {
	setupLogger()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	healthChecker := health.NewChecker(cfg)
	healthHandler := httpapi.NewHealthHandler(healthChecker)
	memoryManager := gateway.NewMemoryManager()
	gatewayHandler := httpapi.NewGatewayHandler(memoryManager)
	router := httpapi.NewRouter(httpapi.RouterDependencies{
		HealthHandler:  healthHandler,
		GatewayHandler: gatewayHandler,
	})
	srv := server.NewServer(cfg, router)

	go func() {
		err := srv.Start()
		if err != nil {
			slog.Error("failed to start server", slog.Any("error", err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop
	slog.Info("shutdown signal received", slog.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("graceful shutdown failed", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("server exited cleanly")
}
