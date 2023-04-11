package http

import (
	"net/http"

	"github.com/Ralphbaer/hubla/backend/common"
)

// ResponseError represents a HTTP response error payload
// swagger:model ResponseError
type ResponseError struct {
	StatusCode int     `json:"status_code,omitempty"`
	ErrCode    string  `json:"err_code,omitempty"`
	Message    string  `json:"message,omitempty"`
	Origin     *string `json:"origin,omitempty"`
}

func (r ResponseError) Error() string {
	return r.Message
}

// ValidationError represents an error occurred when a request to an action is invalid
// swagger:model ValidationError
type ValidationError struct {
	StatusCode int              `json:"status_code,omitempty"`
	ErrCode    string           `json:"err_code,omitempty"`
	Message    string           `json:"message,omitempty"`
	Fields     FieldValidations `json:"fields,omitempty"`
}

func (r ValidationError) Error() string {
	return r.Message
}

// FieldValidations represents a field error response.
type FieldValidations map[string]string

// WithError handle errors in handlers and returns the appropriated response
func WithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case common.EntityNotFoundError:
		NotFound(w, e)
	case common.EntityConflictError:
		Conflict(w, e)
	case common.ValidationError:
		BadRequest(w, ValidationError{
			StatusCode: 400,
			Message:    e.Error(),
			Fields:     nil,
		})
	case common.UnprocessableOperationError:
		UnprocessableEntity(w, e)
	case common.UnauthorizedError:
		Unauthorized(w, e)
	case common.ForbiddenError:
		Forbidden(w, e.Error())
	case *ValidationError, ValidationError:
		BadRequest(w, e)
	case ResponseError:
		rErr, _ := err.(ResponseError)
		JSONResponseError(w, rErr)
	default:
		InternalServerError(w, e.Error())
	}
}
