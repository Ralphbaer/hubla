import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken, checkSession } from './jwt.js';
import { updateTransactionsTable } from './all-file-transactions-dom.js';

export const fetchAllTransactions = async () => {
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

document.addEventListener("DOMContentLoaded", async () => {
    checkSession();

    const transactions = await fetchAllTransactions();
    updateTransactionsTable(transactions);
});
