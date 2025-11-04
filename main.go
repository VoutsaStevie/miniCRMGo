package main

import (
	"crm/contact"
	"crm/memory"
)

func main() {
	store := &memory.MemoryStorage{}
	contact.Init(store)
	contact.Menu()
}
