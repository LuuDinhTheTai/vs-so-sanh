package device

import (
	"vs-so-sanh/internal/model"
)

type UseCase interface {
	FindTop20() ([]model.Device, error)
	FindByName(deviceName string) (*model.Device, error)
}
