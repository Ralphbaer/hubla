import { setJwtToken } from './jwt.js';
import { showError, handleErrors } from './error.js';

document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('login-form');
    loginForm.addEventListener('submit', (event) => {
        event.preventDefault();

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        fetch('http://localhost:4000/api/v1/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password })
        })
            .then(handleErrors)
            .then(response => response.json())
            .then(data => {
                console.log(`Login successful: ${JSON.stringify(data)}`);

                const { access_token, status } = data;
                try {
                    setJwtToken(access_token);
                } catch (error) {
                    console.error('Error setting JWT token:', error);
                }

                window.location.href = 'http://127.0.0.1:5500/frontend/home.html';
            })
            .catch(error => {
                showError(error)
            });
    });
});