package repository

import "errors"

var (
	ErrFailedToConnectToDatabase = errors.New("failed to connect to database")
	ErrFailedToBeginTransaction  = errors.New("failed to begin transaction")
	ErrFailedToInsertSeller      = errors.New("failed to insert seller")
	ErrFailedToCommitTransaction = errors.New("failed to commit transaction")
)
