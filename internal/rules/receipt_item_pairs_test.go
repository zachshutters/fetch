package rules

import (
	"receipt_processor/internal/server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemPairsRule(t *testing.T) {
	testData := []struct {
		name      string
		itemCount int
		points    int
	}{
		{"two_items", 2, 5},
		{"three_items", 3, 5},
		{"four_items", 4, 10},
		{"zero_items", 0, 0},
	}

	var r = &ReceiptItemPairRule{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			// Create the required number of items based on itemCount
			var items []models.Item
			for i := 0; i < test.itemCount; i++ {
				item := models.Item{
					ShortDescription: "test",
					Price:            "23.0", // Assign a price for testing purposes
				}
				items = append(items, item)
			}

			receipt := &models.Receipt{
				Items: items,
			}

			points, err := r.Execute(receipt)

			if test.itemCount == 0 {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points does not match actual points given")
		})
	}
}
