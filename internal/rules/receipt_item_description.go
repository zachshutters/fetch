package rules

import (
	"errors"
	"math"
	"receipt_processor/internal/server/models"
	"strconv"
	"strings"
)

type ReceiptItemDescriptionRule struct {
}

// 6 points if the day in the purchase date is odd.
func (r *ReceiptItemDescriptionRule) Execute(item *models.Item) (int, error) {
	var points = 0

	trimmed := strings.TrimSpace(item.ShortDescription)
	trimmedLength := len(trimmed)

	isMultipleOf3 := trimmedLength%3 == 0

	if isMultipleOf3 {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, errors.New("error parsing float value from string")
		}
		points += int(math.Ceil(price * 0.2))
	}
	return points, nil
}
