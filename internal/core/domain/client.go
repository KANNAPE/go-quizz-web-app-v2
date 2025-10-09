package domain

import "github.com/google/uuid"

type Client struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
