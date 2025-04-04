// Package models contains the data structures and business logic for our application
package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// FinanceManager is responsible for managing all financial transactions
// It handles loading, saving, adding, editing, and deleting transactions
type FinanceManager struct {
	Transactions []*Transaction `json:"transactions"` // List of all transactions
	filePath     string         // Path to the JSON file where transactions are stored
	mutex        sync.Mutex     // Mutex to prevent concurrent access to transactions
}

// NewFinanceManager creates a new finance manager with the given file path
// The file path is where transactions will be saved to and loaded from
func NewFinanceManager(filePath string) *FinanceManager {
	// Create a new finance manager with an empty transactions list
	manager := &FinanceManager{
		Transactions: []*Transaction{}, // Start with an empty list
		filePath:     filePath,         // Set the file path
	}
	
	// Load existing transactions from the file (if any)
	manager.LoadTransactions()
	
	// Return the initialized manager
	return manager
}

// AddTransaction adds a new transaction to the manager
// It also saves the updated transactions list to the file
func (fm *FinanceManager) AddTransaction(transaction *Transaction) {
	// Lock the mutex to prevent concurrent access
	// This ensures that only one goroutine can modify the transactions at a time
	fm.mutex.Lock()
	
	// Make sure to unlock the mutex when this function returns
	// defer ensures this happens even if an error occurs
	defer fm.mutex.Unlock()
	
	// Add the new transaction to the list
	fm.Transactions = append(fm.Transactions, transaction)
	
	// Save the updated transactions list to the file
	fm.SaveTransactions()
}

// EditTransaction modifies an existing transaction with the given ID
// It updates the transaction type, amount, and description
func (fm *FinanceManager) EditTransaction(id int, transactionType string, amount float64, description string) {
	// Lock the mutex to prevent concurrent access
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	// Loop through all transactions to find the one with the matching ID
	for _, transaction := range fm.Transactions {
		if transaction.ID == id {
			// Update the transaction fields
			transaction.Type = transactionType
			transaction.Amount = amount
			transaction.Description = description
			
			// Save the updated transactions list to the file
			fm.SaveTransactions()
			return
		}
	}
	// If we get here, no transaction with the given ID was found
}

// DeleteTransaction removes a transaction with the given ID
// It also saves the updated transactions list to the file
func (fm *FinanceManager) DeleteTransaction(id int) {
	// Lock the mutex to prevent concurrent access
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	// Loop through all transactions to find the one with the matching ID
	for i, transaction := range fm.Transactions {
		if transaction.ID == id {
			// Remove the transaction from the slice
			// This is done by creating a new slice that includes all elements
			// before the transaction to delete and all elements after it
			fm.Transactions = append(fm.Transactions[:i], fm.Transactions[i+1:]...)
			
			// Save the updated transactions list to the file
			fm.SaveTransactions()
			return
		}
	}
	// If we get here, no transaction with the given ID was found
}

// ListTransactions prints all transactions to the console
// This is useful for debugging purposes
func (fm *FinanceManager) ListTransactions() {
	// Lock the mutex to prevent concurrent access
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	// Loop through all transactions and print each one
	for _, transaction := range fm.Transactions {
		fmt.Println(transaction)
	}
}

// GetNextID returns the next available ID for a new transaction
// IDs are sequential, so this returns the highest existing ID + 1
func (fm *FinanceManager) GetNextID() int {
	// Lock the mutex to prevent concurrent access
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	// If there are no transactions, start with ID 1
	if len(fm.Transactions) == 0 {
		return 1
	}
	
	// Otherwise, return the highest existing ID + 1
	// We assume transactions are ordered by ID, so the last one has the highest ID
	return fm.Transactions[len(fm.Transactions)-1].ID + 1
}

// SaveTransactions saves all transactions to the JSON file
// It returns an error if the save operation fails
func (fm *FinanceManager) SaveTransactions() error {
	// Convert the transactions list to JSON
	// MarshalIndent creates formatted JSON with indentation for readability
	data, err := json.MarshalIndent(fm.Transactions, "", "  ")
	if err != nil {
		// If the conversion fails, return an error
		return fmt.Errorf("error marshaling transactions: %v", err)
	}

	// Write the JSON data to the file
	// 0644 is the file permission (readable by everyone, writable by owner)
	err = os.WriteFile(fm.filePath, data, 0644)
	if err != nil {
		// If the write fails, return an error
		return fmt.Errorf("error writing transactions to file: %v", err)
	}

	// If everything succeeded, return nil (no error)
	return nil
}

// LoadTransactions loads transactions from the JSON file
// It returns an error if the load operation fails
func (fm *FinanceManager) LoadTransactions() error {
	// Check if the file exists
	if _, err := os.Stat(fm.filePath); os.IsNotExist(err) {
		// If the file doesn't exist, create the directory structure
		dir := filepath.Dir(fm.filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
		
		// Create an empty file by saving the current (empty) transactions list
		if err := fm.SaveTransactions(); err != nil {
			return fmt.Errorf("error creating empty transactions file: %v", err)
		}
		
		// Return nil since we've successfully created an empty file
		return nil
	}

	// Read the file contents
	data, err := os.ReadFile(fm.filePath)
	if err != nil {
		return fmt.Errorf("error reading transactions file: %v", err)
	}

	// If the file is empty, initialize with an empty transactions list
	if len(data) == 0 {
		fm.Transactions = []*Transaction{}
		return nil
	}

	// Convert the JSON data to a transactions list
	err = json.Unmarshal(data, &fm.Transactions)
	if err != nil {
		return fmt.Errorf("error unmarshaling transactions: %v", err)
	}

	// If everything succeeded, return nil (no error)
	return nil
}

// GetTransactionSummary calculates and returns a summary of all transactions
// It returns the total income, total expenses, and net flow (income - expenses)
func (fm *FinanceManager) GetTransactionSummary() (float64, float64, float64) {
	// Lock the mutex to prevent concurrent access
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	// Initialize variables to track totals
	var incomeTotal float64 = 0  // Total of all income transactions
	var expenseTotal float64 = 0 // Total of all expense transactions
	
	// Loop through all transactions
	for _, transaction := range fm.Transactions {
		// Add to the appropriate total based on transaction type
		if transaction.Type == "Income" {
			incomeTotal += transaction.Amount
		} else if transaction.Type == "Expense" {
			expenseTotal += transaction.Amount
		}
		// Note: If the transaction type is neither "Income" nor "Expense",
		// it won't be included in either total
	}
	
	// Calculate the net flow (income - expenses)
	netFlow := incomeTotal - expenseTotal
	
	// Return all three values
	return incomeTotal, expenseTotal, netFlow
}
