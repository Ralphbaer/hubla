package handler

import (
	"io/ioutil"
	"net/http"

	commonHTTP "github.com/Ralphbaer/hubla/common/net/http"
	uc "github.com/Ralphbaer/hubla/sales/usecase"
)

// SalesHandler represents a handler which deal with Sales resource operations
type SalesHandler struct {
	UseCase *uc.SalesUseCase
}

// Create creates a new Sales in the repository
// swagger:operation POST /sales sales Create
// Register a new Sales into database
// ---
// parameters:
//   - name: input
//     in: body
//     type: string
//     description: The payload
//     required: true
//     schema:
//     "$ref": "#/definitions/CreateSalesInput"
//
// security:
//   - Definitions: []
//
// responses:
//
//	'201':
//	  description: Success Operation
//	  schema:
//	    "$ref": "#/definitions/Sales"
//	'400':
//	  description: Invalid Input - Input has invalid/missing values
//	  schema:
//	    "$ref": "#/definitions/ValidationError"
//	  examples:
//	    "application/json":
//	      code: 400
//	      message: message
//	'409':
//	  description: Conflict - sales document already taken
//	  schema:
//	    "$ref": "#/definitions/ResponseError"
//	  examples:
//	    "application/json":
//	      code: 409
//	      message: message
func (handler *SalesHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		binaryData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Process the binary data
		result, err := handler.UseCase.StoreFileContent(r.Context(), binaryData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		commonHTTP.Accepted(w, result)
	})
}
