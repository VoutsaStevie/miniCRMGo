package contact

import (
	"errors"
	"fmt"
	"strings"
)


type Contact struct {
	ID    int
	Nom   string
	Email string
}


var contacts = make(map[int]*Contact)
var nextID = 1


func NewContact(nom, email string) (*Contact, error) {
	nom = strings.TrimSpace(nom)
	email = strings.TrimSpace(email)

	if nom == "" {
		return nil, errors.New("le nom ne peut pas être vide")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("email invalide")
	}

	c := &Contact{
		ID:    nextID,
		Nom:   nom,
		Email: email,
	}
	nextID++
	return c, nil
}

func (c *Contact) Afficher() {
	fmt.Printf("ID %d : %s - %s\n", c.ID, c.Nom, c.Email)
}

func (c *Contact) Modifier(nouveauNom, nouvelEmail string) error {
	nouveauNom = strings.TrimSpace(nouveauNom)
	nouvelEmail = strings.TrimSpace(nouvelEmail)

	if nouveauNom == "" {
		return errors.New("le nom ne peut pas être vide")
	}
	if !strings.Contains(nouvelEmail, "@") {
		return errors.New("email invalide")
	}

	c.Nom = nouveauNom
	c.Email = nouvelEmail
	return nil
}


func AjouterContact() {
	var nom, email string
	fmt.Print("Entrez le nom du contact : ")
	fmt.Scanln(&nom)
	fmt.Print("Entrez l'email du contact : ")
	fmt.Scanln(&email)

	c, err := NewContact(nom, email)
	if err != nil {
		fmt.Println(" Erreur :", err)
		return
	}

	contacts[c.ID] = c
	fmt.Println(" Contact ajouté avec succès (ID:", c.ID, ")")
}

func ListerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact disponible.")
		return
	}

	fmt.Println("\n Liste des contacts :")
	for _, c := range contacts {
		c.Afficher()
	}
}

func ModifierContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à modifier.")
		return
	}

	ListerContacts()
	var id int
	fmt.Print("Entrez l'ID du contact à modifier : ")
	fmt.Scanln(&id)

	c, exists := contacts[id]
	if !exists {
		fmt.Println(" ID invalide.")
		return
	}

	var nom, email string
	fmt.Print("Nouveau nom : ")
	fmt.Scanln(&nom)
	fmt.Print("Nouvel email : ")
	fmt.Scanln(&email)

	if err := c.Modifier(nom, email); err != nil {
		fmt.Println(" Erreur :", err)
		return
	}

	fmt.Println(" Contact modifié avec succès !")
}

func SupprimerContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer.")
		return
	}

	ListerContacts()
	var id int
	fmt.Print("Entrez l'ID du contact à supprimer : ")
	fmt.Scanln(&id)

	_, exists := contacts[id]
	if !exists {
		fmt.Println(" ID invalide.")
		return
	}

	delete(contacts, id)
	fmt.Println(" Contact supprimé avec succès !")
}


func Menu() {
	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Modifier un contact")
		fmt.Println("4. Supprimer un contact")
		fmt.Println("5. Quitter")

		fmt.Print("Votre choix : ")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			AjouterContact()
		case 2:
			ListerContacts()
		case 3:
			ModifierContact()
		case 4:
			SupprimerContact()
		case 5:
			fmt.Println("À bientôt ")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
