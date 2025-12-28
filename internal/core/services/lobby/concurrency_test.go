package lobby

import (
	"sync"
	"testing"

	"github.com/google/uuid"
)

func TestConcurrency(t *testing.T) {
	srvc := NewService()
	var wg sync.WaitGroup

	// Number of concurrent operations
	numGoroutines := 50

	// 1. Open Lobbies Concurrently
	lobbyIDs := make([]uuid.UUID, numGoroutines)
	var lobbyIDsMu sync.Mutex

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id, err := srvc.OpenLobby()
			if err != nil {
				t.Errorf("Failed to open lobby: %v", err)
				return
			}
			lobbyIDsMu.Lock()
			lobbyIDs[i] = id
			lobbyIDsMu.Unlock()
		}(i)
	}
	wg.Wait()

	// 2. Concurrent Client Connections and Messaging
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lobbyIDsMu.Lock()
			lobbyID := lobbyIDs[i]
			lobbyIDsMu.Unlock()

			// Connect Client
			client, err := srvc.ConnectsClient(lobbyID, "user1")
			if err != nil {
				t.Errorf("Failed to connect client: %v", err)
				return
			}

			// Send Messages
			_, err = srvc.CreateMessage(lobbyID, client.ID, "Hello world")
			if err != nil {
				t.Errorf("Failed to create message: %v", err)
				return
			}

			// Read Messages
			_, err = srvc.GetAllMessagesInLobby(lobbyID)
			if err != nil {
				t.Errorf("Failed to get messages: %v", err)
				return
			}

			// Get Clients
			_, err = srvc.GetClientsInLobby(lobbyID)
			if err != nil {
				t.Errorf("Failed to get clients: %v", err)
				return
			}

			// Disconnect
			_, err = srvc.DisconnectsClient(lobbyID, client)
			if err != nil {
				t.Errorf("Failed to disconnect client: %v", err)
				return
			}
		}(i)
	}
	wg.Wait()

	// 3. Concurrent Random Access on Shared Lobby
	// Create one shared lobby
	sharedLobbyID, _ := srvc.OpenLobby()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Try to connect
			client, err := srvc.ConnectsClient(sharedLobbyID, "concurrent_user")
			if err == nil {
				// If connected, do some stuff then leave
				srvc.GetAllMessagesInLobby(sharedLobbyID)
				srvc.DisconnectsClient(sharedLobbyID, client)
			}
			// If lobby full, that's fine, just check for panics
			srvc.GetClientsInLobby(sharedLobbyID)
		}()
	}
	wg.Wait()
}
