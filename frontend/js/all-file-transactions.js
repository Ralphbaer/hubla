import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken } from './jwt.js';

async function fetchTransactions() {
    const jwtToken = getJwtToken();
    const response = await fetch(
        `http://localhost:3000/api/v1/transaction/file-transactions/transactions`,
        {
            headers: {
                Authorization: `Bearer ${jwtToken}`,
            },
        }
    ).then(handleUnauthorized).then(handleErrors);

    if (!response.ok) {
        throw new Error("Error fetching transactions.");
    }

    return await response.json();
}

function populateTransactionsTable(transactions) {
    const tableBody = document.querySelector("#transactions-table tbody");

    transactions.forEach((transaction) => {
        const row = document.createElement("tr");

        const idCell = document.createElement("td");
        idCell.textContent = transaction.id;
        row.appendChild(idCell);

        const transactionIdCell = document.createElement("td");
        transactionIdCell.textContent = transaction.transaction_id;
        row.appendChild(transactionIdCell);

        const fileIdCell = document.createElement("td");
        const fileIdLink = document.createElement("a");
        fileIdLink.textContent = transaction.file_id;
        fileIdLink.href = `file-transaction.html?id=${transaction.file_id}`;
        fileIdCell.appendChild(fileIdLink);
        row.appendChild(fileIdCell);

        tableBody.appendChild(row);
    });
}

document.addEventListener("DOMContentLoaded", async () => {
    try {
        const transactions = await fetchTransactions();
        populateTransactionsTable(transactions);
    } catch (error) {
        console.error("Error fetching transactions:", error);
    }
});
