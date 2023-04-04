package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type TransactionPostgresRepository struct {
	connection *common.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewTransactionPostgreSQLRepository(c *common.PostgresConnection) *TransactionPostgresRepository {
	return &TransactionPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *TransactionPostgresRepository) Save(ctx context.Context, t *e.Transaction) (*e.Transaction, error) {
	db, dbErr := r.connection.GetDB()
	if dbErr != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, dbErr)
	}

	tx, txErr := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if txErr != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToBeginTransaction, txErr)
	}
	defer tx.Rollback()

	const insertQuery = `INSERT INTO transactions(id, t_type, t_date, product_id, amount, seller_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, DEFAULT) RETURNING id`

	var insertedID string
	if err := tx.QueryRowContext(ctx, insertQuery, t.ID, e.TransactionTypeMapString[t.TType], t.TDate, t.ProductID, t.Amount, t.SellerID).Scan(&insertedID); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToInsertTransaction, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	t.ID = insertedID

	return t, nil
}

func (r *TransactionPostgresRepository) List(ctx context.Context, fileID string) ([]*e.Transaction, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	rows, err := db.QueryContext(ctx, `
		SELECT t.*
		FROM transactions t
		JOIN file_transactions ft ON t.id = ft.transaction_id
		WHERE ft.file_id = $1
	`, fileID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToQueryDatabase, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Failed to close rows: %v", err)
		}
	}()

	var transactions []*e.Transaction
	var tTypeStr string // Add a variable to store the raw TType string
	for rows.Next() {
		transaction := &e.Transaction{} // Create a new transaction variable for each iteration
		err := rows.Scan(&transaction.ID, &tTypeStr, &transaction.TDate, &transaction.ProductID, &transaction.Amount, &transaction.SellerID, &transaction.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrFailedToScanRow, err)
		}
		if tType, ok := e.TransactionTypeMapEnum[tTypeStr]; !ok {
			return nil, fmt.Errorf("%w: invalid transaction type: %s", ErrInvalidDatabaseData, tTypeStr)
		} else {
			transaction.TType = tType // Convert the string to TransactionTypeEnum
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToIterateRows, err)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("%w: no transactions found for file ID %s", ErrNotFound, fileID)
	}

	return transactions, nil
}
