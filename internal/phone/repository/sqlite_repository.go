package repository

import (
	"context"
	"vs-so-sanh/internal/model"
	"vs-so-sanh/internal/phone"

	"gorm.io/gorm"
)

type SqlitePhoneRepository struct {
	db *gorm.DB
}

func NewSqlitePhoneRepository(db *gorm.DB) phone.Repository {
	return &SqlitePhoneRepository{
		db: db,
	}
}

func (s *SqlitePhoneRepository) Save(c context.Context, phone model.Phone) (*model.Phone, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SqlitePhoneRepository) FindByName(c context.Context, name string) (*model.Phone, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SqlitePhoneRepository) Update(c context.Context, phone model.Phone) (*model.Phone, error) {
	//TODO implement me
	panic("implement me")
}
