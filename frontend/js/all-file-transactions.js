import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken, checkSession } from './jwt.js';

const tableBody = document.querySelector("#transactions-table tbody");

const createTableRow = (transaction) => {
    const row = document.createElement("tr");

    const idCell = document.createElement("td");
    const idLink = document.createElement("a");
    idLink.textContent = transaction.id;
    idLink.href = `file-transaction.html?id=${transaction.id}`;
    idCell.appendChild(idLink);
    row.appendChild(idCell);

    return row;
}

const fetchAllTransactions = async () => {
    const jwtToken = getJwtToken();
    try {
        const response = await fetch(
            "http://localhost:3000/api/v1/transaction/file-transactions/transactions",
            {
                headers: {
                    Authorization: `Bearer ${jwtToken}`,
                },
            }
        ).then(handleUnauthorized).then(handleErrors);

        const transactions = await response.json();
        return transactions;
    } catch (error) {
        console.error("Error fetching transactions:", error);
    }
}

const populateTransactionsTable = (transactions) => {
    const rows = transactions.map(createTableRow);
    tableBody.append(...rows);
}

document.addEventListener("DOMContentLoaded", async () => {
    checkSession();

    const transactions = await fetchAllTransactions();
    populateTransactionsTable(transactions);
});
