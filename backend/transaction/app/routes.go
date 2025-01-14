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
	"github.com/Ralphbaer/hubla/backend/common/jwt"
	lib "github.com/Ralphbaer/hubla/backend/common/net/http"
	"github.com/Ralphbaer/hubla/backend/transaction/handler"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(sh *handler.SellerHandler, th *handler.TransactionHandler, logger hlog.Logger) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	api := r.PathPrefix("/api/v1").Subrouter()

	lib.AllowFullOptionsWithCORS(r)
	r.Use(lib.WithCorrelationID)
	r.Use(lib.WithLog(lib.WithLogger(logger)))
	userJWT := jwt.NewJWTAuth(config.AccessTokenPublicKey)

	// Transaction Files

	api.Handle("/transaction/file-transactions", userJWT.Protect(th.Create())).Methods("POST")
	api.Handle("/transaction/file-transactions/{id}/transactions", th.GetFileTransactions()).Methods("GET")
	api.Handle("/transaction/file-transactions/transactions", th.ListFileTransactions()).Methods("GET")

	// Sellers

	api.Handle("/seller/sellers/{id}/balance", sh.GetSellerBalanceByID()).Methods("GET")

	// Common

	api.HandleFunc("/transaction/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "transaction", "transaction API Documentation", r)

	return r
}
