package main

import "fmt"


func main() {
	fmt.Println("--- Mini CRM ---")
	fmt.Println("1. ajouter un contac")
	fmt.Println("2. lister les contacts")
	fmt.Println("4. supprimer un contact")
	fmt.Println("5. quitter")
var choix int
	fmt.Print(" Choisis une option : ")
	fmt.Scanln(&choix)

}

if choix == 1 {
		fmt.Println("\n📋 Liste des contacts :")
		for i, c := range contacts {
			fmt.Printf("%d. %s\n", i+1, c)
		}
	}