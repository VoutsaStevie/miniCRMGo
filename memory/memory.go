package memory

import "crm/contact"
import "fmt"

type MemoryStorage struct {
	data []*contact.Contact
}

func (m *MemoryStorage) GetAll() []*contact.Contact {
	return m.data
}

func (m *MemoryStorage) Add(c *contact.Contact) {
	m.data = append(m.data, c)
}

func (m *MemoryStorage) Update(index int, c *contact.Contact) error {
	if index < 0 || index >= len(m.data) {
		return fmt.Errorf("index invalide")
	}
	m.data[index] = c
	return nil
}

func (m *MemoryStorage) Delete(index int) error {
	if index < 0 || index >= len(m.data) {
		return fmt.Errorf("index invalide")
	}
	m.data = append(m.data[:index], m.data[index+1:]...)
	return nil
}
