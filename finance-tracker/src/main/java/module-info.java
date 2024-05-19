module com.tracker {
    requires javafx.controls;
    requires javafx.fxml;
    requires java.sql;

    opens com.tracker to javafx.fxml;
    exports com.tracker;
}
