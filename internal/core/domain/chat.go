package domain

import "time"

type Message struct {
	Sender   *Client
	TimeSent time.Time
	Body     string
}

type ChatRoom struct {
	Messages []*Message
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}
