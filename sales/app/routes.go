// Package app Sales API.
//
// This guide describes all Hubla Sales API and usecase.
//
//	Schemes: http, https
//	BasePath: /sales
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
	"github.com/Ralphbaer/hubla/sales/handler"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(p *handler.SalesHandler) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	r.Use(lib.WithCorrelationID)

	r.Handle("/sales/upload", p.Create()).Methods("POST")
	r.Handle("/sales/uploads/", p.Create()).Methods("POST")

	// Common

	r.HandleFunc("/sales/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "sales", "sales API Documentation", r)

	return r
}
