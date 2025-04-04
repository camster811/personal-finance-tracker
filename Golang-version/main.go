// Package main is the entry point for our personal finance tracker application
package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"personal-finance-tracker/handlers"
	"personal-finance-tracker/models"
)

// main is the starting point of our application
func main() {
	// Create a new finance manager that will handle our transactions
	// The "transactions.json" file is where all our data will be stored
	financeManager := models.NewFinanceManager("transactions.json")

	// Load all HTML templates from the templates directory
	// template.Must will panic if there's an error loading templates
	// This is okay for startup since we can't run without templates
	htmlTemplates := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	// Create a new HTTP handler that will process web requests
	// We pass in our finance manager and templates so the handler can use them
	requestHandler := handlers.NewHandler(financeManager, htmlTemplates)

	// Set up URL routes - each URL path is connected to a specific handler function
	// These functions will be called when a user visits the corresponding URL
	http.HandleFunc("/", requestHandler.IndexHandler)                 // Home page
	http.HandleFunc("/add", requestHandler.AddTransactionHandler)     // Add transaction page
	http.HandleFunc("/edit", requestHandler.EditTransactionHandler)   // Edit transaction page
	http.HandleFunc("/delete", requestHandler.DeleteTransactionHandler) // Delete transaction page
	http.HandleFunc("/api/summary", requestHandler.SummaryHandler)    // API endpoint for summary data
	http.HandleFunc("/api/transactions", requestHandler.APITransactionsHandler) // API endpoint for all transactions

	// Set up a file server to serve static files (CSS, JavaScript, images)
	// These files are stored in the "static" directory
	staticFileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	// Start the web server on port 8080
	// Log a message so we know the server has started
	log.Println("Server starting on http://localhost:8080")
	
	// ListenAndServe starts the server and blocks until the server is shut down
	// If there's an error starting the server, log.Fatal will log the error and exit
	log.Fatal(http.ListenAndServe(":8080", nil))
}
