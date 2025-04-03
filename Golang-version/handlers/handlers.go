package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"personal-finance-tracker/models"
)

// Handler handles HTTP requests
type Handler struct {
	Manager   *models.FinanceManager
	Templates *template.Template
}

// NewHandler creates a new handler
func NewHandler(manager *models.FinanceManager, templates *template.Template) *Handler {
	return &Handler{
		Manager:   manager,
		Templates: templates,
	}
}

// IndexHandler handles the index page
func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	h.Templates.ExecuteTemplate(w, "index.html", h.Manager.Transactions)
}

// AddTransactionHandler handles adding a new transaction
func (h *Handler) AddTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	transactionType := r.FormValue("type")
	amountStr := r.FormValue("amount")
	description := r.FormValue("description")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	transaction := models.NewTransaction(h.Manager.GetNextID(), transactionType, amount, description)
	h.Manager.AddTransaction(transaction)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// EditTransactionHandler handles editing a transaction
func (h *Handler) EditTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.Templates.ExecuteTemplate(w, "edit.html", nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	transactionType := r.FormValue("type")
	amountStr := r.FormValue("amount")
	description := r.FormValue("description")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	h.Manager.EditTransaction(id, transactionType, amount, description)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteTransactionHandler handles deleting a transaction
func (h *Handler) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.Templates.ExecuteTemplate(w, "delete.html", nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	h.Manager.DeleteTransaction(id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SummaryHandler handles the transaction summary
func (h *Handler) SummaryHandler(w http.ResponseWriter, r *http.Request) {
	incomeTotal, expenseTotal, netFlow := h.Manager.GetTransactionSummary()

	summary := struct {
		IncomeTotal  float64
		ExpenseTotal float64
		NetFlow      float64
	}{
		IncomeTotal:  incomeTotal,
		ExpenseTotal: expenseTotal,
		NetFlow:      netFlow,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

// APITransactionsHandler handles API requests for transactions
func (h *Handler) APITransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Manager.Transactions)
}
