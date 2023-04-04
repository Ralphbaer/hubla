package repository

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type FileMetadataPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewFileMetadataPostgreSQLRepository(c *hpostgres.PostgresConnection) *FileMetadataPostgresRepository {
	return &FileMetadataPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *FileMetadataPostgresRepository) Save(ctx context.Context, fm *e.FileMetadata) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

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

func (r *FileMetadataPostgresRepository) Find(ctx context.Context, hash string) (*e.FileMetadata, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
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
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.Seller{}).Name(),
				Message:    err.Error(),
				Err:        err,
			}
		}
		return nil, err
	}

	return &fileMetadata, nil
}
