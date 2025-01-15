package util

import (
	"encoding/json"
	"fmt"
	"os"
	"receipt_processor/internal/server/models"
)

func LoadTestReceipt(filename string) (*models.Receipt, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var r models.Receipt
	err = json.NewDecoder(file).Decode(&r)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &r, nil

}
