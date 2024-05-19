package com.tracker;

import javafx.fxml.FXML;
import javafx.scene.control.TextField;
import javafx.stage.Stage;

public class DeleteTransactionController {
    @FXML
    private TextField idField;

    private FinanceManager manager;

    public void setFinanceManager(FinanceManager manager) {
        this.manager = manager;
    }

    @FXML
    private void handleDeleteTransaction() {
        int id = Integer.parseInt(idField.getText());
        manager.deleteTransaction(id);
        Stage stage = (Stage) idField.getScene().getWindow();
        stage.close();
    }
}
