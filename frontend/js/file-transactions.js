import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken } from './jwt.js';

function getIdFromURL() {
    const queryParams = new URLSearchParams(window.location.search);
    return queryParams.get('id');
}

async function fetchTransactions(id) {
    const jwtToken = getJwtToken();
    const response = await fetch(
        `http://localhost:3000/api/v1/transaction/file-transactions/${id}/transactions`,
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

        const typeCell = document.createElement("td");
        typeCell.textContent = transaction.t_type;
        row.appendChild(typeCell);

        const dateCell = document.createElement("td");
        dateCell.textContent = transaction.t_date;
        row.appendChild(dateCell);

        const productIdCell = document.createElement("td");
        productIdCell.textContent = transaction.product_id;
        row.appendChild(productIdCell);

        const amountCell = document.createElement("td");
        amountCell.textContent = transaction.amount;
        row.appendChild(amountCell);

        const sellerIdCell = document.createElement("td");
        const sellerLink = document.createElement("a");
        sellerLink.textContent = transaction.seller_id;
        sellerLink.href = `seller.html?id=${transaction.seller_id}`;
        sellerIdCell.appendChild(sellerLink);
        row.appendChild(sellerIdCell);

        tableBody.appendChild(row);
    });
}

document.addEventListener("DOMContentLoaded", async () => {
    const id = getIdFromURL();
    try {
        const transactions = await fetchTransactions(id);
        populateTransactionsTable(transactions);
    } catch (error) {
        console.error("Error fetching transactions:", error);
    }
});
