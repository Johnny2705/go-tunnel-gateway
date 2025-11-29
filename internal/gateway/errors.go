package gateway

import "errors"

var (
	ErrClientNotFound      = errors.New("client not found")
	ErrTunnelNotFound      = errors.New("tunnel not found")
	ErrClientAlreadyExists = errors.New("client already exists")
)
