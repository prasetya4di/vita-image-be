package repository

import (
	"mime/multipart"
	"vita-image-service/data/entity"
	"vita-image-service/data/entity/image"
)

type MessageRepository interface {
	Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error)
	ScanImageMessage(message entity.Message) []image.Possibility
}
