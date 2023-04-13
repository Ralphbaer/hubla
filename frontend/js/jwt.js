import { redirectToLogin } from './redirect.js';

export const getJwtToken = () => {
    try {
        return sessionStorage.getItem("jwt");
    } catch (error) {
        console.error('Error getting JWT token:', error);
        return null;
    }
}

export const setJwtToken = (token) => {
    try {
        sessionStorage.setItem("jwt", token);
    } catch (error) {
        console.error('Error setting JWT token:', error);
    }
}

const isJwtExpired = () => {
    const payload = parseJwt(getJwtToken());
    if (!payload.exp) {
        return false;
    }

    const currentTime = Math.floor(Date.now() / 1000);

    return payload.exp < currentTime;

}

const parseJwt = (token) => {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace('-', '+').replace('_', '/');
    return JSON.parse(atob(base64));
}

export const checkSession = () => {
    const token = getJwtToken();
    if (token) {
        if (isJwtExpired(token)) {
            //logout();
            return false;
        } else {
            return true;
        }
    } else {
        return false;
    }
};

export const logout = () => {
    localStorage.removeItem('jwt');
    redirectToLogin()
}