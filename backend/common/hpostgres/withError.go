package hpostgres

import (
	"database/sql"
	"errors"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/lib/pq"
)

// ResponseError represents a HTTP response error payload
// swagger:model ResponseError
type ResponseError struct {
	Code    int     `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Origin  *string `json:"origin,omitempty"`
}

func (r ResponseError) Error() string {
	return r.Message
}

// ValidationError represents an error occurred when a request to an action is invalid
// swagger:model ValidationError
type ValidationError struct {
	Code    int              `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Fields  FieldValidations `json:"fields,omitempty"`
}

func (r ValidationError) Error() string {
	return r.Message
}

// FieldValidations represents a field error response.
type FieldValidations map[string]string

// WithError handle errors in handlers and returns the appropriated response
func WithError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return common.EntityNotFoundError{
			Message: err.Error(),
			Err:     err,
		}
	case errors.As(err, new(*pq.Error)):
		pqerr := err.(*pq.Error)
		if pqerr.Code == "23505" {
			return common.EntityConflictError{
				Message: err.Error(),
				Err:     err,
			}
		}
	}

	return err
}
