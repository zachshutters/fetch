package rules

import "receipt_processor/internal/server/models"

type ReceiptRules interface {
	Execute(r *models.Receipt) (int, error)
}

type ReceiptItemRules interface {
	Execute(item *models.Item) (int, error)
}
