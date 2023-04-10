$(document).ready(function () {
    $('#login-form').submit(function (event) {
        event.preventDefault(); // prevent the default form submission

        // get the email and password from the form
        const email = $('#email').val();
        const password = $('#password').val();

        // make the AJAX request
        $.ajax({
            url: '/auth/login', // modify the URL to remove the explicit file reference
            method: 'POST',
            data: JSON.stringify({ email, password }),
            contentType: 'application/json'
        })
            .done(function (response, textStatus, jqXHR) {
                console.log('Login successful:', response);

                // extract the access_token and status from the response
                const { access_token, status } = response;
                console.log('access_token:', access_token);
                console.log('status:', status);

                // extract the cookies from the response
                const cookiesHeader = jqXHR.getResponseHeader('Set-Cookie');
                const cookies = cookiesHeader.match(/(\w+)=(.+?)(;|$)/g).map(cookie => cookie.split('=')[1]);
                console.log('cookies:', cookies);

                // do something with the successful response
            })
            .fail(function (jqXHR, textStatus, errorThrown) {
                console.error('Login failed:', textStatus, errorThrown);
                // do something with the failed response
            });
    });
});
