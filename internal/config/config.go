package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port            int
	Env             string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func Load() (*Config, error) {

	cfg := Config{
		Port:            8080,
		Env:             "dev",
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     60 * time.Second,
		ShutdownTimeout: 5 * time.Second,
	}

	portStr := os.Getenv("PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT env var %q: %w", portStr, err)
		}
		if port <= 0 {
			return nil, fmt.Errorf("invalid PORT env var: must be > 0, got %d", port)
		}
		cfg.Port = port
	}

	env := os.Getenv("ENV")
	if env != "" {
		switch env {
		case "dev", "staging", "prod":
			cfg.Env = env
		default:
			return nil, fmt.Errorf("invalid ENV env var: must be one of dev|staging|prod, got %q", env)
		}
	}

	readTimeoutStr := os.Getenv("GATEWAY_READ_TIMEOUT_SECONDS")
	if readTimeoutStr != "" {
		seconds, err := strconv.Atoi(readTimeoutStr)
		if err != nil {
			return nil, fmt.Errorf("invalid GATEWAY_READ_TIMEOUT_SECONDS env var %q: %w", readTimeoutStr, err)
		}
		if seconds <= 0 {
			return nil, fmt.Errorf("invalid GATEWAY_READ_TIMEOUT_SECONDS env var: must be > 0, got %d", seconds)
		}
		cfg.ReadTimeout = time.Duration(seconds) * time.Second
	}

	writeTimeoutStr := os.Getenv("GATEWAY_WRITE_TIMEOUT_SECONDS")
	if writeTimeoutStr != "" {
		seconds, err := strconv.Atoi(writeTimeoutStr)
		if err != nil {
			return nil, fmt.Errorf("invalid GATEWAY_WRITE_TIMEOUT_SECONDS env var %q: %w", writeTimeoutStr, err)
		}
		if seconds <= 0 {
			return nil, fmt.Errorf("invalid GATEWAY_WRITE_TIMEOUT_SECONDS env var: must be > 0, got %d", seconds)
		}
		cfg.WriteTimeout = time.Duration(seconds) * time.Second
	}

	idleTimeoutStr := os.Getenv("GATEWAY_IDLE_TIMEOUT_SECONDS")
	if idleTimeoutStr != "" {
		seconds, err := strconv.Atoi(idleTimeoutStr)
		if err != nil {
			return nil, fmt.Errorf("invalid GATEWAY_IDLE_TIMEOUT_SECONDS env var %q: %w", idleTimeoutStr, err)
		}
		if seconds <= 0 {
			return nil, fmt.Errorf("invalid GATEWAY_IDLE_TIMEOUT_SECONDS env var: must be > 0, got %d", seconds)
		}
		cfg.IdleTimeout = time.Duration(seconds) * time.Second
	}

	shutdownTimeoutStr := os.Getenv("GATEWAY_SHUTDOWN_TIMEOUT_SECONDS")
	if shutdownTimeoutStr != "" {
		seconds, err := strconv.Atoi(shutdownTimeoutStr)
		if err != nil {
			return nil, fmt.Errorf("invalid GATEWAY_SHUTDOWN_TIMEOUT_SECONDS env var %q: %w", shutdownTimeoutStr, err)
		}
		if seconds <= 0 {
			return nil, fmt.Errorf("invalid GATEWAY_SHUTDOWN_TIMEOUT_SECONDS env var: must be > 0, got %d", seconds)
		}
		cfg.ShutdownTimeout = time.Duration(seconds) * time.Second
	}

	return &cfg, nil
}
