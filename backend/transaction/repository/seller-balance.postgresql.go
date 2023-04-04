package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
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

func (r *SellerBalancePostgresRepository) Upsert(ctx context.Context, p *e.SellerBalance) (*float64, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToBeginTransaction, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO seller_balances (id, seller_id, balance, updated_at, created_at)
              VALUES ($1, $2, $3, $4, DEFAULT) 
              ON CONFLICT (seller_id) DO UPDATE SET balance = seller_balances.balance + $3
              RETURNING balance`

	var newBalance float64
	if err := tx.QueryRowContext(ctx, query, p.ID, p.SellerID, p.Balance, p.UpdatedAt).Scan(&newBalance); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToInsertSeller, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	fmt.Printf("Balance for seller %s updated by %s to %f\n", p.SellerID, p.Balance.String(), newBalance)

	return &newBalance, nil
}

func (r *SellerBalancePostgresRepository) Find(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	query := `
        SELECT s.id, s.name, sb.balance, sb.updated_at
        FROM seller_balances sb
        JOIN sellers s ON s.id = sb.seller_id
        WHERE s.id = $1;
		`

	sellerBalanceView := &e.SellerBalanceView{}
	err = db.QueryRowContext(ctx, query, sellerID).Scan(&sellerBalanceView.SellerID, &sellerBalanceView.SellerName,
		&sellerBalanceView.SellerBalance, &sellerBalanceView.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %v", ErrNotFound, err)
		}
		if errors.Is(err, sql.ErrConnDone) {
			return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
		}
		if errors.Is(err, sql.ErrTxDone) {
			return nil, fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
		}
		return nil, fmt.Errorf("%w: %v", ErrFailedToScanRow, err)
	}

	return sellerBalanceView, nil
}
