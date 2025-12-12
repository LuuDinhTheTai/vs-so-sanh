package server

import (
	repository2 "vs-so-sanh/internal/brand/repository"
	usecase2 "vs-so-sanh/internal/brand/usecase"
	"vs-so-sanh/internal/phone/delivery/http"
	"vs-so-sanh/internal/phone/repository"
	"vs-so-sanh/internal/phone/usecase"
)

func (s *Server) MapHandlers() error {
	// Init repository
	brandRepository := repository2.NewMongoDbBrandRepository(s.cfg, s.client)
	phoneRepository := repository.NewSqlitePhoneRepository(s.db)

	// Init usecase
	brandUseCase := usecase2.NewBrandUseCase(brandRepository)
	phoneUseCase := usecase.NewPhoneUseCase(phoneRepository)

	// Init handler
	phoneHandler := http.NewPhoneHandler(phoneUseCase, brandUseCase)

	http.MapPhoneRoutes(s.r, phoneHandler)

	return nil
}
