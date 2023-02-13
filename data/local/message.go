package local

import "vita-image-service/data/entity"

type MessageDao interface {
	Insert(message entity.Message) (entity.Message, error)
	Inserts(messages []entity.Message) ([]entity.Message, error)
}
