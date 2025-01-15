package receipt

import (
	"receipt_processor/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {

	testData := []struct {
		name     string
		filename string
		points   int
	}{
		{"example", "../receipt/examples/example-one.json", 28},
		{"example", "../receipt/examples/example-two.json", 109},
		{"morning", "../receipt/examples/morning-receipt.json", 15},
		{"simple", "../receipt/examples/simple-receipt.json", 31},
	}

	var svc = NewReceiptService()

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			r, err := util.LoadTestReceipt(test.filename)

			if err != nil {
				t.Fail()
			}

			id, err := svc.ProcessReceipt(r)

			assert.Nil(t, err)
			assert.NotNil(t, id)

			points, err := svc.LookupPoints(id)

			assert.Nil(t, err)

			assert.NotNil(t, id)
			assert.NotNil(t, points)
		})
	}
}
