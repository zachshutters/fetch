package rules

import (
	"receipt_processor/internal/server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiptPurchaseTimeRule(t *testing.T) {
	testData := []struct {
		name         string
		purchaseTime string
		inclusive    bool
		points       int
		err          error
	}{
		{"purchase_time_inclusive_start_before", "13:59", true, 0, nil},
		{"purchase_time_inclusive_start", "14:00", true, 10, nil},
		{"purchase_time_exclusive_start_1", "14:00", false, 0, nil},
		{"purchase_time_exclusive_start_2", "14:01", false, 10, nil},
		{"purchase_time_end", "15:59", false, 10, nil},
		{"purchase_time_end_after", "16:01", false, 0, nil},
		{"purchase_time_invalid", "16:99", false, 0, nil},
		{"purchase_time_invalid", "asdfasdfasdf", false, 0, nil},
		{"blank", "", false, 0, nil},
	}

	for _, test := range testData {

		r := &ReceiptPurchaseTimeRule{
			Inclusive: test.inclusive,
		}

		receipt := &models.Receipt{
			PurchaseTime: test.purchaseTime,
		}

		t.Run(test.name, func(t *testing.T) {
			points, err := r.Execute(receipt)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points does not match actual points given")
		})
	}
}
