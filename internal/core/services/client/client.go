package client

import (
	"errors"
	"fmt"
	"go-quizz/m/internal/core/domain"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	clients map[uuid.UUID]*domain.Client
}

func NewService() *Service {
	return &Service{
		clients: make(map[uuid.UUID]*domain.Client),
	}
}

func (clientSrvc *Service) Register(username string) uuid.UUID {
	username = strings.Trim(username, " ")
	if username == "" {
		fmt.Println("username can't be empty")
		return uuid.Nil
	}

	newClientID := uuid.New()
	if _, ok := clientSrvc.clients[newClientID]; !ok {
		clientSrvc.clients[newClientID] = domain.NewClient(newClientID, username)
	} else {
		fmt.Println("client already exists")
	}

	return newClientID
}

func (clientSrvc *Service) GetAll() []domain.Client {
	var clients []domain.Client

	for _, client := range clientSrvc.clients {
		clientCopy := domain.Client{
			ID:       client.ID,
			Username: client.Username,
		}

		clients = append(clients, clientCopy)
	}

	return clients
}

func (clientSrvc *Service) Get(clientID uuid.UUID) (domain.Client, error) {
	if _, ok := clientSrvc.clients[clientID]; !ok {
		return domain.Client{}, errors.New("Client doesn't exists!")
	}

	return *clientSrvc.clients[clientID], nil
}
