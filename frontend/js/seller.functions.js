// functions.js
import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken, checkSession } from './jwt.js';

function getIdFromURL() {
    const queryParams = new URLSearchParams(window.location.search);
    return queryParams.get('id');
}

export async function fetchSellerBalance(id) {
    const jwtToken = getJwtToken();
    const response = await fetch(
        `http://localhost:3000/api/v1/seller/sellers/${id}/balance`,
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
