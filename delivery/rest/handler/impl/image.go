package impl

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
		return
	}
	if !isValidImage(file) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid image format"})
		return
	}
	email := c.Param("email")
	possibilities, err := ih.uploadImage.Invoke(email, file, header)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, possibilities)
}

func isValidImage(file multipart.File) bool {
	buff := make([]byte, 512)
	file.Read(buff)
	contentType := http.DetectContentType(buff)
	return contentType == "image/png" || contentType == "image/jpg" || contentType == "image/gif" || contentType == "image/webp" || contentType == "image/jpeg" || contentType == "image/bmp"
}
