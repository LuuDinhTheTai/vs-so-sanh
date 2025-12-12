package server

import (
	"vs-so-sanh/internal/phone/delivery/http"
	"vs-so-sanh/internal/phone/repository"
	"vs-so-sanh/internal/phone/usecase"
)

func (s *Server) MapHandlers() error {
	// Init repository
	phoneRepository := repository.NewSqlitePhoneRepository(s.db)

	// Init usecase
	phoneUseCase := usecase.NewPhoneUseCase(phoneRepository)

	// Init handler
	phoneHandler := http.NewPhoneHandler(phoneUseCase)

	http.MapPhoneRoutes(s.r, phoneHandler)

	return nil
}
