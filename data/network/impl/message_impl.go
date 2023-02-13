package impl

import (
	"log"
	"vita-image-service/data/entity"
	"vita-image-service/data/entity/image"
	"vita-image-service/data/network"
)

type messageService struct{}

func NewMessageService() network.MessageService {
	return &messageService{}
}

func (m *messageService) SendImageMessage(message entity.Message) []image.Possibility {
	vision, err := network.GetGoogleVision()
	if err != nil {
		log.Fatalf("error when init google vision : %v", err)
		return nil
	}
	defer vision.Close()

	return nil
}
