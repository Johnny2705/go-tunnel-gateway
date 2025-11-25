package gateway

import "sync"

type MemoryManager struct {
	mux     sync.RWMutex
	clients map[ClientID]Client
	tunnels map[TunnelID]Tunnel
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		clients: make(map[ClientID]Client),
		tunnels: make(map[TunnelID]Tunnel),
	}
}

func (m *MemoryManager) CreateTunnel(clientID ClientID) (Tunnel, error) {

}

func (m *MemoryManager) ListTunnels() []Tunnel {

}

func (m *MemoryManager) RegisterClient(id ClientID) error {

}

func (m *MemoryManager) UnregisterClient(id ClientID) error {

}

func (m *MemoryManager) ListTunnelsByClient(clientID ClientID) []Tunnel {

}

func (m *MemoryManager) CloseTunnel(id TunnelID) error {

}
