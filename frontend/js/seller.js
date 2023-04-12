import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken, checkSession } from './jwt.js';
import { fetchSellerBalance } from './seller.functions.js';

function getIdFromURL() {
    const queryParams = new URLSearchParams(window.location.search);
    return queryParams.get('id');
}

function populateSellerBalanceContainer(sellerBalance) {
    const idElement = document.querySelector('.seller-id');
    idElement.textContent = `ID: ${sellerBalance.seller_id}`;

    const typeElement = document.querySelector('.seller-type');
    typeElement.textContent = `Tipo: ${sellerBalance.seller_type}`;

    const nameElement = document.querySelector('.seller-name');
    nameElement.textContent = `Nome: ${sellerBalance.seller_name}`;

    const balanceElement = document.querySelector('.seller-balance');
    balanceElement.textContent = `Balanço (total em vendas): R$ ${sellerBalance.seller_balance}`;

    const balanceUpdatedAtElement = document.querySelector('.seller-balance-updated-at');
    balanceUpdatedAtElement.textContent = `Última atualização do balanço: ${sellerBalance.seller_balance_updated_at}`;

    const createdAtElement = document.querySelector('.seller-created-at');
    createdAtElement.textContent = `Data de criação: ${sellerBalance.seller_created_at}`;
}

document.addEventListener("DOMContentLoaded", async () => {
    checkSession();

    const id = getIdFromURL();
    try {
        const sellerBalance = await fetchSellerBalance(id);

        console.log("asdas" + JSON.stringify(sellerBalance))
        populateSellerBalanceContainer(sellerBalance);
    } catch (error) {
        console.error("Error fetching seller balance information:", error);
    }
});
