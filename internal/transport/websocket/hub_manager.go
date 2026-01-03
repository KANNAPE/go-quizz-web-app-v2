package websocket

import (
	"sync"

	"github.com/google/uuid"
)

type HubManager struct {
	hubs    map[uuid.UUID]*Hub
	mu      sync.RWMutex
	service MessageService
}

func NewHubManager(service MessageService) *HubManager {
	return &HubManager{
		hubs:    make(map[uuid.UUID]*Hub),
		service: service,
	}
}

func (mgr *HubManager) GetHub(lobbyID uuid.UUID) *Hub {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	hub, exists := mgr.hubs[lobbyID]
	if !exists {
		hub = NewHub(lobbyID, mgr.service)
		mgr.hubs[lobbyID] = hub

		// starting the hub in a new goroutine
		go hub.Run()
	}

	return hub
}

func (mgr *HubManager) RemoveHub(lobbyID uuid.UUID) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	// closing channels
	closingHub := mgr.hubs[lobbyID]
	close(closingHub.Broadcast)
	close(closingHub.Register)
	close(closingHub.Unregister)

	// TODO: making a call to the API to delete the lobby?

	delete(mgr.hubs, lobbyID)
}
