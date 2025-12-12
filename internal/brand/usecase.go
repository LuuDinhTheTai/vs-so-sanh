package brand

import (
	"vs-so-sanh/internal/brand/delivery/dto"
)

type UseCase interface {
	FindTop10() ([]dto.BrandResponse, error)
	List() ([]dto.BrandResponse, error)
}
