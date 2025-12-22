package lobby

import (
	"errors"
	"go-quizz/m/internal/core/domain"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (srvc *LobbyService) GetAllMessagesInLobby(lobbyID uuid.UUID) ([]domain.Message, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return nil, errors.New("lobby doesn't exist")
	}

	var messages []domain.Message
	for _, message := range lobby.Messages {
		messageCopy := domain.Message{
			ID:       message.ID,
			SenderID: message.SenderID,
			TimeSent: message.TimeSent,
			Body:     message.Body,
		}

		messages = append(messages, messageCopy)
	}

	return messages, nil
}

func (srvc *LobbyService) GetLobbyMessage(lobbyID uuid.UUID, messageID uuid.UUID) (domain.Message, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return domain.Message{}, errors.New("lobby doesn't exist")
	}

	message, ok := lobby.Messages[messageID]
	if !ok {
		return domain.Message{}, errors.New("message wasn't found")
	}

	return *message, nil
}

func (srvc *LobbyService) CreateMessage(lobbyID uuid.UUID, senderID uuid.UUID, body string) (uuid.UUID, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return uuid.Nil, errors.New("lobby doesn't exist")
	}

	if _, ok := lobby.Clients[senderID]; !ok {
		return uuid.Nil, errors.New("client doesn't exist or is not in this lobby")
	}

	if strings.TrimSpace(body) == "" {
		return uuid.Nil, errors.New("body is empty")
	}

	messageID := uuid.New()
	timeSent := time.Now()

	message := domain.NewMessage(messageID, senderID, timeSent, body)

	lobby.Messages[messageID] = message

	return messageID, nil
}
