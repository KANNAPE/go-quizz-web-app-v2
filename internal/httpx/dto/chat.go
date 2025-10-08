package dto

import "go-quizz/m/internal/core/domain"

type Message struct {
	Sender   Client `json:"sender"`
	TimeSent string `json:"time_sent"`
	Body     string `json:"content"`
}

type ChatRoom struct {
	Messages []Message `json:"messages"`
}

func (messageDto *Message) FromMessage(message *domain.Message) {
	messageDto.Sender = Client{}
	messageDto.Sender.FromClient(message.Sender)

	messageDto.TimeSent = message.TimeSent.GoString()
	messageDto.Body = message.Body
}

func (chatRoomDto *ChatRoom) FromChatRoom(chatRoom *domain.ChatRoom) {
	for _, message := range chatRoom.Messages {
		messageDto := Message{}
		messageDto.FromMessage(message)

		chatRoomDto.Messages = append(chatRoomDto.Messages, messageDto)
	}
}
