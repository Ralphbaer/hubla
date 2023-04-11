export function getJwtToken() {
    try {
        return sessionStorage.getItem("jwt");
    } catch (error) {
        console.error('Error getting JWT token:', error);
        return null;
    }
}

export function setJwtToken(token) {
    try {
        sessionStorage.setItem("jwt", token);
    } catch (error) {
        console.error('Error setting JWT token:', error);
    }
}