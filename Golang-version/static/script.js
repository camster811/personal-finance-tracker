document.addEventListener("DOMContentLoaded", function () {
  // Get the summarize button and summary div
  const summarizeBtn = document.getElementById("summarize-btn");
  const summaryDiv = document.getElementById("summary");
  const incomeTotal = document.getElementById("income-total");
  const expenseTotal = document.getElementById("expense-total");
  const netFlow = document.getElementById("net-flow");

  // Add click event listener to the summarize button
  summarizeBtn.addEventListener("click", function () {
    // Fetch the summary data from the API
    fetch("/api/summary")
      .then((response) => response.json())
      .then((data) => {
        // Update the summary div with the data
        incomeTotal.textContent = data.IncomeTotal.toFixed(2);
        expenseTotal.textContent = data.ExpenseTotal.toFixed(2);
        netFlow.textContent = data.NetFlow.toFixed(2);

        // Show the summary div
        summaryDiv.classList.remove("hidden");
      })
      .catch((error) => {
        console.error("Error fetching summary:", error);
        alert("Error fetching summary. Please try again.");
      });
  });
});
