package phone

import (
	"context"
	"vs-so-sanh/internal/model"
)

type Repository interface {
	Save(c context.Context, phone model.Device) (*model.Device, error)
	FindByName(c context.Context, name string) (*model.Device, error)
	Update(c context.Context, phone model.Device) (*model.Device, error)
}
