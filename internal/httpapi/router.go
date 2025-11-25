package httpapi

import "net/http"

type RouterDependencies struct {
	HealthHandler *HealthHandler
}

func NewRouter(dependencies RouterDependencies) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health/live", dependencies.HealthHandler.Liveness)
	mux.HandleFunc("/health/ready", dependencies.HealthHandler.Readiness)

	return mux
}
