package storage

import "fmt"

// Contact est notre structure de données centrale.
type Contact struct {
	ID    int
	Name  string
	Email string
}

// Storer est notre CONTRAT de stockage.
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie pas de COMMENT c'est fait
// (en mémoire, fichier, BDD), seulement de CE QUI peut être fait.
type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetByID(id int) (*Contact, error) // Méthode pratique pour la suite
	Update(id int, newName string, newEmail string) error
	Delete(id int) error
}

var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
}
