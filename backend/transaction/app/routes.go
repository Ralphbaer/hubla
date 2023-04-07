// Package app Sales API.
//
// This guide describes all Hubla Sales API and usecase.
//
//	Schemes: http, https
//	BasePath: /transaction
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Extensions:
//	x-meta-value: value
//	x-meta-array:
//	  - value1
//	  - value2
//	x-meta-array-obj:
//	  - name: obj
//	    value: field
//
// swagger:meta
package app

import (
	"github.com/Ralphbaer/hubla/backend/common/hlog"
	lib "github.com/Ralphbaer/hubla/backend/common/net/http"
	"github.com/Ralphbaer/hubla/backend/transaction/handler"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(sh *handler.SellerHandler, th *handler.TransactionHandler, logger hlog.Logger) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	lib.AllowFullOptionsWithCORS(r)
	r.Use(lib.WithCorrelationID)
	r.Use(lib.WithLog(lib.WithLogger(logger)))

	// Transaction Files

	r.Handle("/file-transactions", th.Create()).Methods("POST")
	r.Handle("/file-transactions/{id}/transactions", th.GetFileTransactions()).Methods("GET")

	// Sellers

	r.Handle("/sellers/{id}/balance", sh.GetSellerBalanceByID()).Methods("GET")

	// Common

	r.HandleFunc("/transaction/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "transaction", "transaction API Documentation", r)

	return r
}
