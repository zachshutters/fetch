package models

import "time"

type LedgerEntryItem struct {
	description string
	price       string
	createdAt   time.Time
}

func CreateLedgerEntryItem(description, price string) (LedgerEntryItem, error) {

	// validations?

	return LedgerEntryItem{
		description: description,
		price:       price,
		createdAt:   time.Now(),
	}, nil
}
