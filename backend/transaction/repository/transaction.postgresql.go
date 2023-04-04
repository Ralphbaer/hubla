package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type TransactionPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewTransactionPostgreSQLRepository(c *hpostgres.PostgresConnection) *TransactionPostgresRepository {
	return &TransactionPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *TransactionPostgresRepository) Save(ctx context.Context, t *e.Transaction) (*string, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, hpostgres.WithError(err)
	}
	defer tx.Rollback()

	const insertQuery = `INSERT INTO transactions(id, t_type, t_date, product_id, amount, seller_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, DEFAULT) RETURNING id`
	var insertedID string
	if err := tx.QueryRowContext(ctx, insertQuery, t.ID, e.TransactionTypeMapString[t.TType], t.TDate, t.ProductID, t.Amount, t.SellerID).Scan(&insertedID); err != nil {
		return nil, hpostgres.WithError(err)
	}

	return &insertedID, nil
}

func (r *TransactionPostgresRepository) List(ctx context.Context, fileID string) ([]*e.Transaction, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	rows, err := db.QueryContext(ctx, `
		SELECT t.*
		FROM transactions t
		JOIN file_transactions ft ON t.id = ft.transaction_id
		WHERE ft.file_id = $1
	`, fileID)
	if err != nil {
		return nil, hpostgres.WithError(err)
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
		if err := rows.Scan(&transaction.ID, &tTypeStr, &transaction.TDate, &transaction.ProductID,
			&transaction.Amount, &transaction.SellerID, &transaction.CreatedAt); err != nil {
			return nil, hpostgres.WithError(err)
		}
		transaction.TType = e.TransactionTypeMapEnum[tTypeStr] // Convert the string to TransactionTypeEnum
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, hpostgres.WithError(err)
	}

	return transactions, nil
}
