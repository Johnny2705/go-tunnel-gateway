package httpapi

import (
	"net/http"

	"github.com/Johnny2705/go-tunnel-gateway/internal/health"
)

type HealthHandler struct {
	checker *health.Checker
}

func NewHealthHandler(checker *health.Checker) *HealthHandler {
	return &HealthHandler{
		checker: checker,
	}
}

func (h *HealthHandler) writeStatus(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func (h *HealthHandler) Liveness(w http.ResponseWriter, r *http.Request) {
	err := h.checker.CheckLiveness()

	if err != nil {
		h.writeStatus(w, 500, "unhealthy")
		return
	}
	h.writeStatus(w, 200, "ok")
}

func (h *HealthHandler) Readiness(w http.ResponseWriter, r *http.Request) {
	err := h.checker.CheckReadiness()

	if err != nil {
		h.writeStatus(w, 500, "unhealthy")
		return
	}
	h.writeStatus(w, 200, "ok")
}
