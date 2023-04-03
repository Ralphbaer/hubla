const loginForm = document.getElementById("login-form");

loginForm.addEventListener("submit", (e) => {
    e.preventDefault();
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    // Authenticate the user using your internal process.
    // Replace this with your actual authentication process.
    if (email === "user@example.com" && password === "password123") {
        alert("Logged in successfully!");
        // Redirect to your internal process page.
        // window.location.href = "/your-internal-process-page";
    } else {
        alert("Invalid email or password. Please try again.");
    }
});
