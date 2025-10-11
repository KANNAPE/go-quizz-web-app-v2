package message

import (
	"errors"
	"go-quizz/m/internal/core/domain"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	messages map[uuid.UUID]*domain.Message
}

func NewService() *Service {
	return &Service{
		messages: make(map[uuid.UUID]*domain.Message),
	}
}

func (messageSrvc *Service) Create(body string) (uuid.UUID, error) {
	if strings.Trim(body, "") == "" {
		return uuid.Nil, errors.New("can't send an empty message")
	}

	newMessageID := uuid.New()
	if _, ok := messageSrvc.messages[newMessageID]; ok {
		return newMessageID, errors.New("message ID duplicates are forbidden")
	}

	timeSent := time.Now()

	messageSrvc.messages[newMessageID] = domain.NewMessage(newMessageID, timeSent, body)

	return newMessageID, nil
}

func (messageSrvc *Service) GetAll() []domain.Message {
	var messages []domain.Message

	for _, message := range messageSrvc.messages {
		messageCopy := domain.Message{
			ID:       message.ID,
			Sender:   message.Sender,
			TimeSent: message.TimeSent,
			Body:     message.Body,
		}

		messages = append(messages, messageCopy)
	}

	return messages
}

func (messageSrvc *Service) Get(messageID uuid.UUID) (domain.Message, error) {
	if _, ok := messageSrvc.messages[messageID]; !ok {
		return domain.Message{}, errors.New("message doesn't exist")
	}

	return *messageSrvc.messages[messageID], nil
}
