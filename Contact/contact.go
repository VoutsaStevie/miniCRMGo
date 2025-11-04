package contact

import "fmt"

type Contact struct {
	Nom   string
	Email string
}

var contacts []Contact


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
				fmt.Println("A bientôt !")
				return
			default:
				fmt.Println("Choix invalide")
		}
	}
}

func AjouterContact() {
	var nom, email string
	fmt.Print("Entrez le nom du contact : ")
	fmt.Scanln(&nom)
	fmt.Print("Entrez l'email du contact : ")
	fmt.Scanln(&email)

	newContact := Contact{Nom: nom, Email: email}
	contacts = append(contacts, newContact)

	fmt.Println("Contact ajouté avec succès :", nom)
}

func ListerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact disponible.")
		return
	}

	fmt.Println("\n Liste des contacts :")
	for i, contact := range contacts {
		fmt.Printf("%d. %s - %s\n", i+1, contact.Nom, contact.Email)
	}
}

func ModifierContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à modifier.")
		return
	}

	var index int
	ListerContacts()
	fmt.Print("Entrez l'index du contact à modifier : ")
	fmt.Scanln(&index)

	if index < 0 || index >= len(contacts) {
		fmt.Println("Index invalide.")
		return
	}

	var nom, email string
	fmt.Print("Entrez le nouveau nom : ")
	fmt.Scanln(&nom)
	fmt.Print("Entrez le nouvel email : ")
	fmt.Scanln(&email)

	contacts[index-1] = Contact{Nom: nom, Email: email}
	fmt.Println("Contact modifié avec succès !")
}

func SupprimerContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer.")
		return
	}

	var index int
	ListerContacts()
	fmt.Print("Entrez l'index du contact à supprimer : ")
	fmt.Scanln(&index)

	if index < 0 || index >= len(contacts) {
		fmt.Println("Index invalide.")
		return
	}

	contacts = append(contacts[:index], contacts[index+1:]...)
	fmt.Println(" Contact supprimé avec succès !")
}