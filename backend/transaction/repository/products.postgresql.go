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
type ProductPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewProductPostgreSQLRepository(c *hpostgres.PostgresConnection) *ProductPostgresRepository {
	return &ProductPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *ProductPostgresRepository) Save(ctx context.Context, s *e.Product) error {
	db, err := r.connection.GetDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO products(id, name, creator_id, created_at) VALUES ($1, $2, $3, DEFAULT)`
	if _, err = tx.ExecContext(ctx, query, s.ID, s.Name, s.CreatorID); err != nil {
		if pqerr := err.(*pq.Error); pqerr.Code == "23505" {
			return common.EntityConflictError{
				Message: err.Error(),
				Err:     err,
			}
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductPostgresRepository) Find(ctx context.Context, productName string) (*e.Product, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, name, creator_id, created_at
        FROM products
        WHERE name = $1`
	var product e.Product
	err = db.QueryRowContext(ctx, query, productName).Scan(&product.ID, &product.Name, &product.CreatorID, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.Seller{}).Name(),
				Err:        err,
			}
		}
		return nil, err
	}

	return &product, nil
}
