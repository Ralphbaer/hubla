import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken, checkSession } from './jwt.js';
import { fetchSellerBalance } from './seller.functions.js';

function getIdFromURL() {
    const queryParams = new URLSearchParams(window.location.search);
    return queryParams.get('id');
}

async function fetchTransactions(id) {
    const jwtToken = getJwtToken();
    const response = await fetch(`http://localhost:3000/api/v1/transaction/file-transactions/${id}/transactions`, {
        headers: {
            Authorization: `Bearer ${jwtToken}`,
        },
    }).then(handleUnauthorized).then(handleErrors);

    if (!response.ok) {
        throw new Error("Error fetching transactions.");
    }

    return await response.json();
}

function createTableCell(value) {
    const cell = document.createElement("td");
    cell.textContent = value;
    return cell;
}

function createLinkCell(value, href) {
    const cell = document.createElement("td");
    const link = document.createElement("a");
    link.textContent = value;
    link.href = href;
    cell.appendChild(link);
    return cell;
}

async function populateSellerBalanceRow(transaction) {
    if (!populateSellerBalanceRow.uniqueIds) {
        populateSellerBalanceRow.uniqueIds = [];
    }

    const sellerId = transaction.seller_id;
    if (populateSellerBalanceRow.uniqueIds.includes(sellerId)) {
        return null;
    } else {
        populateSellerBalanceRow.uniqueIds.push(sellerId);
    }

    const sellerBalance = await fetchSellerBalance(sellerId);
    const row = document.createElement("tr");
    row.appendChild(createTableCell(sellerBalance.seller_name));
    row.appendChild(createTableCell(formatCurrency(sellerBalance.seller_balance)));
    return row;
}


async function populateTransactionsTable(transactions) {
    const tableBody = document.querySelector("#transactions-table tbody");

    transactions.forEach((transaction) => {
        const row = document.createElement("tr");
        row.appendChild(createTableCell(transaction.id));
        row.appendChild(createTableCell(transaction.t_type));
        row.appendChild(createTableCell(transaction.t_date));
        row.appendChild(createTableCell(transaction.product_id));
        row.appendChild(createTableCell(formatCurrency(transaction.amount)));
        row.appendChild(createLinkCell(transaction.seller_id, `seller.html?id=${transaction.seller_id}`));
        tableBody.appendChild(row);
    });
}

async function populateSellerBalanceTable(transactions) {
    const tableBody = document.querySelector("#sellers-table tbody");

    const rows = await Promise.all(transactions.map(populateSellerBalanceRow));

    rows
        .filter(row => row !== null)
        .forEach(row => tableBody.appendChild(row));
}


document.addEventListener("DOMContentLoaded", async () => {
    checkSession();

    const id = getIdFromURL();
    try {
        const transactions = await fetchTransactions(id);
        await populateTransactionsTable(transactions);
        await populateSellerBalanceTable(transactions);
    } catch (error) {
        console.error("Error fetching transactions:", error);
    }
});

function formatCurrency(amount) {
    const formatter = new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL'
    });
    return formatter.format(amount).replace(/\s/, ''); // remove whitespace between symbol and amount
}
