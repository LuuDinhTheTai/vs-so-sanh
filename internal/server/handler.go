package server

import (
	repository2 "vs-so-sanh/internal/brand/repository"
	usecase2 "vs-so-sanh/internal/brand/usecase"
	"vs-so-sanh/internal/device/delivery/http"
	"vs-so-sanh/internal/device/repository"
	"vs-so-sanh/internal/device/usecase"
)

func (s *Server) MapHandlers() error {
	// Init repository
	brandRepository := repository2.NewMongoDbBrandRepository(s.cfg, s.client)
	mongoDbDeviceRepository := repository.NewMongoDbDeviceRepository(s.cfg, s.client)

	// Init usecase
	brandUseCase := usecase2.NewBrandUseCase(brandRepository)
	deviceUseCase := usecase.NewPhoneUseCase(mongoDbDeviceRepository)

	// Init handler
	phoneHandler := http.NewDeviceHandler(deviceUseCase, brandUseCase)

	http.MapDeviceRoutes(s.r, phoneHandler)

	return nil
}
