package phone

import (
	"context"
	"vs-so-sanh/internal/model"
)

type Repository interface {
	Save(c context.Context, phone model.Phone) (*model.Phone, error)
	FindByName(c context.Context, name string) (*model.Phone, error)
	Update(c context.Context, phone model.Phone) (*model.Phone, error)
}
