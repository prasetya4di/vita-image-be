package network

import (
	"vita-image-service/data/entity"
	"vita-image-service/data/entity/image"
)

type MessageService interface {
	SendImageMessage(message entity.Message) []image.Possibility
}
