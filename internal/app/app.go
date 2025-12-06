package app

import (
	"crm/internal/handler"
	"crm/internal/repository"
	"crm/internal/service"
	"crm/internal/storage"
	"fmt"
	"net/http"
	"strconv"
)

func Run() {
	// stockage mémoire
	mem := storage.NewMemoryStorage()

	// repository / service / handler
	repo := repository.NewContactRepository(mem)
	svc := service.NewContactService(repo)
	h := handler.NewContactHandler(svc)

	// page d'accueil + gestion complète
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Mini CRM</title>
</head>
<body>
<h1>Mini CRM — Gérer les contacts</h1>

<h2>Ajouter un contact</h2>
<form id="contactForm">
  <label>Nom:<br><input type="text" id="name" required></label><br><br>
  <label>Email:<br><input type="email" id="email" required></label><br><br>
  <button type="submit">Ajouter</button>
</form>

<h2>Liste des contacts</h2>
<table border="1" id="contactsTable" style="border-collapse: collapse;">
  <thead>
    <tr><th>ID</th><th>Nom</th><th>Email</th><th>Actions</th></tr>
  </thead>
  <tbody></tbody>
</table>

<div id="result" style="margin-top:20px;color:green"></div>

<script>
async function fetchContacts() {
  const resp = await fetch('/contacts');
  const contacts = await resp.json();
  const tbody = document.querySelector('#contactsTable tbody');
  tbody.innerHTML = '';
  contacts.forEach(function(c) {
    const row = document.createElement('tr');
    row.innerHTML =
      '<td>' + c.id + '</td>' +
      '<td><input type="text" value="' + c.name + '" id="name-' + c.id + '"></td>' +
      '<td><input type="email" value="' + c.email + '" id="email-' + c.id + '"></td>' +
      '<td>' +
      '<button onclick="updateContact(' + c.id + ')">Modifier</button> ' +
      '<button onclick="deleteContact(' + c.id + ')">Supprimer</button>' +
      '</td>';
    tbody.appendChild(row);
  });
}

async function addContact() {
  const name = document.getElementById('name').value;
  const email = document.getElementById('email').value;
  await fetch('/contacts/add', {
    method: 'POST',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify({name, email})
  });
  document.getElementById('name').value = '';
  document.getElementById('email').value = '';
  fetchContacts();
}

async function updateContact(id) {
  const name = document.getElementById('name-' + id).value;
  const email = document.getElementById('email-' + id).value;
  await fetch('/contacts/update?id=' + id + '&name=' + encodeURIComponent(name) + '&email=' + encodeURIComponent(email));
  fetchContacts();
}

async function deleteContact(id) {
  await fetch('/contacts/remove?id=' + id);
  fetchContacts();
}

document.getElementById('contactForm').addEventListener('submit', function(e) {
  e.preventDefault();
  addContact();
});

fetchContacts();
</script>

</body>
</html>`))
	})

	// lister tous les contacts
	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.GetContacts(w, r)
	})

	// ajouter un contact (POST JSON)
	http.HandleFunc("/contacts/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.AddContact(w, r)
	})

	// endpoints navigateur pour update/delete
	http.HandleFunc("/contacts/update", func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	if err := h.UpdateContactByParams(id, name, email); err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Indiquer que la mise à jour a réussi
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Contact mis à jour !"))
})


	http.HandleFunc("/contacts/remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		if err := h.DeleteContactByID(id); err != nil {
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Contact supprimé !"))
	})

	// GET/PUT/DELETE /contacts/{id} (existants)
	http.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		h.ContactByID(w, r)
	})

	fmt.Println("CRM running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
