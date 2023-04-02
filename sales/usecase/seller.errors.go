package usecase

import "errors"

// ErrCreatorNotFound is throwed when a document (sales unique field) already exists in the repository
var ErrCreatorNotFound = errors.New("creator not found")
