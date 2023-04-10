import { getJwtToken } from './jwt.js';

function getIdFromHash() {
    return window.location.hash.slice(1).split("/")[1];
}

async function fetchTransactions(id) {
    const jwtToken = getJwtToken();
    const response = await fetch(
        `http://localhost:3000/file-transactions/${id}/transactions`,
        {
            headers: {
                Authorization: `Bearer ${jwtToken}`,
            },
        }
    );

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

        const typeCell = document.createElement("td");
        typeCell.textContent = transaction.type;
        row.appendChild(typeCell);

        const dateCell = document.createElement("td");
        dateCell.textContent = transaction.date;
        row.appendChild(dateCell);

        const productIdCell = document.createElement("td");
        productIdCell.textContent = transaction.product_id;
        row.appendChild(productIdCell);

        const amountCell = document.createElement("td");
        amountCell.textContent = transaction.amount;
        row.appendChild(amountCell);

        const sellerIdCell = document.createElement("td");
        sellerIdCell.textContent = transaction.seller_id;
        row.appendChild(sellerIdCell);

        tableBody.appendChild(row);
    });
}

document.addEventListener("DOMContentLoaded", async () => {
    const id = getIdFromHash();
    try {
        const transactions = await fetchTransactions(id);
        populateTransactionsTable(transactions);
    } catch (error) {
        console.error("Error fetching transactions:", error);
    }
});
