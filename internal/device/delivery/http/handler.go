package http

import (
	"log/slog"
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/device"
	"vs-so-sanh/internal/device/delivery/dto"
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
	top20Brands, err := p.brandUseCase.FindTop20()
	if err != nil {
		slog.Error("Error fetching top 10 brands: ", err)
		return
	}

	top20Devices, err := p.deviceUseCase.FindTop20()
	if err != nil {
		slog.Error("Error fetching top 10 devices: ", err)
		return
	}

	response.HTML(ctx, page.HomePage(top20Brands, top20Devices))
}

func (p *DeviceHandler) Details(ctx *gin.Context) {
	deviceName := ctx.Param("modelName")
	deviceResponse, err := p.deviceUseCase.FindByName(ctx, deviceName)
	if err != nil {
		slog.Error("(handler) details: \n", err)
		return
	}
	response.HTML(ctx, page.DetailsPage(deviceResponse))
}

func (p *DeviceHandler) Compare(ctx *gin.Context) {
	device1Name := ctx.Param("device1")
	device2Name := ctx.Param("device2")

	device1, err := p.deviceUseCase.FindByName(ctx, device1Name)
	if err != nil {
		slog.Error("(handler) compare: \n", err)
		device1 = dto.EmptyDeviceResponse
	}

	device2, err := p.deviceUseCase.FindByName(ctx, device2Name)
	if err != nil {
		slog.Error("(handler) compare: \n", err)
		device2 = dto.EmptyDeviceResponse
	}

	response.HTML(ctx, page.ComparePage(device1, device2))
}
