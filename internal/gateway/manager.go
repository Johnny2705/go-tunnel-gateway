package gateway

type Manager interface {
	CreateTunnel(clientID ClientID) (Tunnel, error)
	ListTunnels() []Tunnel
	RegisterClient(id ClientID) error
	UnregisterClient(id ClientID) error
	ListTunnelsByClient(clientID ClientID) []Tunnel
	CloseTunnel(id TunnelID) error
}
