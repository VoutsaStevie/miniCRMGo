package storage

import "fmt"

// MemoryStore est une implémentation CONCRÈTE de l'interface Storer.
// Elle utilise une map en mémoire pour stocker les données.
// Elle respecte le contrat Storer car elle possède toutes les méthodes demandées.
type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

// NewMemoryStore est un "constructeur" qui initialise proprement notre store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {
	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	// On crée une slice pour éviter de retourner directement
	// une référence à la map interne.
	var allContacts []*Contact
	for _, c := range ms.contacts {
		allContacts = append(allContacts, c)
	}
	return allContacts, nil
}

func (ms *MemoryStore) GetByID(id int) (*Contact, error) {
	contact, ok := ms.contacts[id]
	if !ok {
		return nil, fmt.Errorf("contact avec l'ID %d non trouvé", id)
	}
	return contact, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	contact, err := ms.GetByID(id)
	if err != nil {
		return err // Retourne l'erreur "non trouvé"
	}

	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return fmt.Errorf("contact avec l'ID %d non trouvé", id)
	}
	delete(ms.contacts, id)
	return nil
}
