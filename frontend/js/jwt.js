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

function isJwtExpired() {
    const payload = parseJwt(getJwtToken());

    if (!payload.exp) {
        return false;
    }

    const currentTime = Math.floor(Date.now() / 1000);

    return payload.exp < currentTime;
}


function parseJwt(token) {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace('-', '+').replace('_', '/');
    return JSON.parse(atob(base64));
}


export function checkSession() {
    const token = getJwtToken()
    if (token) {
        if (isJwtExpired(token)) {
            logout();
            return false;
        } else {
            return true;
        }
    } else {
        logout()
        return false;
    }
}

function logout() {
    localStorage.removeItem('jwt');
    window.location.href = '/frontend/login.html';
}


