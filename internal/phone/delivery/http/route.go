package http

import (
	"net/http"
	"vs-so-sanh/internal/phone"

	"github.com/gin-gonic/gin"
	"maragu.dev/gomponents"
)

func MapPhoneRoutes(r *gin.Engine, handler phone.Handler) {
	r.GET("/", adapt(handler.HomePage))
}

func adapt(h func(ctx *gin.Context) (gomponents.Node, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		n, err := h(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Header("Content-Type", "text/html")
		_ = n.Render(c.Writer)
	}
}
