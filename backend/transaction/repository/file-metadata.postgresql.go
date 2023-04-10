package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hlog"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// FileMetadataPostgresRepository implements the Repository interface for storing and retrieving file metadata in Postgres.
type FileMetadataPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewFileMetadataPostgreSQLRepository creates a new instance of FileMetadataPostgresRepository with the given Postgres connection.
func NewFileMetadataPostgreSQLRepository(c *hpostgres.PostgresConnection) *FileMetadataPostgresRepository {
	return &FileMetadataPostgresRepository{
		connection: c,
	}
}

// Save inserts the given file metadata into the Postgres database, wrapped in a transaction.
// If the entity already exists, it returns a common.EntityConflictError.
func (r *FileMetadataPostgresRepository) Save(ctx context.Context, fm *e.FileMetadata) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			hlog.NewLoggerFromContext(ctx).Errorf("Failed to rollback transaction: %v", err)
		}
	}()

	query := `INSERT INTO file_metadata(id, file_size, disposition, hash, binary_data, created_at) VALUES ($1, $2, $3, $4, $5, DEFAULT)`
	if _, err = tx.ExecContext(ctx, query, fm.ID, fm.FileSize, fm.Disposition, fm.Hash, fm.BinaryData); err != nil {
		if pqerr := err.(*pq.Error); pqerr.Code == "23505" {
			return common.EntityConflictError{
				Message: err.Error(),
				Err:     err,
			}
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
