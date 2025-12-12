package device

import (
	"vs-so-sanh/internal/model"
)

type UseCase interface {
	FindTop20() ([]model.Device, error)
}
