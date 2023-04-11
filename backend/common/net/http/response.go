package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hlog"
)

// Unauthorized respond with HTTP 401 Unauthorized and payload.
func Unauthorized(w http.ResponseWriter, err common.UnauthorizedError) {
	JSONResponse(w, http.StatusUnauthorized, &ResponseError{
		ErrCode:    err.ErrCode,
		StatusCode: http.StatusUnauthorized,
		Message:    err.Message,
	})
}

// Forbidden respond with HTTP 403 Forbidden.
func Forbidden(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusForbidden, &ResponseError{
		StatusCode: http.StatusForbidden,
		Message:    message,
	})
}

// BadRequest respond with HTTP 400 BadRequest and payload.
func BadRequest(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusBadRequest, s)
}

// Created respond with HTTP 201 StatusOK and payload.
func Created(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusCreated, s)
}

// OK respond with HTTP 200 StatusOK and payload.
func OK(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusOK, s)
}

// NoContent respond with success HTTP 204 StatusNoContent.
func NoContent(w http.ResponseWriter) {
	JSONResponse(w, http.StatusNoContent, nil)
}

// Accepted respond with HTTP 202 StatusAccepted and payload.
func Accepted(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusAccepted, s)
}

// PartialContent respond with HTTP 206 PartialContent and payload.
func PartialContent(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusPartialContent, s)
}

// RangeNotSatisfiable respond with HTTP 416 RangeNotSatisfiable.
func RangeNotSatisfiable(w http.ResponseWriter) {
	JSONResponse(w, http.StatusRequestedRangeNotSatisfiable, nil)
}

// NotFound respond with HTTP 404 NotFound and payload.
func NotFound(w http.ResponseWriter, err common.EntityNotFoundError) {
	JSONResponse(w, http.StatusNotFound, &ResponseError{
		StatusCode: http.StatusNotFound,
		ErrCode:    err.ErrCode,
		Message:    err.Error(),
	})
}

// Conflict respond with HTTP 409 Conflict.
func Conflict(w http.ResponseWriter, err common.EntityConflictError) {
	JSONResponse(w, http.StatusConflict, &ResponseError{
		StatusCode: http.StatusConflict,
		ErrCode:    err.ErrCode,
		Message:    err.Error(),
	})
}

// UnprocessableEntity respond with HTTP 422 UnprocessableEntity.
func UnprocessableEntity(w http.ResponseWriter, err common.UnprocessableOperationError) {
	JSONResponse(w, http.StatusUnprocessableEntity, &ResponseError{
		StatusCode: http.StatusUnprocessableEntity,
		ErrCode:    err.ErrCode,
		Message:    err.Error(),
	})
}

// InternalServerError respond with HTTP 500 InternalServerError and message.
func InternalServerError(w http.ResponseWriter, err string) {
	JSONResponse(w, http.StatusInternalServerError, &ResponseError{
		Message: "Internal Server Error",
	})
}

// JSONResponseError respond with a ResponseError
func JSONResponseError(w http.ResponseWriter, err ResponseError) {
	JSONResponse(w, err.StatusCode, err)
}

// JSONResponse respond with given HTTP status and given payload.
func JSONResponse(w http.ResponseWriter, status int, s interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w)
	if s != nil {
		payload, _ := json.Marshal(s)
		if _, err := w.Write(payload); err != nil {
			hlog.NewLoggerFromContext(context.TODO()).Errorf("Failed to write JSON response: %v", err)
		}
	}
}
