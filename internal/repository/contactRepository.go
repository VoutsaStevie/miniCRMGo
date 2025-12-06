package repository

import (
	"crm/internal/model"
	"crm/internal/storage"
	"errors"
)

var ErrContactNotFound = errors.New("contact not found")

type ContactRepository struct {
	store storage.Storage
}

func NewContactRepository(s storage.Storage) *ContactRepository {
	return &ContactRepository{store: s}
}

func (r *ContactRepository) Create(c model.Contact) (model.Contact, error) {
	if err := r.store.Save(c); err != nil {
		return model.Contact{}, err
	}
	// return last inserted (memory store)
	all := r.store.FindAll()
	return all[len(all)-1], nil
}

func (r *ContactRepository) GetAll() []model.Contact {
	return r.store.FindAll()
}

func (r *ContactRepository) GetByID(id int) (model.Contact, error) {
	c, ok := r.store.FindByID(id)
	if !ok {
		return model.Contact{}, ErrContactNotFound
	}
	return c, nil
}

func (r *ContactRepository) Update(id int, c model.Contact) error {
	return r.store.Update(id, c)
}

func (r *ContactRepository) Delete(id int) error {
	return r.store.Delete(id)
}
