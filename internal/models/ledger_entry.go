package models

import (
	"log"
	"receipt_processor/internal/server/models"
	"time"

	"github.com/google/uuid"
)

type LedgerEntry struct {
	Id           string
	retailer     string
	purchaseDate time.Time
	purchaseTime string
	items        []LedgerEntryItem
	total        string
	Points       int
	createdAt    time.Time
}

func CreateLedgerEntry(r *models.Receipt, points int) (LedgerEntry, error) {

	var ledgerEntryItems = []LedgerEntryItem{}

	for _, item := range r.Items {

		ledgerEntryItem, err := CreateLedgerEntryItem(item.ShortDescription, item.Price)

		if err != nil {
			log.Fatalf("Critical error creating ledger entry for item: %s, price: %s. Error: %v", item.ShortDescription, item.Price, err)
		}

		ledgerEntryItems = append(ledgerEntryItems, ledgerEntryItem)

	}

	return LedgerEntry{
		Id:           uuid.New().String(),
		retailer:     r.Retailer,
		purchaseDate: r.PurchaseDate.Time,
		purchaseTime: r.PurchaseTime,
		total:        r.Total,
		Points:       points,
		items:        ledgerEntryItems,
		createdAt:    time.Now(),
	}, nil
}
