package rules

import (
	"errors"
	"receipt_processor/internal/server/models"
	"time"
)

type ReceiptPurchaseTimeRule struct {
	Inclusive bool
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func (r *ReceiptPurchaseTimeRule) Execute(receipt *models.Receipt) (int, error) {
	var points = 0

	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")

	t, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, errors.New("error parsing time from string")
	}

	if !r.Inclusive {
		if t.After(startTime) && t.Before(endTime) {
			points += 10
		}
	} else {
		if !t.Before(startTime) && !t.After(endTime) {
			points += 10
		}
	}

	return points, nil
}
