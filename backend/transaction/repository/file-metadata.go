package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/file_metadata_repository_mock.go -package=mocks . FileMetadataRepository

// FileMetadataRepository manages transaction repository operations
type FileMetadataRepository interface {
	Save(ctx context.Context, fm *e.FileMetadata) error
}
