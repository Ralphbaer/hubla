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
	lib "github.com/Ralphbaer/hubla/common/net/http"
	"github.com/Ralphbaer/hubla/transaction/handler"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(sh *handler.SellerHandler, th *handler.TransactionHandler) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	r.Use(lib.WithCorrelationID)

	r.Handle("/transactions/upload", th.Create()).Methods("POST")
	r.Handle("/sellers/{id}/balance", sh.GetSellerBalanceByID()).Methods("GET")

	// Common

	r.HandleFunc("/transaction/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "transaction", "transaction API Documentation", r)

	return r
}
