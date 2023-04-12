import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken, checkSession } from './jwt.js';

async function fetchAllTransactions() {
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
        const idLink = document.createElement("a");
        idLink.textContent = transaction.id;
        idLink.href = `file-transaction.html?id=${transaction.id}`;
        idCell.appendChild(idLink);
        row.appendChild(idCell);

        tableBody.appendChild(row);
    });
}

document.addEventListener("DOMContentLoaded", async () => {
    checkSession();

    try {
        const transactions = await fetchAllTransactions();
        populateTransactionsTable(transactions);
    } catch (error) {
        console.error("Error fetching transactions:", error);
    }
});
