package usecase

import (
	"log/slog"
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/brand/delivery/dto"
)

type BrandUseCase struct {
	BrandRepository brand.Repository
}

func NewBrandUseCase(repo brand.Repository) brand.UseCase {
	return &BrandUseCase{
		BrandRepository: repo,
	}
}

func (b *BrandUseCase) FindTop10() ([]dto.BrandResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BrandUseCase) List() ([]dto.BrandResponse, error) {
	brands, err := b.BrandRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var brandResponses []dto.BrandResponse
	for _, br := range brands {
		brandResponses = append(brandResponses, dto.BrandResponse{
			Name: br.Name,
		})
	}

	slog.Info("brand response: ", brandResponses)

	return brandResponses, nil
}
