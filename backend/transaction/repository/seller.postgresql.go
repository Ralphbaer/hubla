package repository

import (
	"context"
	"database/sql"
	"reflect"

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
	if _, err := tx.ExecContext(ctx, query, s.ID, s.Name, e.SellerTypeMap[s.SellerType]); err != nil {
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

func (r *SellerPostgresRepository) Find(ctx context.Context, sellerName string) (*e.Seller, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	var seller e.Seller
	query := `SELECT id, name, seller_type, created_at FROM seller WHERE name = $1`
	err = db.QueryRowContext(ctx, query, sellerName).Scan(&seller.ID, &seller.Name, &seller.SellerType, &seller.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.Seller{}).Name(),
				Err:        err,
			}
		}
		return nil, err
	}

	return &seller, nil
}
