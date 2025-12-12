package http

import (
	"net/http"
	"vs-so-sanh/internal/device"

	"github.com/gin-gonic/gin"
	"maragu.dev/gomponents"
)

func MapDeviceRoutes(r *gin.Engine, handler device.Handler) {
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
