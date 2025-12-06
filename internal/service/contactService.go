package service

import (
	"crm/internal/model"
	"crm/internal/repository"
)

type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(r *repository.ContactRepository) *ContactService {
	return &ContactService{r}
}

func (s *ContactService) AddContact(c model.Contact) (model.Contact, error) {
	return s.repo.Create(c)
}

func (s *ContactService) ListContacts() []model.Contact {
	return s.repo.GetAll()
}

func (s *ContactService) GetContact(id int) (model.Contact, error) {
	return s.repo.GetByID(id)
}

func (s *ContactService) UpdateContact(id int, c model.Contact) error {
	return s.repo.Update(id, c)
}

func (s *ContactService) DeleteContact(id int) error {
	return s.repo.Delete(id)
}
