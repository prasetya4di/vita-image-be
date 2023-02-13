package impl

import (
	"log"
	"mime/multipart"
	"vita-image-service/data/entity/image"
	"vita-image-service/repository"
	"vita-image-service/usecase"
)

type uploadImage struct {
	repo repository.MessageRepository
}

func NewUploadImage(messageRepository repository.MessageRepository) usecase.UploadImage {
	return &uploadImage{
		messageRepository,
	}
}

func (sm *uploadImage) Invoke(email string, file multipart.File, header *multipart.FileHeader) ([]image.Possibility, error) {
	message, err := sm.repo.Insert(email, file, header)
	if err != nil {
		log.Fatalf("error insert image: %v", err)
		return nil, err
	}

	result := sm.repo.ScanImageMessage(message)
	return result, nil
}
