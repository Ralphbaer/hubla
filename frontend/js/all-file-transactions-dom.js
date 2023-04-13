export const createTableRowElement = (transaction) => {
    const row = document.createElement("tr");

    const idCell = document.createElement("td");
    const idLink = document.createElement("a");
    idLink.textContent = transaction.id;
    idLink.href = `file-transaction.html?id=${transaction.id}`;
    idCell.appendChild(idLink);
    row.appendChild(idCell);

    return row;
}

export const updateTransactionsTable = (transactions) => {
    const tableBody = document.querySelector("#transactions-table tbody");
    const rows = transactions.map(createTableRowElement);
    tableBody.append(...rows);
}
