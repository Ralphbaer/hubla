package handler

import (
	"encoding/json"
	"net/http"

	uc "github.com/Ralphbaer/hubla/backend/auth/usecase"

	"github.com/Ralphbaer/hubla/backend/common/jwt"
)

// LoginHandler represents a handler which deal with Transaction resource operations
type LoginHandler struct {
	UseCase *uc.UserUseCase
	JWTAuth *jwt.JWTAuth
}

// Create creates a new Transaction in the repository
// swagger:operation POST /Transaction Transaction Create
// Register a new Transaction into database
// ---
// parameters:
//   - name: input
//     in: body
//     type: string
//     description: The payload
//     required: true
//     schema:
//     "$ref": "#/definitions/CreateTransactionInput"
//
// security:
//   - Definitions: []
//
// responses:
//
//	'201':
//	  description: Success Operation
//	  schema:
//	    "$ref": "#/definitions/Transaction"
//	'400':
//	  description: Invalid Input - Input has invalid/missing values
//	  schema:
//	    "$ref": "#/definitions/ValidationError"
//	  examples:
//	    "application/json":
//	      code: 400
//	      message: message
//	'409':
//	  description: Conflict - Transaction document already taken
//	  schema:
//	    "$ref": "#/definitions/ResponseError"
//	  examples:
//	    "application/json":
//	      code: 409
//	      message: message
func (handler *LoginHandler) SignInUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse input
		var credentials uc.LoginInput
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Replace this with your actual method to get the user by email
		user, err := handler.UseCase.GetUserByEmail(r.Context(), credentials.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusBadRequest)
			return
		}

		if err := jwt.ComparePassword(user.Password, credentials.Password); err != nil {
			http.Error(w, "Invalid email or password", http.StatusBadRequest)
			return
		}

		access_token, err := handler.JWTAuth.CreateAccessToken("1asdasds")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		refresh_token, err := handler.JWTAuth.CreateRefreshToken(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Set Cookies
		http.SetCookie(w, &http.Cookie{
			Name:  "access_token",
			Value: access_token,
			//		MaxAge:   config.AccessTokenMaxAge * 60,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: true,
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "refresh_token",
			Value: refresh_token,
			//	MaxAge:   config.RefreshTokenMaxAge * 60,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: true,
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "logged_in",
			Value: "true",
			//	MaxAge:   config.AccessTokenMaxAge * 60,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: false,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":       "success",
			"access_token": access_token,
		})
	})
}
