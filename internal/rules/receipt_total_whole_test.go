package rules

import (
	"receipt_processor/internal/server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiptTotalWholeRule(t *testing.T) {
	testData := []struct {
		name   string
		total  string
		points int
		err    error
	}{
		{"whole_number", "23.0", 50, nil},
		{"not_whole_number", "23.5", 0, nil},
		{"negative_number", "-23.0", 0, nil},
		{"blank_name", "", 0, nil},
	}

	r := &ReceiptTotalWholeRule{}

	for _, test := range testData {

		receipt := &models.Receipt{
			Total: test.total,
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
