package impl

import (
	"mime/multipart"
	"vita-image-service/data/entity"
	"vita-image-service/data/entity/image"
	"vita-image-service/data/local"
	"vita-image-service/data/network"
	"vita-image-service/repository"
)

type messageRepository struct {
	messageDao     local.MessageDao
	messageService network.MessageService
}

func NewMessageRepository(dao local.MessageDao, service network.MessageService) repository.MessageRepository {
	return &messageRepository{
		dao,
		service,
	}
}

func (mr *messageRepository) Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error) {
	return mr.messageDao.Insert(email, file, header)
}

func (mr *messageRepository) ScanImageMessage(message entity.Message) []image.Possibility {
	return mr.messageService.ScanImageMessage(message)
}
