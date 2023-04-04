package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type FileTransactionPostgresRepository struct {
	connection *common.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewFileTransactionPostgreSQLRepository(c *common.PostgresConnection) *FileTransactionPostgresRepository {
	return &FileTransactionPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *FileTransactionPostgresRepository) Save(ctx context.Context, ft *e.FileTransaction) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO file_transactions(id, file_id, transaction_Id) VALUES ($1, $2, $3) RETURNING id`
	var fileTransactionID string
	if err := tx.QueryRowContext(ctx, query, ft.ID, ft.FileID, ft.TransactionID).Scan(&fileTransactionID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *FileTransactionPostgresRepository) Find(ctx context.Context, ID string) (*e.FileTransaction, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	var fileTransaction e.FileTransaction
	if err := db.QueryRow(`
		 SELECT *
		 FROM file_transactions
		 WHERE ID = $1`, ID).Scan(&fileTransaction.ID, &fileTransaction.FileID, &fileTransaction.TransactionID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &fileTransaction, nil
}
