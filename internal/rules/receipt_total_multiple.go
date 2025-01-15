package rules

import (
	"math"
	"receipt_processor/internal/server/models"
	"strconv"
)

type ReceiptTotalMultipleRule struct{}

// 25 points if the total is a multiple of 0.25.
func (r *ReceiptTotalMultipleRule) Execute(receipt *models.Receipt) (int, error) {

	var points = 0

	value, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		return 0, nil
	}

	if math.Mod(value*4, 1) == 0 {
		points += 25
	}
	return points, nil
}
