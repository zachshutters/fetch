package rules

import (
	"receipt_processor/internal/server/models"
)

type ReceiptPurchaseDateRule struct{}

// 6 points if the day in the purchase date is odd.
func (r *ReceiptPurchaseDateRule) Execute(receipt *models.Receipt) (int, error) {
	var points = 0

	if receipt.PurchaseDate.Day()%2 != 0 {
		//even
		points += 6
	}
	return points, nil
}
