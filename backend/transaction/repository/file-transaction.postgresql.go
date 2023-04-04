package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type FileTransactionPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewFileTransactionPostgreSQLRepository(c *hpostgres.PostgresConnection) *FileTransactionPostgresRepository {
	return &FileTransactionPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *FileTransactionPostgresRepository) Save(ctx context.Context, ft *e.FileTransaction) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return hpostgres.WithError(err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return hpostgres.WithError(err)
	}
	defer tx.Rollback()

	query := `INSERT INTO file_transactions(id, file_id, transaction_Id) VALUES ($1, $2, $3) RETURNING id`
	var fileTransactionID string
	err = tx.QueryRowContext(ctx, query, ft.ID, ft.FileID, ft.TransactionID).Scan(&fileTransactionID)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return hpostgres.WithError(err)
		}
		if errors.Is(err, sql.ErrTxDone) {
			return hpostgres.WithError(err)
		}
		return hpostgres.WithError(err)
	}

	err = tx.Commit()
	if err != nil {
		return hpostgres.WithError(err)
	}

	return nil
}

func (r *FileTransactionPostgresRepository) Find(ctx context.Context, ID string) (*e.FileTransaction, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	query := `
        SELECT id, file_id, transaction_id
        FROM file_transactions
        WHERE id = $1`
	var fileTransaction e.FileTransaction
	err = db.QueryRowContext(ctx, query, ID).Scan(&fileTransaction.ID, &fileTransaction.FileID, &fileTransaction.TransactionID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, hpostgres.WithError(err)
		}
		if errors.Is(err, sql.ErrConnDone) {
			return nil, hpostgres.WithError(err)
		}
		if errors.Is(err, sql.ErrTxDone) {
			return nil, hpostgres.WithError(err)
		}
		return nil, hpostgres.WithError(err)
	}

	return &fileTransaction, nil
}
