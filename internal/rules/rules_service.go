package rules

import (
	"errors"
	"fmt"
	"math"
	"receipt_processor/internal/server/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type RulesService struct {
}

// Given a receipt, this code will execute all the rules and calculate the price
func (s *RulesService) CalculatePoints(r *models.Receipt) (int, error) {

	var totalPoints = 0

	if points, err := s.RetailerNameRule(r.Retailer); err == nil {
		totalPoints += points
	}

	if points, err := s.TotalWholeNumberRule(r.Total); err == nil {
		totalPoints += points
	}

	if points, err := s.TotalMultipleRule(r.Total); err == nil {
		totalPoints += points
	}

	if points, err := s.PurchaseDateRule(r.PurchaseDate.String()); err == nil {
		totalPoints += points
	}

	if points, err := s.PurchaseTimeRule(r.PurchaseTime, false); err == nil {
		totalPoints += points
	}

	if points, err := s.CalculateItemPoints(r.Items); err == nil {
		totalPoints += points
	}

	return totalPoints, nil
}

// executes the rules for items
func (s *RulesService) CalculateItemPoints(items []models.Item) (int, error) {

	var totalPoints = 0

	if points, err := s.ItemPairsRule(items); err == nil {
		totalPoints += points
	}

	for _, item := range items {

		if points, err := s.ItemDescriptionRule(item); err == nil {
			totalPoints += points
		}

	}

	return totalPoints, nil
}

/*
* 5 points for every two items on the receipt.
 */
func (s *RulesService) ItemPairsRule(items []models.Item) (int, error) {
	if items == nil {
		return 0, errors.New("nil items not allowed")
	}
	return (len(items) / 2 * 5), nil
}

/**
 * If the trimmed length of the item description is a multiple of 3,
 * multiply the price by `0.2` and round up to the nearest integer.
 **/
func (r *RulesService) ItemDescriptionRule(item models.Item) (int, error) {

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

/*
* One point for every alphanumeric character in the retailer name.
 */
func (r *RulesService) RetailerNameRule(name string) (int, error) {

	var points = 0
	fmt.Print("hmm")

	for _, char := range name {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points += 1
		}
	}
	return points, nil
}

/*
 * 50 points if the total is a round dollar amount with no cents.
 */
func (r *RulesService) TotalWholeNumberRule(total string) (int, error) {

	var points = 0

	value, err := strconv.ParseFloat(total, 64)

	if err != nil {
		return 0, nil
	}

	if value > 0 && math.Mod(value, 1) == 0 {
		points += 50
	}
	return points, nil
}

/*
 * 25 points if the total is a multiple of `0.25`.
 */
func (r *RulesService) TotalMultipleRule(total string) (int, error) {

	var points = 0

	value, err := strconv.ParseFloat(total, 64)

	if err != nil {
		return 0, nil
	}

	if math.Mod(value*4, 1) == 0 {
		points += 25
	}
	return points, nil
}

/*
 * 6 points if the day in the purchase date is odd.
 */
func (r *RulesService) PurchaseDateRule(purchaseDate string) (int, error) {

	var points = 0

	d, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return 0, errors.New("error parsing date from string")
	}

	if d.Day()%2 != 0 {
		//even
		points += 6
	}
	return points, nil
}

/*
 * 10 points if the time of purchase is after 2:00pm and before 4:00pm.
 */
func (r *RulesService) PurchaseTimeRule(purchaseTime string, inclusive bool) (int, error) {

	var points = 0

	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")

	t, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0, errors.New("error parsing time from string")
	}

	if !inclusive {
		if t.After(startTime) && t.Before(endTime) {
			points += 10
		}
	} else {
		if !t.Before(startTime) && !t.After(endTime) {
			points += 10
		}
	}

	return points, nil
}
