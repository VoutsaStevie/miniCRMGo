package contact

import "fmt"

// Contact représente une personne dans le CRM
type Contact struct {
	Nom   string
	Email string
}

// Storage définit les opérations CRUD
type Storage interface {
	GetAll() []*Contact
	Add(c *Contact)
	Update(index int, c *Contact) error
	Delete(index int) error
}

// store est une variable globale du type Storage
var store Storage

// Init permet d’injecter une implémentation concrète (comme Memory)
func Init(s Storage) {
	store = s
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

	if nom == "" || email == "" {
		fmt.Println("Nom ou email invalide")
		return
	}

	newContact := &Contact{Nom: nom, Email: email}
	store.Add(newContact)
	fmt.Println("✅ Contact ajouté avec succès :", nom)
}

func ListerContacts() {
	contacts := store.GetAll()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact disponible.")
		return
	}

	fmt.Println("\nListe des contacts :")
	for i, c := range contacts {
		fmt.Printf("%d. %s - %s\n", i+1, c.Nom, c.Email)
	}
}

func ModifierContact() {
	contacts := store.GetAll()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à modifier.")
		return
	}

	ListerContacts()
	var index int
	fmt.Print("Entrez l'index du contact à modifier : ")
	fmt.Scanln(&index)

	if index < 1 || index > len(contacts) {
		fmt.Println("Index invalide.")
		return
	}

	var nom, email string
	fmt.Print("Entrez le nouveau nom : ")
	fmt.Scanln(&nom)
	fmt.Print("Entrez le nouvel email : ")
	fmt.Scanln(&email)

	updated := &Contact{Nom: nom, Email: email}
	err := store.Update(index-1, updated)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("✏️ Contact modifié avec succès !")
}

func SupprimerContact() {
	contacts := store.GetAll()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer.")
		return
	}

	ListerContacts()
	var index int
	fmt.Print("Entrez l'index du contact à supprimer : ")
	fmt.Scanln(&index)

	if index < 1 || index > len(contacts) {
		fmt.Println("Index invalide.")
		return
	}

	err := store.Delete(index - 1)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("🗑️ Contact supprimé avec succès !")
}
