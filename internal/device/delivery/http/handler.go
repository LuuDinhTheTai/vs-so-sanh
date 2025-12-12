package http

import (
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/device"
	_ "vs-so-sanh/internal/model"
	"vs-so-sanh/util/response"
	"vs-so-sanh/web/page"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	phoneUseCase device.UseCase
	brandUseCase brand.UseCase
}

func NewDeviceHandler(useCase device.UseCase, brandUseCase brand.UseCase) device.Handler {
	return &DeviceHandler{
		phoneUseCase: useCase,
		brandUseCase: brandUseCase,
	}
}

func (p *DeviceHandler) HomePage(ctx *gin.Context) {
	response.HTML(ctx, page.HomePage(p.brandUseCase, p.phoneUseCase))
}
