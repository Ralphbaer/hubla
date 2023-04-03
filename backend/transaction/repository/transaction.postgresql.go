package repository

import (
	"context"
	"database/sql"
	"fmt"

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
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToBeginTransaction, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO transactions(id, t_type, t_date, product_id, amount, seller_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, DEFAULT) RETURNING id`
	var sellerID string
	if err := tx.QueryRowContext(ctx, query, t.ID, e.TransactionTypeMapString[t.TType], t.TDate, t.ProductID, t.Amount, t.SellerID).Scan(&sellerID); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToInsertSeller, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	return nil, nil
}
