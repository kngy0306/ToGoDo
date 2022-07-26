package repository

import (
	"go-api/domain/model"
)

type EntryRepository interface {
	FindByID(id int) (*model.Entry, error)
	Store(entry *model.Entry) error
	Update(entry *model.Entry) error
	Delete(entry *model.Entry) error
	FindAll() ([]*model.Entry, error)
}
