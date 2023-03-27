package handler

import (
	"context"
	"net/http"
	"strings"

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
		contentType := r.Header.Get("Content-Type")

		if !strings.HasPrefix(contentType, "multipart/form-data") {
			commonHTTP.BadRequest(w, "Unsupported content type")
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			commonHTTP.BadRequest(w, err.Error())
			return
		}
		defer file.Close()

		sfp := &uc.SalesFileUpload{File: file}

		result, err := handler.UseCase.StoreFileContent(context.Background(), sfp)
		if err != nil {
			commonHTTP.InternalServerError(w, err.Error())
			return
		}

		commonHTTP.OK(w, result)
	})
}
