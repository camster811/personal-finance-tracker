package models

import (
	"fmt"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID          int     `json:"id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// NewTransaction creates a new transaction with the given parameters
func NewTransaction(id int, transactionType string, amount float64, description string) *Transaction {
	return &Transaction{
		ID:          id,
		Type:        transactionType,
		Amount:      amount,
		Description: description,
	}
}

// String returns a string representation of the transaction
func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction %d\nType: %s\nAmount: %.2f\nDescription: %s\n",
		t.ID, t.Type, t.Amount, t.Description)
}
