package domain

import "github.com/google/uuid"

type Client struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

func NewClient(id uuid.UUID, username string) *Client {
	return &Client{
		ID:       id,
		Username: username,
	}
}
