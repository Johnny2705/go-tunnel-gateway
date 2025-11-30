package gateway

type Manager interface {
	CreateTunnel(clientID ClientID) (Tunnel, error)
	ListTunnels() []Tunnel
	RegisterClient() ClientID
	UnregisterClient(id ClientID) error
	ListTunnelsByClient(clientID ClientID) []Tunnel
	CloseTunnel(id TunnelID) error
}
