package main

import "fmt"

type Contact struct {
	Nom   string
	Email string
}

func main() {
	var contacts []Contact
	var count int            

	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Choisis une option : ")
		fmt.Scanln(&choix)

		if choix == 1 {
			var nom, email string
			fmt.Print("Nom du contact : ")
			fmt.Scanln(&nom)
			fmt.Print("Email du contact : ")
			fmt.Scanln(&email)

			contacts[count] = nom + " - " + email
			count++
			fmt.Println("✅ Contact ajouté !")

		} else if choix == 2 {
			if count == 0 {
				fmt.Println("📭 Aucun contact pour le moment.")
			} else {
				fmt.Println("\n📋 Liste des contacts :")
				for i := 0; i < count; i++ {
					fmt.Printf("%d. %s\n", i+1, contacts[i])
				}
			}

		}
	}
}