package handler

import (
	"crm/internal/model"
	"crm/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ContactHandler struct {
	service *service.ContactService
}

func NewContactHandler(s *service.ContactService) *ContactHandler {
	return &ContactHandler{s}
}

func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.ListContacts())
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		// fallback pour formulaire POST
		if err := r.ParseForm(); err == nil {
			name := r.FormValue("name")
			email := r.FormValue("email")
			c := model.Contact{Name: name, Email: email}
			created, _ := h.service.AddContact(c)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(created)
			return
		}
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	var c model.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	created, err := h.service.AddContact(c)
	if err != nil {
		http.Error(w, "Failed to create contact: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *ContactHandler) ContactByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
	if idStr == "" || idStr == "/" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	idStr = strings.TrimSuffix(idStr, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		c, err := h.service.GetContact(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	case http.MethodPut:
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
			return
		}
		var upd model.Contact
		if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if err := h.service.UpdateContact(id, upd); err != nil {
			http.Error(w, "Update failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if err := h.service.DeleteContact(id); err != nil {
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ContactHandler) UpdateContactByParams(id int, name, email string) error {
	contact, err := h.service.GetContact(id)
	if err != nil {
		return err
	}
	if name != "" {
		contact.Name = name
	}
	if email != "" {
		contact.Email = email
	}
	return h.service.UpdateContact(id, contact)
}


func (h *ContactHandler) DeleteContactByID(id int) error {
	return h.service.DeleteContact(id)
}
