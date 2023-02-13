package handler

import "github.com/gin-gonic/gin"

type ImageHandler interface {
	UploadImage(c *gin.Context)
}
