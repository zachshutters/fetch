package ledger

import (
	"receipt_processor/internal/database"
	"receipt_processor/internal/models"
)

type LedgerService struct {
	db *database.InMemoryDb
}

func New() LedgerService {
	return LedgerService{
		db: database.New(),
	}
}

func (l *LedgerService) Insert(entry models.LedgerEntry) error {

	return l.db.Put(entry.Id, entry)
}

func (l *LedgerService) Find(id string) (models.LedgerEntry, bool) {
	entry, found := l.db.Get(id)

	return entry, found
}
