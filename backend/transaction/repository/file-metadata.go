package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// FileMetadataRepository manages transaction repository operations
type FileMetadataRepository interface {
	Find(ctx context.Context, hash string) (*e.FileMetadata, error)
	Save(ctx context.Context, fm *e.FileMetadata) error
}
