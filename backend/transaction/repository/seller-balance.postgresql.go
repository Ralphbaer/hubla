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

// SellerBalancePostgresRepository represents a MongoDB implementation of PartnerRepository interface
type SellerBalancePostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSellerBalancePostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewSellerBalancePostgreSQLRepository(c *hpostgres.PostgresConnection) *SellerBalancePostgresRepository {
	return &SellerBalancePostgresRepository{
		connection: c,
	}
}

// Upsert upserts a SellerBalance entity in the Postgres database with the given values.
// If the entity already exists, it updates its balance value by adding the new value to the existing one.
// It returns the new balance value after the upsert operation is performed.
func (r *SellerBalancePostgresRepository) Upsert(ctx context.Context, p *e.SellerBalance) (*float64, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	success := false
	defer func() {
		if !success {
			if err := tx.Rollback(); err != nil {
				hlog.NewLoggerFromContext(ctx).Errorf("Failed to rollback transaction: %v", err)
			}
		}
	}()

	query := `INSERT INTO seller_balance (id, seller_id, balance, updated_at, created_at)
              VALUES ($1, $2, $3, $4, DEFAULT) 
              ON CONFLICT (seller_id) DO UPDATE SET balance = seller_balance.balance + $3
              RETURNING balance`
	var newBalance float64
	if err := tx.QueryRowContext(ctx, query, p.ID, p.SellerID, p.Balance, p.UpdatedAt).Scan(&newBalance); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &newBalance, nil
}

// Find retrieves a SellerBalanceView entity from the Postgres database by the given seller ID.
// It returns a SellerBalanceView object containing the seller ID, name, balance, and updated time.
func (r *SellerBalancePostgresRepository) Find(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT s.id, s.name, s.seller_type, s.created_at, sb.balance, sb.updated_at
        FROM seller_balance sb
        JOIN seller s ON s.id = sb.seller_id
        WHERE s.id = $1;
		`

	sellerBalanceView := &e.SellerBalanceView{}
	err = db.QueryRowContext(ctx, query, sellerID).Scan(&sellerBalanceView.SellerID, &sellerBalanceView.SellerName,
		&sellerBalanceView.SellerType, &sellerBalanceView.SellerCreatedAt, &sellerBalanceView.SellerBalance, &sellerBalanceView.SellerBalanceUpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.SellerBalance{}).Name(),
				Err:        err,
			}
		}
		return nil, err
	}

	return sellerBalanceView, nil
}
