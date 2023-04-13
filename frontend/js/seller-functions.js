import { handleErrors, handleUnauthorized } from './error.js';
import { getJwtToken } from './jwt.js';

export const fetchSellerBalance = async (id) => {
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
