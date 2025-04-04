// Package handlers contains all the HTTP request handlers for our application
package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"personal-finance-tracker/models"
)

// Handler is a struct that processes HTTP requests
// It contains references to our finance manager and HTML templates
type Handler struct {
	Manager   *models.FinanceManager // Manages all our financial transactions
	Templates *template.Template     // Contains all our HTML templates
}

// NewHandler creates a new Handler with the given finance manager and templates
// This is called when our application starts up
func NewHandler(manager *models.FinanceManager, templates *template.Template) *Handler {
	return &Handler{
		Manager:   manager,
		Templates: templates,
	}
}

// IndexHandler displays the main page of our application
// It shows a list of all transactions
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Execute the index.html template and pass in all transactions as data
	// The template will use this data to display the transactions
	h.Templates.ExecuteTemplate(w, "index.html", h.Manager.Transactions)
}

// AddTransactionHandler processes requests to add a new transaction
// It handles the form submission when a user adds a transaction
func (h *Handler) AddTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a POST request (form submission)
	// We only want to process POST requests for this endpoint
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the form values submitted by the user
	transactionType := r.FormValue("type")       // "Income" or "Expense"
	amountStr := r.FormValue("amount")           // Amount as a string
	description := r.FormValue("description")    // Description of the transaction

	// Convert the amount from string to float64
	// This is necessary because form values are always strings
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		// If the conversion fails, return an error to the user
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Create a new transaction with the form values
	// GetNextID() generates a unique ID for the transaction
	transaction := models.NewTransaction(h.Manager.GetNextID(), transactionType, amount, description)
	
	// Add the transaction to our finance manager
	// This will also save it to the transactions.json file
	h.Manager.AddTransaction(transaction)

	// Redirect the user back to the home page
	// StatusSeeOther (303) is the standard redirect status code for POST requests
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// EditTransactionHandler processes requests to edit an existing transaction
// It handles both displaying the edit form and processing form submissions
func (h *Handler) EditTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// If it's a GET request, just display the edit form
	if r.Method == http.MethodGet {
		h.Templates.ExecuteTemplate(w, "edit.html", nil)
		return
	}

	// If it's not a POST request, return an error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the form values submitted by the user
	idStr := r.FormValue("id")                   // Transaction ID as a string
	transactionType := r.FormValue("type")       // "Income" or "Expense"
	amountStr := r.FormValue("amount")           // Amount as a string
	description := r.FormValue("description")    // Description of the transaction

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If the conversion fails, return an error to the user
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Convert the amount from string to float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		// If the conversion fails, return an error to the user
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Edit the transaction with the given ID
	// This will also save the changes to the transactions.json file
	h.Manager.EditTransaction(id, transactionType, amount, description)

	// Redirect the user back to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteTransactionHandler processes requests to delete a transaction
// It handles both displaying the delete confirmation form and processing form submissions
func (h *Handler) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// If it's a GET request, just display the delete confirmation form
	if r.Method == http.MethodGet {
		h.Templates.ExecuteTemplate(w, "delete.html", nil)
		return
	}

	// If it's not a POST request, return an error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the transaction ID from the form
	idStr := r.FormValue("id")

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If the conversion fails, return an error to the user
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the transaction with the given ID
	// This will also save the changes to the transactions.json file
	h.Manager.DeleteTransaction(id)

	// Redirect the user back to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SummaryHandler returns a JSON summary of all transactions
// This is used by the frontend to display summary information
func (h *Handler) SummaryHandler(w http.ResponseWriter, r *http.Request) {
	// Get the transaction summary from the finance manager
	// This returns the total income, total expenses, and net flow (income - expenses)
	incomeTotal, expenseTotal, netFlow := h.Manager.GetTransactionSummary()

	// Create a struct to hold the summary data
	// This will be converted to JSON and sent to the client
	summary := struct {
		IncomeTotal  float64 // Total of all income transactions
		ExpenseTotal float64 // Total of all expense transactions
		NetFlow      float64 // Income total minus expense total
	}{
		IncomeTotal:  incomeTotal,
		ExpenseTotal: expenseTotal,
		NetFlow:      netFlow,
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Encode the summary struct as JSON and write it to the response
	json.NewEncoder(w).Encode(summary)
}

// APITransactionsHandler returns all transactions as JSON
// This is used by the frontend to get all transaction data
func (h *Handler) APITransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Encode all transactions as JSON and write them to the response
	json.NewEncoder(w).Encode(h.Manager.Transactions)
}
