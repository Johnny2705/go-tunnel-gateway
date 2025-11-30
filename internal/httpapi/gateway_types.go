package httpapi

import "time"

type registerClientRequest struct {
	ClientID string `json:"client_id"`
}

type createTunnelRequest struct {
	ClientID string `json:"client_id"`
}

type tunnelResponse struct {
	ID        string    `json:"id"`
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
}
