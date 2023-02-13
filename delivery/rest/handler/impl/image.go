package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vita-image-service/delivery/rest/handler"
	"vita-image-service/usecase"
)

type imageHandler struct {
	uploadImage usecase.UploadImage
}

func NewImageHandler(uploadImage usecase.UploadImage) handler.ImageHandler {
	return &imageHandler{
		uploadImage,
	}
}

func (ih *imageHandler) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	email := c.Param("email")
	possibilities, err := ih.uploadImage.Invoke(email, file, header)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusCreated, possibilities)
}
