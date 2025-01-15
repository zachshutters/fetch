package rules

import (
	"receipt_processor/internal/server/models"
	"unicode"
)

type ReceiptRetailerRule struct{}

// One point for every alphanumeric character in the retailer name.
func (r *ReceiptRetailerRule) Execute(receipt *models.Receipt) (int, error) {
	var points = 0

	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points += 1
		}
	}
	return points, nil
}
