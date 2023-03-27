package usecase

import "errors"

// ErrSalesDocumentConflict is throwed when a document (sales unique field) already exists in the repository
var ErrSalesDocumentConflict = errors.New("sales document already taken")
