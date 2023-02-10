package entity

import "time"

type Message struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Message     string    `json:"message"`
	CreatedDate time.Time `json:"created_date"`
	MessageType string    `json:"message_type"`
	FileType    string    `json:"fileType"`
}
