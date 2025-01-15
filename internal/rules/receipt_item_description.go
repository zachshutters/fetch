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

// If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
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
