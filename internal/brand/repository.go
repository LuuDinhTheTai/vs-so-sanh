package brand

import "vs-so-sanh/internal/model"

type Repository interface {
	FindTop10() ([]model.Brand, error)
	FindAll() ([]model.Brand, error)
}
