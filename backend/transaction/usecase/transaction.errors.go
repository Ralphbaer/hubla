package usecase

import "errors"

// ErrSalesDocumentConflict is throwed when a document (transaction unique field) already exists in the repository
var (
	ErrInvalidLineFormat         = errors.New("invalid line format")
	ErrInvalidTransactionType    = errors.New("invalid transaction type")
	ErrInvalidSellerName         = errors.New("invalid seller name")
	ErrScanningFile              = errors.New("error scanning file %v")
	ErrUpsertingSellerBalance    = errors.New("error upserting seller balance")
	ErrParsingParsingLine        = errors.New("error parsing line %d: %v")
	ErrFileMetadataAlreadyExists = errors.New("file metadata with id %s already exists")
	ErrInvalidAmount             = errors.New("invalid amount")
	ErrInvalidDate               = errors.New("invalid date")
)
