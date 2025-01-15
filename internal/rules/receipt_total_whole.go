package rules

import (
	"math"
	"receipt_processor/internal/server/models"
	"strconv"
)

type ReceiptTotalWholeRule struct{}

// 50 points if the total is a round dollar amount with no cents.
func (r *ReceiptTotalWholeRule) Execute(receipt *models.Receipt) (int, error) {

	var points = 0

	value, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		return 0, nil
	}

	if value > 0 && math.Mod(value, 1) == 0 {
		points += 50
	}
	return points, nil
}
