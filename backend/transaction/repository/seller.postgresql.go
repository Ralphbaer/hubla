package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hlog"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// SellerPostgresRepository implements the SellerRepository interface for storing and retrieving sellers in Postgres.
type SellerPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSellerPostgreSQLRepository creates a new instance of SellerPostgresRepository with the given Postgres connection.
func NewSellerPostgreSQLRepository(c *hpostgres.PostgresConnection) *SellerPostgresRepository {
	return &SellerPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *SellerPostgresRepository) Save(ctx context.Context, s *e.Seller) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	success := false
	defer func() {
		if !success {
			if err := tx.Rollback(); err != nil {
				hlog.NewLoggerFromContext(ctx).Errorf("Failed to rollback transaction: %v", err)
			}
		}
	}()

	query := `INSERT INTO seller(id, name, seller_type, created_at) VALUES ($1, $2, $3, DEFAULT)`
	if _, err := tx.ExecContext(ctx, query, s.ID, s.Name, e.SellerTypeMapString[s.SellerType]); err != nil {
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

// FindBySellerName retrieves a Seller entity from the Postgres database by the given seller name.
// It returns a Seller object if found, or nil if no seller is found.
func (r *SellerPostgresRepository) FindBySellerName(ctx context.Context, sellerName string) (*e.Seller, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM seller WHERE name = $1`
	rows, err := db.QueryContext(ctx, query, sellerName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seller *e.Seller
	if rows.Next() {
		seller := &e.Seller{}
		var sellerTypeStr string
		err := rows.Scan(&seller.ID, &seller.Name, &sellerTypeStr, &seller.CreatedAt)
		if err != nil {
			return nil, err
		}
		seller.SellerType = e.SellerTypeFromString[sellerTypeStr]
	}

	return seller, nil
}
