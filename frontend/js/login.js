import { checkSession, setJwtToken } from './jwt.js';
import { showError, handleErrors, } from './error.js';
import { redirectToHome } from './redirect.js';


const loginForm = document.getElementById('login-form');
document.addEventListener('DOMContentLoaded', () => {
    if (checkSession()) {
        redirectToHome();
    }

    loginForm.addEventListener('submit', handleLoginSubmit);
});

async function handleLoginSubmit(event) {
    event.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:4000/api/v1/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password })
        });

        const data = await handleErrors(response).json();
        console.log(`Login successful: ${JSON.stringify(data)}`);

        const { access_token } = data;
        setJwtToken(access_token);

        redirectToHome();
    } catch (error) {
        showError(error.message);
    }
}

