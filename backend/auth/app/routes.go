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
	"github.com/Ralphbaer/hubla/backend/auth/handler"
	"github.com/Ralphbaer/hubla/backend/auth/usecase"
	lib "github.com/Ralphbaer/hubla/backend/common/net/http"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(l *handler.LoginHandler) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	lib.AllowFullOptionsWithCORS(r)
	r.Use(lib.WithCorrelationID)
	api := r.PathPrefix("/api/v1").Subrouter()

	api.Handle("/auth/login", lib.WithBody(new(usecase.SignInInput), l.SignInUser)).Methods("POST")

	// Common

	api.HandleFunc("/auth/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "auth", "auth API Documentation", r)

	return r
}
