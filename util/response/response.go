package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"maragu.dev/gomponents"
)

func HTML(ctx *gin.Context, node gomponents.Node) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")

	ctx.Status(http.StatusOK)

	err := node.Render(ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}
