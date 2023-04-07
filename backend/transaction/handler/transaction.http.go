package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"
	uc "github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/gorilla/mux"
)

// TransactionHandler represents a handler which deal with Transaction resource operations
type TransactionHandler struct {
	UseCase *uc.TransactionUseCase
}

// Create is a method that handles incoming requests for creating file transactions,
// processing the request data, storing the metadata and content, and returning an appropriate response.
// swagger:operation POST /transactions/upload nil Create
// Register a new Transaction into database
// ---
// parameters:
//   - name: input
//     in: body
//     type: string
//     description: The payload
//     required: true
//     schema:
//     "$ref": "#/definitions/CreateFileMetadata"
//
// security:
//   - Definitions: []
//
// responses:
//
//	'204':
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
		logger := hlog.NewLoggerFromContext(ctx)
		logger.Debug("Create transactions")

		binaryData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Error(err.Error())
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
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		transactions, err := handler.UseCase.StoreFileContent(ctx, binaryData)
		if err != nil {
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		if err := handler.UseCase.CreateFileTransactions(ctx, fileID, transactions); err != nil {
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s/transaction-files/%s/transactions", r.Host, fileID))
		w.Header().Set("Content-Type", "application/json")

		commonHTTP.Created(w, nil)
	})
}

func (handler *TransactionHandler) GetFileTransactions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := hlog.NewLoggerFromContext(ctx)
		logger.Debug("Get file transactions")

		fileID := mux.Vars(r)["id"]
		transactions, err := handler.UseCase.GetFileTransactions(r.Context(), fileID)
		if err != nil {
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}
		if transactions == nil {
			logger.Error(err.Error())
			commonHTTP.OK(w, []interface{}{})
			return
		}

		commonHTTP.OK(w, transactions)
	})
}
