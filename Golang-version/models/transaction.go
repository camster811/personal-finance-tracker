// Package models contains the data structures and business logic for our application
package models

import (
	"fmt"
)

// Transaction represents a single financial transaction
// It could be either an income or an expense
type Transaction struct {
	ID          int     `json:"id"`          // Unique identifier for the transaction
	Type        string  `json:"type"`        // Type of transaction: "Income" or "Expense"
	Amount      float64 `json:"amount"`      // Amount of money involved in the transaction
	Description string  `json:"description"` // Description of what the transaction was for
}

// NewTransaction creates a new Transaction with the given parameters
// This is a constructor function that makes it easy to create new transactions
func NewTransaction(id int, transactionType string, amount float64, description string) *Transaction {
	// Create and return a new Transaction with the provided values
	return &Transaction{
		ID:          id,          // Unique identifier
		Type:        transactionType, // "Income" or "Expense"
		Amount:      amount,      // How much money was involved
		Description: description, // What the transaction was for
	}
}

// String returns a formatted string representation of the transaction
// This is useful for printing transactions in a readable format
// It implements the Stringer interface, so fmt.Println can print transactions nicely
func (t *Transaction) String() string {
	// Format the transaction as a multi-line string
	return fmt.Sprintf(
		"Transaction %d\nType: %s\nAmount: %.2f\nDescription: %s\n",
		t.ID,          // Transaction ID
		t.Type,        // Transaction type
		t.Amount,      // Amount (formatted to 2 decimal places)
		t.Description, // Description
	)
}
