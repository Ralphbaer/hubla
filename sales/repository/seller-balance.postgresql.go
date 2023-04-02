package repository

import (
	"context"
	"fmt"

	"github.com/Ralphbaer/hubla/common"
	e "github.com/Ralphbaer/hubla/sales/entity"
	"github.com/shopspring/decimal"
)

// SellerBalancePostgresRepository represents a MongoDB implementation of PartnerRepository interface
type SellerBalancePostgresRepository struct {
	connection *common.PostgresConnection
}

// NewSellerBalancePostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewSellerBalancePostgreSQLRepository(c *common.PostgresConnection) *SellerBalancePostgresRepository {
	return &SellerBalancePostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *SellerBalancePostgresRepository) Upsert(ctx context.Context, p *e.SellerBalance) (*decimal.Decimal, error) {
	db, err := r.connection.Connect()
	if err != nil {
		return nil, err
	}

	// Start a transaction with default options
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Prepare the INSERT statement with the ON CONFLICT DO UPDATE clause
	query := `INSERT INTO seller_balances (id, seller_id, balance, created_at)
              VALUES ($1, $2, $3, DEFAULT) 
              ON CONFLICT (seller_id) DO UPDATE SET balance = seller_balances.balance + $3
              RETURNING balance`
	//stmt, err := tx.Prepare(query)
	//if err != nil {
	//	return nil, err
	//}
	//defer stmt.Close()

	// Execute the prepared statement with the given sellerID and amount
	var newBalance float64
	if err := tx.QueryRowContext(ctx, query, p.ID, p.Balance).Scan(&newBalance); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToInsertSeller, err)
	}

	// Execute the prepared statement with the given sellerID and amount
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	fmt.Printf("Balance for seller %s updated by %d to %f\n", p.SellerID, p.Balance, newBalance)

	return nil, nil
}
