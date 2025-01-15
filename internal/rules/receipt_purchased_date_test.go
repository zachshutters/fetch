package rules

import (
	"receipt_processor/internal/server/models"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

func TestReceiptPurchaseDateRule(t *testing.T) {
	testData := []struct {
		name         string
		purchaseDate string
		points       int
		err          error
	}{
		{"purchase_data_even", "2022-01-03", 6, nil},
		{"purchase_data_odd", "2022-01-02", 0, nil},
	}

	var r = &ReceiptPurchaseDateRule{}

	for _, test := range testData {

		layout := "2006-01-02"
		purchaseDate, _ := time.Parse(layout, test.purchaseDate)

		receipt := &models.Receipt{
			PurchaseDate: openapi_types.Date{Time: purchaseDate},
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
