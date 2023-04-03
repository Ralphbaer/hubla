package common

import (
	"github.com/pkg/errors"
)

// ErrPostgreSQLDuplicatedDocument is throwed when a Document already exists in the repository.
var ErrPostgreSQLDuplicatedDocument = errors.New("Duplicated Document")
