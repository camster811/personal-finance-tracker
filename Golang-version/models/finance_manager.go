package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// FinanceManager manages financial transactions
type FinanceManager struct {
	Transactions []*Transaction `json:"transactions"`
	filePath     string
	mutex        sync.Mutex
}

// NewFinanceManager creates a new finance manager
func NewFinanceManager(filePath string) *FinanceManager {
	manager := &FinanceManager{
		Transactions: []*Transaction{},
		filePath:     filePath,
	}
	manager.LoadTransactions()
	return manager
}

// AddTransaction adds a new transaction
func (fm *FinanceManager) AddTransaction(transaction *Transaction) {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	fm.Transactions = append(fm.Transactions, transaction)
	fm.SaveTransactions()
}

// EditTransaction edits an existing transaction
func (fm *FinanceManager) EditTransaction(id int, transactionType string, amount float64, description string) {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	for _, transaction := range fm.Transactions {
		if transaction.ID == id {
			transaction.Type = transactionType
			transaction.Amount = amount
			transaction.Description = description
			fm.SaveTransactions()
			return
		}
	}
}

// DeleteTransaction deletes a transaction by ID
func (fm *FinanceManager) DeleteTransaction(id int) {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	for i, transaction := range fm.Transactions {
		if transaction.ID == id {
			// Remove the transaction from the slice
			fm.Transactions = append(fm.Transactions[:i], fm.Transactions[i+1:]...)
			fm.SaveTransactions()
			return
		}
	}
}

// ListTransactions prints all transactions
func (fm *FinanceManager) ListTransactions() {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	for _, transaction := range fm.Transactions {
		fmt.Println(transaction)
	}
}

// GetNextID returns the next available ID for a new transaction
func (fm *FinanceManager) GetNextID() int {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	if len(fm.Transactions) == 0 {
		return 1
	}
	return fm.Transactions[len(fm.Transactions)-1].ID + 1
}

// SaveTransactions saves all transactions to a file
func (fm *FinanceManager) SaveTransactions() error {
	data, err := json.MarshalIndent(fm.Transactions, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling transactions: %v", err)
	}

	err = os.WriteFile(fm.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing transactions to file: %v", err)
	}

	return nil
}

// LoadTransactions loads transactions from a file
func (fm *FinanceManager) LoadTransactions() error {
	// Check if file exists
	if _, err := os.Stat(fm.filePath); os.IsNotExist(err) {
		// Create directory if it doesn't exist
		dir := filepath.Dir(fm.filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
		
		// Create empty file
		if err := fm.SaveTransactions(); err != nil {
			return fmt.Errorf("error creating empty transactions file: %v", err)
		}
		
		return nil
	}

	data, err := os.ReadFile(fm.filePath)
	if err != nil {
		return fmt.Errorf("error reading transactions file: %v", err)
	}

	if len(data) == 0 {
		fm.Transactions = []*Transaction{}
		return nil
	}

	err = json.Unmarshal(data, &fm.Transactions)
	if err != nil {
		return fmt.Errorf("error unmarshaling transactions: %v", err)
	}

	return nil
}

// GetTransactionSummary returns a summary of all transactions
func (fm *FinanceManager) GetTransactionSummary() (float64, float64, float64) {
	fm.mutex.Lock()
	defer fm.mutex.Unlock()
	
	var incomeTotal, expenseTotal float64
	
	for _, transaction := range fm.Transactions {
		if transaction.Type == "Income" {
			incomeTotal += transaction.Amount
		} else if transaction.Type == "Expense" {
			expenseTotal += transaction.Amount
		}
	}
	
	return incomeTotal, expenseTotal, incomeTotal - expenseTotal
}
