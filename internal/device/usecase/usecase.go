package usecase

import (
	"context"
	"fmt"
	"vs-so-sanh/internal/device"
	"vs-so-sanh/internal/device/delivery/dto"
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

func (p *DeviceUseCase) FindByName(c context.Context, deviceName string) (*dto.DeviceResponse, error) {
	devices, err := p.repository.FindByName(c, deviceName)
	if err != nil {
		return nil, fmt.Errorf("(usecase) FindByName: %w", err)
	}

	return dto.From(devices), nil
}
