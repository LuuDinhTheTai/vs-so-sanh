package brand

import (
	"vs-so-sanh/internal/brand/delivery/dto"
)

type UseCase interface {
	FindTop20() ([]dto.BrandResponse, error)
	List() ([]dto.BrandResponse, error)
}
