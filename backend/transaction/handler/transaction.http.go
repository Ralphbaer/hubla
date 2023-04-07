package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"
	uc "github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/gorilla/mux"
)

// TransactionHandler represents a handler which deal with Transaction resource operations
type TransactionHandler struct {
	UseCase *uc.TransactionUseCase
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
func (handler *TransactionHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		binaryData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}
		ctfm := &uc.CreateFileMetadata{
			FileSize:    r.Header.Get("Content-length"),
			Disposition: r.Header.Get("Content-Disposition"),
			BinaryData:  binaryData,
		}

		fileID, err := handler.UseCase.StoreFileMetadata(ctx, ctfm)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		transactions, err := handler.UseCase.StoreFileContent(ctx, binaryData)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		if err := handler.UseCase.CreateFileTransactions(ctx, fileID, transactions); err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s/files/%s/transactions", r.Host, fileID))
		w.Header().Set("Content-Type", "application/json")

		commonHTTP.Created(w, nil)
	})
}

func (handler *TransactionHandler) GetFileTransactions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileID := mux.Vars(r)["id"]
		transactions, err := handler.UseCase.GetFileTransactions(r.Context(), fileID)
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}
		if transactions == nil {
			commonHTTP.OK(w, []interface{}{})
			return
		}

		commonHTTP.OK(w, transactions)
	})
}
