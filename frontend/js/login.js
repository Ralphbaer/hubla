$(document).ready(function () {
    $('#login-form').submit(function (event) {
        event.preventDefault(); // prevent the default form submission

        // get the email and password from the form
        const email = $('#email').val();
        const password = $('#password').val();

        // make the AJAX request
        $.ajax({
            url: 'http://localhost:4000/auth/login', // modify the URL to remove the explicit file reference
            method: 'POST',
            data: JSON.stringify({ email, password }),
            contentType: 'application/json',
        })
            .done(function (response, textStatus, jqXHR) {
                console.log('Login successful:', response);

                const { access_token, status } = response;
                console.log('access_token:', access_token);
                console.log('status:', status);

                // set the JWT token in sessionStorage
                try {
                    setJwtToken(access_token);
                } catch (error) {
                    console.error('Error setting JWT token:', error);
                }

                window.location.href = 'http://127.0.0.1:5500/frontend/home.html';
            })
            .fail(function (jqXHR, textStatus, errorThrown) {
                console.error('Login failed:', textStatus, errorThrown);
            });
    });
});

// Short duration JWT token (5-10 min)
function getJwtToken() {
    try {
        return sessionStorage.getItem("jwt");
    } catch (error) {
        console.error('Error getting JWT token:', error);
        return null;
    }
}

function setJwtToken(token) {
    try {
        sessionStorage.setItem("jwt", token);
    } catch (error) {
        console.error('Error setting JWT token:', error);
    }
}
