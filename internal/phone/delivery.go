package phone

import (
	"github.com/gin-gonic/gin"
	"maragu.dev/gomponents"
)

type Handler interface {
	HomePage(ctx *gin.Context) (gomponents.Node, error)
}
