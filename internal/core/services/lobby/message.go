package lobby

import (
	"errors"
	"go-quizz/m/internal/core/domain"

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
			Sender:   message.Sender,
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
		return domain.Message{}, errors.New("invalid message ID")
	}

	return *message, nil
}

func (srvc *LobbyService) CreateMessage(senderID uuid.UUID, body string) (uuid.UUID, error) {
	// check que le client existe dans le lobby

	// check que le body est pas empty (on veut pas de messages vides)

	// créer le message et le rajouter dans le tableau

	// renvoyer l'ID du message nouvellement créé
}
