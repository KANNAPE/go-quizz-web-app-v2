package dto

type Message struct {
	Sender   string `json:"sender_id"`
	TimeSent string `json:"time_sent"`
	Body     string `json:"content"`
}

type ChatRoom struct {
	Messages []Message `json:"messages"`
}
