package receipt

import (
	"errors"
	"receipt_processor/internal/ledger"
	ledgermodels "receipt_processor/internal/models"
	"receipt_processor/internal/rules"
	"receipt_processor/internal/server/models"
)

type ReceiptService struct {
	rules  rules.RulesService
	ledger ledger.LedgerService
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{
		rules:  rules.RulesService{},
		ledger: ledger.New(),
	}
}

func (s *ReceiptService) LookupPoints(id string) (int, error) {

	if id == "" {
		return 0, errors.New("receipt ID cannot be empty")
	}

	entry, exists := s.ledger.Find(id)

	if !exists {
		return 0, errors.New("not found")
	}
	return entry.Points, nil
}

func (s *ReceiptService) ProcessReceipt(receipt *models.Receipt) (string, error) {

	if receipt == nil {
		return "", errors.New("receipt data cannot be nil")
	}

	points, err := s.rules.CalculatePoints(receipt)

	if err != nil {
		return "", err
	}

	entry, err := ledgermodels.CreateLedgerEntry(receipt, points)

	if err != nil {
		return "", errors.New("error creating ledger entry")
	}

	err = s.ledger.Insert(entry)

	if err != nil {
		return "", err
	} else {
		return entry.Id, nil
	}

}
