package device

import (
	"context"
	"vs-so-sanh/internal/device/delivery/dto"
	"vs-so-sanh/internal/model"
)

type UseCase interface {
	FindTop20() ([]model.Device, error)
	FindByName(c context.Context, deviceName string) (*dto.DeviceResponse, error)
}
