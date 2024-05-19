package com.tracker;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

import javafx.scene.control.Alert;

public class FinanceManager {
    private List<Transaction> transactions = new ArrayList<>();
    private final String filePath = "transactions.dat";

    public FinanceManager() {
        loadTransactions();
    }

    public void addTransaction(Transaction transaction) {
        transactions.add(transaction);
        saveTransactions();
    }

    public void editTransaction(int id, String type, double amount, String description) {
        for (Transaction transaction : transactions) {
            if (transaction.getId() == id) {
                transaction.setType(type);
                transaction.setAmount(amount);
                transaction.setDescription(description);
                saveTransactions();
            }
        }
    }

    public void deleteTransaction(int id) {
        transactions.removeIf(transaction -> transaction.getId() == id);
        saveTransactions();
    }

    public void listTransactions() {
        transactions.forEach(System.out::println);
    }

    private void saveTransactions() {
        try (ObjectOutputStream out = new ObjectOutputStream(new FileOutputStream(filePath))) {
            out.writeObject(transactions);
        } catch (IOException error) {
            System.out.println(error);
        }
    }

    public void loadTransactions() {
        try (ObjectInputStream in = new ObjectInputStream(new FileInputStream(filePath))) {
            transactions = (List<Transaction>) in.readObject();
        } catch (FileNotFoundException e) {
            Alert alert = new Alert(Alert.AlertType.INFORMATION);
            alert.setTitle("No file found");
            alert.setHeaderText(null);
            alert.setContentText("No data found, starting fresh");
            alert.showAndWait();
        } catch (IOException | ClassNotFoundException e) {
            Alert alert = new Alert(Alert.AlertType.INFORMATION);
            alert.setTitle("Error");
            alert.setHeaderText(null);
            alert.setContentText("{e}");
            alert.showAndWait();
        }
    }

    public int getNextID() {
        return transactions.isEmpty() ? 1 : transactions.get(transactions.size() - 1).getId() + 1;
    }

    public List<Transaction> getTransactions() {
        return transactions;
    }
}
