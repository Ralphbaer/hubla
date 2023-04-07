package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type SellerPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
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
	defer tx.Rollback()

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

func (r *SellerPostgresRepository) FindBySellerName(ctx context.Context, sellerName string) (*e.Seller, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `SELECT id, name, seller_type, created_at FROM seller WHERE name = $1`
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
