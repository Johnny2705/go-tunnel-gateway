package gateway

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MemoryManager struct {
	mux     sync.RWMutex
	clients map[ClientID]*Client
	tunnels map[TunnelID]*Tunnel
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		clients: make(map[ClientID]*Client),
		tunnels: make(map[TunnelID]*Tunnel),
	}
}

func (m *MemoryManager) CreateTunnel(clientID ClientID) (Tunnel, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if _, ok := m.clients[clientID]; !ok {
		return Tunnel{}, fmt.Errorf("%w: %s", ErrClientNotFound, clientID)
	}

	tunnelID := TunnelID(uuid.NewString())
	tunnel := &Tunnel{
		ID:        tunnelID,
		ClientID:  clientID,
		CreatedAt: time.Now().UTC(),
	}
	m.tunnels[tunnelID] = tunnel

	return *tunnel, nil
}

func (m *MemoryManager) ListTunnels() []Tunnel {
	m.mux.RLock()
	defer m.mux.RUnlock()
	tunnels := make([]Tunnel, 0, len(m.tunnels))

	for _, tunnel := range m.tunnels {
		tunnels = append(tunnels, *tunnel)
	}

	return tunnels
}

func (m *MemoryManager) RegisterClient(id ClientID) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.clients[id]; ok {
		return fmt.Errorf("%w: %s", ErrClientAlreadyExists, id)
	}

	c := &Client{
		ID: id,
	}
	m.clients[id] = c

	return nil
}

func (m *MemoryManager) UnregisterClient(id ClientID) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.clients[id]; !ok {
		return fmt.Errorf("%w: %s", ErrClientNotFound, id)
	}

	for tid, t := range m.tunnels {
		if t.ClientID == id {
			delete(m.tunnels, tid)
		}
	}
	delete(m.clients, id)

	return nil

}

func (m *MemoryManager) ListTunnelsByClient(clientID ClientID) []Tunnel {
	m.mux.RLock()
	defer m.mux.RUnlock()
	tunnels := make([]Tunnel, 0)

	for _, tunnel := range m.tunnels {
		if tunnel.ClientID == clientID {
			tunnels = append(tunnels, *tunnel)
		}
	}

	return tunnels
}

func (m *MemoryManager) CloseTunnel(id TunnelID) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	if _, ok := m.tunnels[id]; !ok {
		return fmt.Errorf("%w: %s", ErrTunnelNotFound, id)
	}

	delete(m.tunnels, id)
	return nil
}

var _ Manager = (*MemoryManager)(nil)
