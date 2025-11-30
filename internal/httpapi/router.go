package httpapi

import "net/http"

type RouterDependencies struct {
	HealthHandler  *HealthHandler
	GatewayHandler *GatewayHandler
}

func NewRouter(dependencies RouterDependencies) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health/live", dependencies.HealthHandler.Liveness)
	mux.HandleFunc("GET /health/ready", dependencies.HealthHandler.Readiness)
	mux.HandleFunc("POST /clients", dependencies.GatewayHandler.RegisterClient)
	mux.HandleFunc("DELETE /clients/{id}", dependencies.GatewayHandler.UnregisterClient)
	mux.HandleFunc("GET /clients/{id}/tunnels", dependencies.GatewayHandler.ListTunnelsByClient)
	mux.HandleFunc("POST /tunnels", dependencies.GatewayHandler.CreateTunnel)
	mux.HandleFunc("GET /tunnels", dependencies.GatewayHandler.ListTunnels)
	mux.HandleFunc("DELETE /tunnels/{id}", dependencies.GatewayHandler.CloseTunnel)

	return mux
}
