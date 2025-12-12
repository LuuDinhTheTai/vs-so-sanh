package http

import (
	"log/slog"
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/device"
	_ "vs-so-sanh/internal/model"
	"vs-so-sanh/util/response"
	"vs-so-sanh/web/page"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	deviceUseCase device.UseCase
	brandUseCase  brand.UseCase
}

func NewDeviceHandler(useCase device.UseCase, brandUseCase brand.UseCase) device.Handler {
	return &DeviceHandler{
		deviceUseCase: useCase,
		brandUseCase:  brandUseCase,
	}
}

func (p *DeviceHandler) HomePage(ctx *gin.Context) {
	response.HTML(ctx, page.HomePage(p.brandUseCase, p.deviceUseCase))
}

func (p *DeviceHandler) Details(ctx *gin.Context) {
	deviceName := ctx.Param("modelName")
	phone, err := p.deviceUseCase.FindByName(deviceName)
	if err != nil {
		slog.Error("(handler) details: ", err)
		return
	}
	response.HTML(ctx, page.DetailsPage(phone))
}
