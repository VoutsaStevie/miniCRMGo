package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"crm/internal/storage"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Mini CRM v3!")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Add a contact")
		fmt.Println("2. List contacts")
		fmt.Println("3. Update a contact")
		fmt.Println("4. Delete a contact")
		fmt.Println("5. Exit")
		fmt.Print("Your choice: ")

		choice := readUserChoice(reader)

		switch choice {
		case 1:
			handleAddContact(reader, store)
		case 2:
			handleListContacts(store)
		case 3:
			handleUpdateContact(reader, store)
		case 4:
			handleDeleteContact(reader, store)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

// Les fonctions "handle..." s'occupent de l'interaction avec l'utilisateur
// et appellent la couche de stockage (le store) pour effectuer les opérations.
// Elles sont découplées du stockage : elles fonctionnent avec N'IMPORTE quel Storer.

func handleAddContact(reader *bufio.Reader, store storage.Storer) {
	var (
		name  string
		email string
		err   error
	)

	// Boucle pour le nom (obligatoire)
	for {
		fmt.Print("Enter contact name: ")
		name, err = readLine(reader)
		if err != nil {
			// Erreur système lors de la lecture (EOF, etc.)
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if name != "" {
			break // L'entrée est valide (non vide)
		}
		fmt.Println("Name cannot be empty. Please enter a value.")
	}

	// Boucle pour l'email (obligatoire)
	for {
		fmt.Print("Enter contact email: ")
		email, err = readLine(reader)
		if err != nil {
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if email != "" {
			break // L'entrée est valide (non vide)
		}
		fmt.Println("Email cannot be empty. Please enter a value.")
	}
	// Le reste de la logique d'ajout
	contact := &storage.Contact{Name: name, Email: email}
	err = store.Add(contact)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf(" Contact '%s' added with ID %d.\n", contact.Name, contact.ID)
}

func handleListContacts(store storage.Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println(" No contacts to display.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the contact to update: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingContact.Name)

	fmt.Printf("New name (%s): ", existingContact.Name)
	newName, err := readLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("New email (%s): ", existingContact.Email)
	newEmail, err := readLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully.")
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the contact to delete: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
}

func readLine(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		// Affiche erreur de lecture (EOF, etc.)
		fmt.Printf("Error during read: %v\n", err)
		return "", err
	}

	trimmedInput := strings.TrimSpace(input)
	return trimmedInput, nil
}

func readUserChoice(reader *bufio.Reader) int {
	choice, err := readLine(reader)
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	nbrChoice, err := strconv.Atoi(choice)
	if err != nil {
		return -1
	}
	return nbrChoice
}

func readInteger(reader *bufio.Reader) int {
	id, err := readLine(reader)
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	nbrId, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}

	return nbrId
}
