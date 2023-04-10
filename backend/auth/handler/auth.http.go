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

		access_token, err := handler.JWTAuth.CreateAccessToken(user.ID)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		// Set Cookies
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    access_token,
			MaxAge:   tokenMaxAge * 60,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: true,
		})
		w.Header().Set("Content-Type", "application/json")
		commonHTTP.OK(w, json.NewEncoder(w).Encode(map[string]interface{}{
			"status":       "success",
			"access_token": access_token,
		}))
	})
}
