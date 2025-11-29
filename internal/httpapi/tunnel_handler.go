package httpapi

import "github.com/Johnny2705/go-tunnel-gateway/internal/gateway"

type TunnelHandler struct{ manager gateway.Manager }

func NewTunnelHandler(m gateway.Manager) *TunnelHandler {
	return &TunnelHandler{
		manager: m,
	}
}
