package rules

import (
	"fmt"
	"receipt_processor/internal/server/models"
)

type RulesService struct {
	ReceiptRules     []ReceiptRules
	ReceiptItemRules []ReceiptItemRules
}

func New() *RulesService {
	return &RulesService{
		ReceiptRules: []ReceiptRules{
			&ReceiptRetailerRule{},
			&ReceiptTotalWholeRule{},
			&ReceiptTotalMultipleRule{},
			&ReceiptPurchaseTimeRule{},
			&ReceiptPurchaseDateRule{},
			&ReceiptItemPairRule{},
		},
		ReceiptItemRules: []ReceiptItemRules{
			&ReceiptItemDescriptionRule{},
		},
	}
}

func (s *RulesService) CalculatePoints(r *models.Receipt) (int, error) {

	totalPoints := 0

	for _, rule := range s.ReceiptRules {
		points, err := rule.Execute(r)
		if err != nil {
			return 0, fmt.Errorf("error executing rule: %w", err)
		}
		totalPoints += points
	}

	for _, item := range r.Items {
		for _, rule := range s.ReceiptItemRules {
			points, err := rule.Execute(&item)
			if err != nil {
				return 0, fmt.Errorf("error executing receipt item rule: %w", err)
			}
			totalPoints += points
		}
	}

	return totalPoints, nil

}
