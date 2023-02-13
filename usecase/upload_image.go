package usecase

import (
	"mime/multipart"
	"vita-image-service/data/entity/image"
)

type UploadImage interface {
	Invoke(email string, file multipart.File, header *multipart.FileHeader) ([]image.Possibility, error)
}
