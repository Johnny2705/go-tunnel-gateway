package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Johnny2705/go-tunnel-gateway/internal/gateway"
)

type GatewayHandler struct{ manager gateway.Manager }

func NewGatewayHandler(m gateway.Manager) *GatewayHandler {
	return &GatewayHandler{
		manager: m,
	}
}

func (h *GatewayHandler) RegisterClient(w http.ResponseWriter, r *http.Request) {

	clientID := h.manager.RegisterClient()

	writeJSON(w, http.StatusCreated, map[string]string{
		"client_id": string(clientID),
	})
}

func (h *GatewayHandler) UnregisterClient(w http.ResponseWriter, r *http.Request) {
	clientID := gateway.ClientID(r.PathValue("id"))
	if clientID == "" {
		writeJSONError(w, http.StatusBadRequest, "invalid client_id")
		return
	}

	err := h.manager.UnregisterClient(clientID)
	if err != nil {
		if errors.Is(err, gateway.ErrClientNotFound) {
			writeJSONError(w, http.StatusNotFound, "no client registered with this id")
		} else {
			writeJSONError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *GatewayHandler) CreateTunnel(w http.ResponseWriter, r *http.Request) {
	var request createTunnelRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if strings.TrimSpace(request.ClientID) == "" {
		writeJSONError(w, http.StatusBadRequest, "client_id is required")
		return
	}

	t, err := h.manager.CreateTunnel(gateway.ClientID(request.ClientID))
	if err != nil {
		var status int
		if errors.Is(err, gateway.ErrClientNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		writeJSONError(w, status, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

func (h *GatewayHandler) ListTunnels(w http.ResponseWriter, r *http.Request) {
	tunnels := h.manager.ListTunnels()
	writeJSON(w, http.StatusOK, tunnels)
}

func (h *GatewayHandler) ListTunnelsByClient(w http.ResponseWriter, r *http.Request) {
	clientID := gateway.ClientID(r.PathValue("id"))
	if clientID == "" {
		writeJSONError(w, http.StatusBadRequest, "invalid client_id")
		return
	}

	tunnels := h.manager.ListTunnelsByClient(clientID)
	writeJSON(w, http.StatusOK, tunnels)
}

func (h *GatewayHandler) CloseTunnel(w http.ResponseWriter, r *http.Request) {
	tunnelID := gateway.TunnelID(r.PathValue("id"))
	if tunnelID == "" {
		writeJSONError(w, http.StatusBadRequest, "invalid tunnel_id")
		return
	}

	err := h.manager.CloseTunnel(tunnelID)
	if err != nil {
		if errors.Is(err, gateway.ErrTunnelNotFound) {
			writeJSONError(w, http.StatusNotFound, "tunnel not found")
		} else {
			writeJSONError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
