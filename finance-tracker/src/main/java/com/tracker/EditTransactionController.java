package com.tracker;

import javafx.fxml.FXML;
import javafx.scene.control.TextField;
import javafx.stage.Stage;

public class EditTransactionController {
    @FXML
    private TextField idField;
    @FXML
    private TextField typeField;
    @FXML
    private TextField amountField;
    @FXML
    private TextField descriptionField;

    private FinanceManager manager;

    public void setFinanceManager(FinanceManager manager) {
        this.manager = manager;
    }

    @FXML
    private void handleEditTransaction() {
        int id = Integer.parseInt(idField.getText());
        String type = typeField.getText();
        double amount = Double.parseDouble(amountField.getText());
        String description = descriptionField.getText();
        manager.editTransaction(id, type, amount, description);
        Stage stage = (Stage) idField.getScene().getWindow();
        stage.close();
    }
}
