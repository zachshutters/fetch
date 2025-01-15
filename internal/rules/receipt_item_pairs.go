package rules

import (
	"errors"
	"receipt_processor/internal/server/models"
)

type ReceiptItemPairRule struct {
}

// 5 points for every two items on the receipt.
func (r *ReceiptItemPairRule) Execute(receipt *models.Receipt) (int, error) {
	if receipt.Items == nil {
		return 0, errors.New("nil items not allowed")
	}
	return (len(receipt.Items) / 2 * 5), nil
}
