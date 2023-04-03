package usecase

import "errors"

// ErrCreatorNotFound is throwed when a document (transaction unique field) already exists in the repository
var ErrCreatorNotFound = errors.New("creator not found")
