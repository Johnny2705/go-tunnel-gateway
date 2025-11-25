package gateway

import "time"

type ClientID string
type TunnelID string

type Client struct {
	ID ClientID
}

type Tunnel struct {
	ID        TunnelID
	ClientId  ClientID
	CreatedAt time.Time
}
