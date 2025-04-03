# Personal Finance Tracker - Golang Version

This is a Golang implementation of the Personal Finance Tracker application, converted from the original Java/JavaFX version.

## Features

- Add income and expense transactions
- Edit existing transactions
- Delete transactions
- View a summary of all transactions (total income, expenses, and net flow)
- View all transactions in a table

## Project Structure

- `models/`: Contains the data models and business logic
  - `transaction.go`: Defines the Transaction struct and related methods
  - `finance_manager.go`: Manages transactions and file operations
- `handlers/`: Contains HTTP request handlers
  - `handlers.go`: Defines handlers for different routes
- `templates/`: Contains HTML templates
  - `index.html`: Main page template
  - `edit.html`: Edit transaction page template
  - `delete.html`: Delete transaction page template
- `static/`: Contains static assets
  - `style.css`: CSS styles
  - `script.js`: Client-side JavaScript
- `main.go`: Application entry point

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to the `Golang-version` directory
3. Run the application:
   ```
   go run main.go
   ```
4. Open a web browser and go to `http://localhost:8080`

## Data Storage

Transactions are stored in a JSON file (`transactions.json`) in the root directory. The file is created automatically if it doesn't exist.

## Differences from the Java Version

- The Java version uses JavaFX for the UI, while the Golang version uses a web-based UI with HTML, CSS, and JavaScript
- The Java version uses Java serialization for data storage, while the Golang version uses JSON
- The Golang version includes a RESTful API for accessing transaction data and summaries
