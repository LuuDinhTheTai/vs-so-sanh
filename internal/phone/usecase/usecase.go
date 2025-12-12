package usecase

import "vs-so-sanh/internal/phone"

type PhoneUseCase struct {
	repository phone.Repository
}

func NewPhoneUseCase(repository phone.Repository) *PhoneUseCase {
	return &PhoneUseCase{repository: repository}
}
