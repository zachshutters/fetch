package rules

import (
	"receipt_processor/internal/server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiptRetailRule(t *testing.T) {
	testData := []struct {
		name         string
		retailerName string
		points       int
		err          error
	}{
		{"hasnonumbers", "target", 6, nil},
		{"hasnumbers", "target123", 9, nil},
		{"blank_name", "", 0, nil},
	}

	r := &ReceiptRetailerRule{}

	for _, test := range testData {

		receipt := &models.Receipt{
			Retailer: test.retailerName,
		}

		t.Run(test.name, func(t *testing.T) {
			points, err := r.Execute(receipt)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points and actual points are not equal")
		})
	}
}
