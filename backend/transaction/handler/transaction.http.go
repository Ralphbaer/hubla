package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"
	uc "github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// TransactionHandler represents a handler which deal with Transaction resource operations
type TransactionHandler struct {
	UseCase *uc.TransactionUseCase
}

// Create returns an http.Handler that creates a new file and its transactions.
// It reads binary data from the request body, stores the file metadata, content, and transactions.
func (handler *TransactionHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := hlog.NewLoggerFromContext(ctx)
		logger.Debug("Create transactions")

		binaryData, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		ctfm := &uc.CreateFileMetadata{
			ID:          uuid.New().String(),
			FileSize:    r.Header.Get("Content-length"),
			Disposition: r.Header.Get("Content-Disposition"),
			BinaryData:  binaryData,
		}

		contentType := http.DetectContentType(binaryData)
		if ok := strings.HasPrefix(contentType, "text/plain"); !ok {
			commonHTTP.WithError(w, commonHTTP.ValidationError{
				Code:    400,
				Message: "Only text files with a .txt format are accepted",
			})
			return
		}

		if ctfm.FileSize == "0" {
			commonHTTP.WithError(w, commonHTTP.ValidationError{
				Code:    400,
				Message: "Please provide a file or ensure the file is not empty.",
			})
			return
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

		if err := handler.UseCase.CreateFileTransactions(ctx, fileID.ID, transactions); err != nil {
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s/file-transactions/%s/transactions", r.Host, fileID.ID))
		w.Header().Set("Content-Type", "application/json")

		commonHTTP.Created(w, fileID)
	})
}

// GetFileTransactions handles to retrieve transactions for a given file ID.
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
			commonHTTP.OK(w, []interface{}{})
			return
		}

		commonHTTP.OK(w, transactions)
	})
}
