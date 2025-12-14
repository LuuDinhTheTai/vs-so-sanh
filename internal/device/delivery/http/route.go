package http

import (
	"vs-so-sanh/internal/device"

	"github.com/gin-gonic/gin"
)

func MapDeviceRoutes(r *gin.Engine, handler device.Handler) {
	r.GET("/", handler.HomePage)
	r.GET("/devices/:modelName", handler.Details)
	r.GET("/compare/:device1/vs/:device2", handler.Compare)
}
