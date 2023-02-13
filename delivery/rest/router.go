package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"vita-image-service/delivery/rest/handler"
)

func LoadRoutes(handler handler.ImageHandler) {
	router := gin.Default()
	router.POST("/image/:email", handler.UploadImage)
	router.GET("/ping", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
	})
	router.Static("/image", "upload/images")
	err := router.Run(os.Getenv("BASEURL") + ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
