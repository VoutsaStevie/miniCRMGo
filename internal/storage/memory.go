package storage

import (
	"errors"
	"crm/internal/model"
)

var ErrNotFound = errors.New("not found")

type MemoryStorage struct {
	contacts []model.Contact
	nextID   int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		contacts: []model.Contact{},
		nextID:   1,
	}
}

func (m *MemoryStorage) Save(c model.Contact) error {
	if c.ID == 0 {
		c.ID = m.nextID
		m.nextID++
	}
	m.contacts = append(m.contacts, c)
	return nil
}

func (m *MemoryStorage) FindAll() []model.Contact {
	return m.contacts
}

func (m *MemoryStorage) FindByID(id int) (model.Contact, bool) {
	for _, c := range m.contacts {
		if c.ID == id {
			return c, true
		}
	}
	return model.Contact{}, false
}

func (m *MemoryStorage) Update(id int, updated model.Contact) error {
	for i, c := range m.contacts {
		if c.ID == id {
			// update only non-empty fields
			if updated.Name != "" {
				c.Name = updated.Name
			}
			if updated.Email != "" {
				c.Email = updated.Email
			}
			m.contacts[i] = c
			return nil
		}
	}
	return ErrNotFound
}

func (m *MemoryStorage) Delete(id int) error {
	for i, c := range m.contacts {
		if c.ID == id {
			m.contacts = append(m.contacts[:i], m.contacts[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
