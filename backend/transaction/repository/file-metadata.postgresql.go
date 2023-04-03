package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
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
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO file_metadata(id, file_size, disposition, hash, binary_data, created_at) VALUES ($1, $2, $3, $4, $5, DEFAULT) RETURNING id`
	var FileMetadataID string
	if err := tx.QueryRowContext(ctx, query, fm.ID,
		fm.FileSize, fm.Disposition, fm.Hash, fm.BinaryData, fm.CreatedAt).Scan(&FileMetadataID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *FileMetadataPostgresRepository) Find(ctx context.Context, hash string) (*e.FileMetadata, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	var FileMetadata e.FileMetadata
	if err := db.QueryRow(`
		 SELECT *
		 FROM file_metadata
		 WHERE hash = $1`, hash).Scan(&FileMetadata.ID, &FileMetadata.FileSize, &FileMetadata.Disposition,
		&FileMetadata.Hash, &FileMetadata.BinaryData); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &FileMetadata, nil
}
