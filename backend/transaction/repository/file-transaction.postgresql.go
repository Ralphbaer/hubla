package repository

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hlog"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// FileTransactionPostgresRepository implements the FileTransactionRepository interface for storing and retrieving file transactions in Postgres.
type FileTransactionPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewFileTransactionPostgreSQLRepository creates a new instance of FileTransactionPostgresRepository with the given Postgres connection.
func NewFileTransactionPostgreSQLRepository(c *hpostgres.PostgresConnection) *FileTransactionPostgresRepository {
	return &FileTransactionPostgresRepository{
		connection: c,
	}
}

// Save saves the provided file transaction to the PostgreSQL database.
// Returns an error if there's any issue.
func (r *FileTransactionPostgresRepository) Save(ctx context.Context, ft *e.FileTransaction) error {
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

	query := `INSERT INTO file_transaction(id, file_id, transaction_Id) VALUES ($1, $2, $3)`
	if _, err = tx.ExecContext(ctx, query, ft.ID, ft.FileID, ft.TransactionID); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Find retrieves a file transaction by its ID from the PostgreSQL database.
// Returns a pointer to the transaction and an error if there's any issue.
func (r *FileTransactionPostgresRepository) Find(ctx context.Context, ID string) (*e.FileTransaction, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, file_id, transaction_id
        FROM file_transaction
        WHERE id = $1`
	var fileTransaction e.FileTransaction
	err = db.QueryRowContext(ctx, query, ID).Scan(&fileTransaction.ID, &fileTransaction.FileID, &fileTransaction.TransactionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.Seller{}).Name(),
				Err:        err,
			}
		}
		return nil, err
	}

	return &fileTransaction, nil
}
