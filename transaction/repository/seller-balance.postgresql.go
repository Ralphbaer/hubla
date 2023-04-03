package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ralphbaer/hubla/common"
	e "github.com/Ralphbaer/hubla/transaction/entity"
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
func (r *SellerBalancePostgresRepository) Upsert(ctx context.Context, p *e.SellerBalance) (*float64, error) {
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
	query := `INSERT INTO seller_balances (id, seller_id, balance, updated_at, created_at)
              VALUES ($1, $2, $3, $4, DEFAULT) 
              ON CONFLICT (seller_id) DO UPDATE SET balance = seller_balances.balance + $3
              RETURNING balance`
	//stmt, err := tx.Prepare(query)
	//if err != nil {
	//	return nil, err
	//}
	//defer stmt.Close()

	// Execute the prepared statement with the given sellerID and amount
	var newBalance float64
	if err := tx.QueryRowContext(ctx, query, p.ID, p.SellerID, p.Balance, p.UpdatedAt).Scan(&newBalance); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToInsertSeller, err)
	}

	// Execute the prepared statement with the given sellerID and amount
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	fmt.Printf("Balance for seller %s updated by %s to %f\n", p.SellerID, p.Balance.String(), newBalance)

	return &newBalance, nil
}

func (r *SellerBalancePostgresRepository) Find(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	db, err := r.connection.Connect()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	sellerBalanceView := &e.SellerBalanceView{}
	if err := db.QueryRow(`
		SELECT s.id, s.name, sb.balance, sb.updated_at
		FROM seller_balances sb
		JOIN sellers s ON s.id = sb.seller_id
		WHERE s.id = $1;`, sellerID).Scan(&sellerBalanceView.SellerID, &sellerBalanceView.SellerName,
		&sellerBalanceView.SellerBalance, &sellerBalanceView.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Seller not found
		}
		return nil, err // Other error
	}

	return sellerBalanceView, nil
}
