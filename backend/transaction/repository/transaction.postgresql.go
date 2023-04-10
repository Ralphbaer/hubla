package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// TransactionPostgresRepository implements the TransactionRepository interface for storing and retrieving transactions in Postgres.
type TransactionPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewTransactionPostgreSQLRepository creates a new instance of TransactionPostgresRepository with the given Postgres connection.
func NewTransactionPostgreSQLRepository(c *hpostgres.PostgresConnection) *TransactionPostgresRepository {
	return &TransactionPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *TransactionPostgresRepository) Save(ctx context.Context, t *e.Transaction) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO transaction_record(id, t_type, t_date, product_id, amount, seller_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, DEFAULT)`
	if _, err := tx.ExecContext(ctx, query, t.ID, e.TransactionTypeMapString[t.TType], t.TDate, t.ProductID, t.Amount, t.SellerID); err != nil {
		if pqerr := err.(*pq.Error); pqerr.Code == "23505" {
			return common.EntityConflictError{
				Message: err.Error(),
				Err:     err,
			}
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// ListTransactionsByFileID retrieves a slice of Transaction entities from the Postgres database by the given file ID.
// It returns a slice of Transaction objects if found, or an error if the query fails.
func (r *TransactionPostgresRepository) ListTransactionsByFileID(ctx context.Context, fileID string) ([]*e.Transaction, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, `
		SELECT t.*
		FROM transaction_record t
		JOIN file_transaction ft ON t.id = ft.transaction_id
		WHERE ft.file_id = $1
	`, fileID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Failed to close rows: %v", err)
		}
	}()

	var transactions []*e.Transaction
	var tTypeStr string
	for rows.Next() {
		transaction := &e.Transaction{}
		if err := rows.Scan(&transaction.ID, &tTypeStr, &transaction.TDate, &transaction.ProductID,
			&transaction.Amount, &transaction.SellerID, &transaction.CreatedAt); err != nil {
			return nil, err
		}
		transaction.TType = e.TransactionTypeMapEnum[tTypeStr]
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
