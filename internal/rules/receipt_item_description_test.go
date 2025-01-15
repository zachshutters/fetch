package rules

import (
	"receipt_processor/internal/server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemDescriptionRule(t *testing.T) {

	testData := []struct {
		name        string
		description string
		price       string
		points      int
		err         error
	}{
		{"purchase_time_inclusive_start_before", "Emils Cheese Pizza", "12.25", 3, nil},
		{"purchase_time_inclusive_start_before", "Klarbrunn 12-PK 12 FL OZ", "12.00", 3, nil},
		{"purchase_time_inclusive_start_before", "Klarbrunn 12-PK 12 FL OZ", "", 0, nil},
		{"purchase_time_inclusive_start_before", "Klarbrunn 12-PK 12 FL OZ", "abcd", 0, nil},
		{"purchase_time_inclusive_start_before", "Klarbrunn 12-PK 12 FL OZ", "0", 0, nil},
		{"purchase_time_inclusive_start_before", "Klarbrunn 12-PK 12 FL OZ", "-1", 0, nil},
		{"blank", "", "", 0, nil},
	}

	var r = &ReceiptItemDescriptionRule{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			item := models.Item{
				ShortDescription: test.description,
				Price:            test.price,
			}
			points, err := r.Execute(&item)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points does not match actual points given")
		})
	}
}
