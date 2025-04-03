package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"personal-finance-tracker/handlers"
	"personal-finance-tracker/models"
)

func main() {
	// Create a new finance manager
	manager := models.NewFinanceManager("transactions.json")

	// Parse templates
	templates := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	// Create a new handler
	handler := handlers.NewHandler(manager, templates)

	// Set up routes
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/add", handler.AddTransactionHandler)
	http.HandleFunc("/edit", handler.EditTransactionHandler)
	http.HandleFunc("/delete", handler.DeleteTransactionHandler)
	http.HandleFunc("/api/summary", handler.SummaryHandler)
	http.HandleFunc("/api/transactions", handler.APITransactionsHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
