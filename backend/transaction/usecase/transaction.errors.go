package usecase

import "errors"

// ErrSalesDocumentConflict is throwed when a document (transaction unique field) already exists in the repository
var (
	ErrInvalidLineFormat         = errors.New("invalid line format")
	ErrParsingTransactionType    = errors.New("error parsing transaction type")
	ErrInvalidTransactionType    = errors.New("invalid transaction type")
	ErrParsingDate               = errors.New("error parsing date")
	ErrParsingAmount             = errors.New("error parsing amount")
	ErrInvalidSellerName         = errors.New("invalid seller name")
	ErrScanningFile              = errors.New("error scanning file %v")
	ErrFindingOrCreateSeller     = errors.New("error finding or creating seller")
	ErrFindingOrCreateProduct    = errors.New("error finding or creating product")
	ErrSavingTransaction         = errors.New("error saving transaction")
	ErrUpdatingSellerBalance     = errors.New("error updating seller balance")
	ErrParsingParsingLine        = errors.New("error parsing line %d: %v")
	ErrFileMetadataAlreadyExists = errors.New("file metadata with id %s already exists")
	ErrInvalidAmount             = errors.New("invalid amount")
	ErrInvalidDate               = errors.New("invalid date")
)
