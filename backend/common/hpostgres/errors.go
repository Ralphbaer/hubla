package hpostgres

import (
	"github.com/pkg/errors"
)

// ErrPostgreSQLDuplicatedDocument is throwed when a Document already exists in the repository.
var (
	ErrPostgreSQLDuplicatedDocument = errors.New("Duplicated Document")
	ErrFailedToConnectToDatabase    = errors.New("failed to connect to database")
	ErrFailedToBeginTransaction     = errors.New("failed to begin transaction")
	ErrFailedToCommitTransaction    = errors.New("failed to commit transaction")
	ErrInvalidDatabaseData          = errors.New("invalid database data")
	ErrFailedToIterateRows          = errors.New("failed to iterate rows")
	ErrNotFound                     = errors.New("not found")
	ErrFailedToQueryDatabase        = errors.New("failed to query database")
	ErrFailedToScanRow              = errors.New("failed to scan row")
	ErrFailedToInsertTransaction    = errors.New("failed to insert transaction")
	ErrFailedToUpsertTransaction    = errors.New("failed to upsert transaction")
)
