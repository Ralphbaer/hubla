package repository

import "errors"

var (
	ErrFailedToConnectToDatabase = errors.New("failed to connect to database")
	ErrFailedToBeginTransaction  = errors.New("failed to begin transaction")
	ErrFailedToCommitTransaction = errors.New("failed to commit transaction")
	ErrFailedToInsertTransaction = errors.New("failed to insert transaction")

	ErrFailedToInsertSeller  = errors.New("failed to insert seller")
	ErrFailedToQueryDatabase = errors.New("failed to query database")
	ErrFailedToScanRow       = errors.New("failed to scan row")
	ErrInvalidDatabaseData   = errors.New("invalid database data")
	ErrFailedToIterateRows   = errors.New("failed to iterate rows")
	ErrNotFound              = errors.New("not found")
)
