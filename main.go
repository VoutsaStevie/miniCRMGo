package main

import "fmt"

type Contact struct {
	Nom   string
	Email string
}

var contacts []Contact 

func main() {
	fmt.Println("--- Mini CRM --- \n 1. Ajouter un contact \n 2. Lister les contacts \n 3. Modifier un contact \n 4. Supprimer un contact \n 5. Quitter")
	fmt.Print("Votre choix : ")
	var choix int
	fmt.Scanln(&choix)

	if choix != 1 && choix != 2 && choix != 3 && choix != 4 && choix != 5 {
		fmt.Println("Choix invalide")
	} else{
		switch choix {
			case 1:
				ajouterContact()
			case 2:
				listerContacts()
			case 3:
				ModifierContact()
			case 4:
				SupprimerContact()
			case 5:
				fmt.Println("A bientôt !")
		}
	
	}

func ajouterContact() {
	var nom, email string
	fmt.Print("Entrez le nom du contact : ")
	fmt.Scanln(&nom)
	fmt.Print("Entrez l'email du contact : ")
	fmt.Scanln(&email)

	newContact := Contact{Nom: nom, Email: email}
	contacts = append(contacts, newContact)

	fmt.Println("Contact " + nom + " ajouté avec succès !")
}

func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact disponible.")
		return
	}
}
