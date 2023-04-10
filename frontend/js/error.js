export function showError(message) {
    const toast = document.getElementById('toast');
    toast.innerText = message;
    toast.classList.add('show');
    setTimeout(() => {
        toast.classList.remove('show');
    }, 3000);
}

export async function handleErrors(response) {
    if (!response.ok) {
        const errorData = await response.json();
        const errorMessage =
            errorData.code !== 500
                ? errorData.message
                : 'Internal server error';

        throw new Error(errorMessage);
    }

    return response;
}