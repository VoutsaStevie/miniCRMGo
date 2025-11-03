package main

import "fmt"

type Contact struct {
	Nom   string
	Email string
}

func main() {
	var contacts []Contact

	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Modifier un contact")
		fmt.Println("5. Quitter")

		var choix int
		fmt.Print("Choisis une option : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			var nom, email string
			fmt.Print("Nom du contact : ")
			fmt.Scanln(&nom)
			fmt.Print("Email du contact : ")
			fmt.Scanln(&email)
			contacts = append(contacts, Contact{Nom: nom, Email: email})
			fmt.Println("Contact ajouté.")

		case 2:
			if len(contacts) == 0 {
				fmt.Println("Aucun contact pour le moment.")
			} else {
				fmt.Println("\nListe des contacts :")
				for i, c := range contacts {
					fmt.Printf("%d. %s - %s\n", i+1, c.Nom, c.Email)
				}
			}

		case 3:
			if len(contacts) == 0 {
				fmt.Println("Aucun contact à supprimer.")
				continue
			}
			var num int
			fmt.Print("Numéro du contact à supprimer : ")
			fmt.Scanln(&num)
			if num < 1 || num > len(contacts) {
				fmt.Println("Numéro invalide.")
			} else {
				contacts = append(contacts[:num-1], contacts[num:]...)
				fmt.Println("Contact supprimé.")
			}

		case 4:
			if len(contacts) == 0 {
				fmt.Println("Aucun contact à modifier.")
				continue
			}
			var num int
			fmt.Print("Numéro du contact à modifier : ")
			fmt.Scanln(&num)
			if num < 1 || num > len(contacts) {
				fmt.Println("Numéro invalide.")
			} else {
				var nom, email string
				fmt.Print("Nouveau nom : ")
				fmt.Scanln(&nom)
				fmt.Print("Nouvel email : ")
				fmt.Scanln(&email)
				contacts[num-1] = Contact{Nom: nom, Email: email}
				fmt.Println("Contact modifié.")
			}

		case 5:
			fmt.Println("Au revoir.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
