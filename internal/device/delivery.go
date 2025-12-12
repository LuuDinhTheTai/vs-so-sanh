package device

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	HomePage(ctx *gin.Context)
	Details(ctx *gin.Context)
}
