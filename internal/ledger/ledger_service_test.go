package ledger

import (
	"receipt_processor/internal/models"
	"receipt_processor/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {

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

	var svc = New()

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			r, err := util.LoadTestReceipt(test.filename)

			if err != nil {
				t.Fail()
			}

			entry, err := models.CreateLedgerEntry(r, test.points)

			assert.Nil(t, err)
			err = svc.Insert(entry)

			assert.Nil(t, err)
			assert.NotNil(t, entry.Id)
			assert.Equal(t, test.points, entry.Points)
		})
	}
}

func TestFind(t *testing.T) {

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

	var svc = New()

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			r, err := util.LoadTestReceipt(test.filename)

			if err != nil {
				t.Fail()
			}

			existingEntry, err := models.CreateLedgerEntry(r, test.points)
			assert.Nil(t, err)

			err = svc.Insert(existingEntry)
			assert.Nil(t, err)

			entry, found := svc.Find(existingEntry.Id)

			assert.True(t, found)
			assert.NotNil(t, entry.Id)
			assert.Equal(t, existingEntry.Id, entry.Id)
		})
	}
}
