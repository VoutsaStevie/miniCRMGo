package storage

import "crm/internal/model"

type Storage interface {
	Save(c model.Contact) error
	FindAll() []model.Contact
	FindByID(id int) (model.Contact, bool)
	Update(id int, updated model.Contact) error
	Delete(id int) error
}
