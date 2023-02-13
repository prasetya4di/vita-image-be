package network

import (
	"vita-image-service/data/entity"
	"vita-image-service/data/entity/image"
)

type MessageService interface {
	ScanImageMessage(message entity.Message) []image.Possibility
}
