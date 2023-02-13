package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"vita-image-service/delivery/rest/handler"
)

func LoadRoutes(handler handler.ImageHandler) {
	router := gin.Default()
	router.POST("/image/:email", handler.UploadImage)
	err := router.Run(os.Getenv("BASEURL") + ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
