package local

import (
	"mime/multipart"
	"vita-image-service/data/entity"
)

type MessageDao interface {
	Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error)
}
