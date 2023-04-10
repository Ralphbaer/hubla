package handler

import (
	"encoding/json"
	"net/http"

	uc "github.com/Ralphbaer/hubla/backend/auth/usecase"
	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"

	"github.com/Ralphbaer/hubla/backend/common/jwt"
)

const tokenMaxAge = 60

// LoginHandler represents a handler which deal with Transaction resource operations
type LoginHandler struct {
	UseCase *uc.UserUseCase
	JWTAuth *jwt.Auth
}

// SignInUser handles HTTP requests for user sign-in by verifying credentials,
// generating an access token, and sending a JSON response with the token.
func (handler *LoginHandler) SignInUser(p interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credentials := p.(*uc.SignInInput)
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		user, err := handler.UseCase.GetUserByEmail(r.Context(), credentials.Email)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		if err := jwt.ComparePassword(user.Password, credentials.Password); err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		accessToken, err := handler.JWTAuth.CreateAccessToken(user.ID)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		// Set Cookies
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			MaxAge:   tokenMaxAge * 60,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: true,
		})
		w.Header().Set("Content-Type", "application/json")
		commonHTTP.OK(w, json.NewEncoder(w).Encode(map[string]interface{}{
			"status":       "success",
			"access_token": accessToken,
		}))
	})
}
