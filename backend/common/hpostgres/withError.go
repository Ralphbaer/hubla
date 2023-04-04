package hpostgres

import (
	"database/sql"
	"errors"
	"fmt"
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
	case errors.Is(err, sql.ErrConnDone):
		return fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	case errors.Is(err, sql.ErrTxDone):
		return fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	case errors.Is(err, ErrPostgreSQLDuplicatedDocument):
		return fmt.Errorf("%w: %v", ErrPostgreSQLDuplicatedDocument, err)
	case errors.Is(err, ErrInvalidDatabaseData):
		return fmt.Errorf("%w: %v", ErrInvalidDatabaseData, err)
	case errors.Is(err, ErrFailedToIterateRows):
		return fmt.Errorf("%w: %v", ErrFailedToIterateRows, err)
	case errors.Is(err, ErrNotFound):
		return fmt.Errorf("%w: %v", ErrNotFound, err)
	case errors.Is(err, ErrFailedToQueryDatabase):
		return fmt.Errorf("%w: %v", ErrFailedToQueryDatabase, err)
	case errors.Is(err, ErrFailedToScanRow):
		return fmt.Errorf("%w: %v", ErrFailedToScanRow, err)
	case errors.Is(err, ErrFailedToInsertTransaction):
		return fmt.Errorf("%w: %v", ErrFailedToInsertTransaction, err)
	default:
		return err
	}
}
