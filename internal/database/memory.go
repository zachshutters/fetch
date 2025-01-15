package database

import (
	"receipt_processor/internal/models"
	"sync"
)

type InMemoryDb struct {
	// store  map[string]int
	store map[string]models.LedgerEntry
	mutex sync.Mutex
}

func (db *InMemoryDb) Put(id string, entry models.LedgerEntry) error {

	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.store[id] = entry

	return nil
}

func (db *InMemoryDb) Get(id string) (models.LedgerEntry, bool) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	entry, exists := db.store[id]
	return entry, exists
}

func New() *InMemoryDb {
	return &InMemoryDb{
		store: make(map[string]models.LedgerEntry),
	}
}
