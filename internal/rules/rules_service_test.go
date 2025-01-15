package rules

import (
	"receipt_processor/internal/server/models"
	"receipt_processor/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullExample(t *testing.T) {

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

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {

			r, err := util.LoadTestReceipt(test.filename)

			if err != nil {
				t.Fail()
			}

			points, err := rulesSvc.CalculatePoints(r)
			assert.Nil(t, err)

			assert.Equal(t, test.points, points, "expected points and actual points are not equal")
		})
	}
}

/*
 * One point for every alphanumeric character in the retailer name.
 */
func TestAlphaNumericRule(t *testing.T) {

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

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			points, err := rulesSvc.RetailerNameRule(test.retailerName)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points and actual points are not equal")
		})
	}
}

// 50 points if the total is a round dollar amount with no cents.
func TestTotalRule(t *testing.T) {

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

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			points, err := rulesSvc.TotalWholeNumberRule(test.total)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points and actual points are not equal")
		})
	}
}

/*
* 25 points if the total is a multiple of `0.25`.
 */
func TestTotalMultipleRule(t *testing.T) {

	testData := []struct {
		name   string
		total  string
		points int
		err    error
	}{
		{"whole_number_multiple", "23.0", 25, nil},
		{"fractional_multiple", "23.25", 25, nil},
		{"fractional_multiple", "23.5", 25, nil},
		{"not_multiple", "23.9", 0, nil},
		{"blank", "", 0, nil},
	}

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			points, err := rulesSvc.TotalMultipleRule(test.total)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points and actual points are not equal")
		})
	}
}

/*
 * 6 points if the day in the purchase date is odd.
 */
func TestPurchaseDateRule(t *testing.T) {

	testData := []struct {
		name         string
		purchaseDate string
		points       int
		err          error
	}{
		{"purchase_data_even", "2022-01-01", 6, nil},
		{"purchase_data_odd", "2022-01-02", 0, nil},
		{"purchase_data_invalid", "2022-01-68", 0, nil},
		{"purchase_data_invalid", "adslfhajdskfy234642", 0, nil},
		{"blank", "", 0, nil},
	}

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			points, err := rulesSvc.PurchaseDateRule(test.purchaseDate)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points and actual points are not equal")
		})
	}
}

/*
 * 10 points if the time of purchase is after 2:00pm and before 4:00pm.
 */
func TestPurchaseTimeRule(t *testing.T) {

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

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			points, err := rulesSvc.PurchaseTimeRule(test.purchaseTime, test.inclusive)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points does not match actual points given")
		})
	}
}

func TestItemPurchasePriceRule(t *testing.T) {

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

	var rulesSvc = &RulesService{}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			item := models.Item{
				ShortDescription: test.description,
				Price:            test.price,
			}
			points, err := rulesSvc.ItemDescriptionRule(item)

			if test.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, test.points, points, "expected points does not match actual points given")
		})
	}
}
