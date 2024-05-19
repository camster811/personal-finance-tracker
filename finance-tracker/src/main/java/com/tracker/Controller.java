package com.tracker;

import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.control.cell.PropertyValueFactory;
import javafx.stage.Modality;
import javafx.stage.Stage;

import java.io.IOException;
import java.text.DecimalFormat;

public class Controller {
    @FXML
    private TextField typeField;
    @FXML
    private TextField amountField;
    @FXML
    private TextField descriptionField;
    @FXML
    private TableView<Transaction> transactionTable;
    @FXML
    private TableColumn<Transaction, Number> idColumn;
    @FXML
    private TableColumn<Transaction, String> typeColumn;
    @FXML
    private TableColumn<Transaction, Number> amountColumn;
    @FXML
    private TableColumn<Transaction, String> descriptionColumn;

    private FinanceManager manager = new FinanceManager();
    private ObservableList<Transaction> transactionList = FXCollections.observableArrayList();

    @FXML
    public void initialize() {
        idColumn.setCellValueFactory(new PropertyValueFactory<>("id"));
        typeColumn.setCellValueFactory(new PropertyValueFactory<>("type"));
        amountColumn.setCellValueFactory(new PropertyValueFactory<>("amount"));
        descriptionColumn.setCellValueFactory(new PropertyValueFactory<>("description"));

        manager.loadTransactions();
        transactionTable.getItems().addAll(manager.getTransactions());
    }

    @FXML
    private void handleAddTransaction() {
        String type = typeField.getText();
        double amount = Double.parseDouble(amountField.getText());
        String description = descriptionField.getText();
        Transaction transaction = new Transaction(manager.getNextID(), type, amount, description);
        manager.addTransaction(transaction);
        transactionList.add(transaction);
        refreshTable();
    }

    @FXML
    private void openEditTransactionWindow() {
        try {
            FXMLLoader loader = new FXMLLoader(getClass().getResource("EditUI.fxml"));
            Parent root = loader.load();

            EditTransactionController controller = loader.getController();
            controller.setFinanceManager(manager);

            Stage stage = new Stage();
            stage.setTitle("Edit Transaction");
            stage.initModality(Modality.WINDOW_MODAL);
            stage.setScene(new Scene(root));
            stage.showAndWait();

            refreshTable();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @FXML
    private void openDeleteTransactionWindow() {
        try {
            FXMLLoader loader = new FXMLLoader(getClass().getResource("DeleteUI.fxml"));
            Parent root = loader.load();

            DeleteTransactionController controller = loader.getController();
            controller.setFinanceManager(manager);

            Stage stage = new Stage();
            stage.setTitle("Delete Transaction");
            stage.initModality(Modality.WINDOW_MODAL);
            stage.setScene(new Scene(root));
            stage.showAndWait();

            transactionList.setAll(manager.getTransactions());
            refreshTable();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @FXML
    private void handleSummarizeTransactions() {
        DecimalFormat df = new DecimalFormat("##.00");
        double incomeTotal = manager.getTransactions().stream()
                .filter(t -> t.getType().equalsIgnoreCase("Income"))
                .mapToDouble(Transaction::getAmount)
                .sum() * 100 / 100;

        double expenseTotal = manager.getTransactions().stream()
                .filter(t -> t.getType().equalsIgnoreCase("Expense"))
                .mapToDouble(Transaction::getAmount)
                .sum() * 100 / 100;

        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("Transaction Summary");
        alert.setHeaderText(null);
        alert.setContentText("Total income: " + incomeTotal + "\nTotal expenses: " + expenseTotal + "\nNet flow: " + (df.format(incomeTotal - expenseTotal)));
        alert.showAndWait();
    }

    private void refreshTable() {
        // Clear and re-populate the table with updated transactions
        transactionTable.getItems().clear();
        transactionTable.getItems().addAll(manager.getTransactions());
    }
}
