package usecase

import (
	"vs-so-sanh/internal/device"
	"vs-so-sanh/internal/model"
)

type DeviceUseCase struct {
	repository device.Repository
}

func NewPhoneUseCase(repository device.Repository) device.UseCase {
	return &DeviceUseCase{repository: repository}
}

func (p *DeviceUseCase) FindTop20() ([]model.Device, error) {
	return p.repository.FindTop20(nil)
}
