package main

import (
	"github.com/joho/godotenv"
	"log"
	"vita-image-service/data/local"
	"vita-image-service/data/local/impl"
	impl2 "vita-image-service/data/network/impl"
	"vita-image-service/delivery/rest"
	impl5 "vita-image-service/delivery/rest/handler/impl"
	impl3 "vita-image-service/repository/impl"
	impl4 "vita-image-service/usecase/impl"
)

func init() {
	//Get environment data
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := local.GetDB()

	messageDao := impl.NewMessageDao(db)
	messageService := impl2.NewMessageService()

	messageRepository := impl3.NewMessageRepository(messageDao, messageService)

	uploadImageUseCase := impl4.NewUploadImage(messageRepository)

	imageHandler := impl5.NewImageHandler(uploadImageUseCase)

	rest.LoadRoutes(imageHandler)
}
