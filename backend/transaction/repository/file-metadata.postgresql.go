package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type FileMetadataPostgresRepository struct {
	connection *common.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewFileMetadataPostgreSQLRepository(c *common.PostgresConnection) *FileMetadataPostgresRepository {
	return &FileMetadataPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *FileMetadataPostgresRepository) Save(ctx context.Context, fm *e.FileMetadata) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrFailedToBeginTransaction, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO file_metadata(id, file_size, disposition, hash, binary_data, created_at) VALUES ($1, $2, $3, $4, $5, DEFAULT) RETURNING id`
	var fileMetadataID string
	err = tx.QueryRowContext(ctx, query, fm.ID, fm.FileSize, fm.Disposition, fm.Hash, fm.BinaryData).Scan(&fileMetadataID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" && pqErr.Constraint == "file_metadata_hash_key" {
			return fmt.Errorf("file metadata with hash '%s' already exists", fm.Hash)
		}
		return fmt.Errorf("%w: %v", ErrFailedToInsertTransaction, err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	return nil
}

func (r *FileMetadataPostgresRepository) Find(ctx context.Context, hash string) (*e.FileMetadata, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	query := `
        SELECT id, file_size, disposition, hash, binary_data
        FROM file_metadata
        WHERE hash = $1`
	var fileMetadata e.FileMetadata
	err = db.QueryRowContext(ctx, query, hash).Scan(&fileMetadata.ID, &fileMetadata.FileSize, &fileMetadata.Disposition,
		&fileMetadata.Hash, &fileMetadata.BinaryData)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %v", ErrNotFound, err)
		}
		if errors.Is(err, sql.ErrConnDone) {
			return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
		}
		if errors.Is(err, sql.ErrTxDone) {
			return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
		}
		return nil, fmt.Errorf("%w: %v", ErrFailedToScanRow, err)
	}

	return &fileMetadata, nil
}
