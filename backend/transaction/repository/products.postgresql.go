package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/lib/pq"
)

// ProductPostgresRepository is a struct that implements the ProductRepository interface
// for retrieving and saving product data from a PostgreSQL database.
type ProductPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewProductPostgreSQLRepository creates a new instance of the
// ProductPostgresRepository, which implements the ProductRepository interface
// for retrieving and saving product data from a PostgreSQL database.
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

	query := `INSERT INTO product(id, name, creator_id, created_at) VALUES ($1, $2, $3, DEFAULT)`
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

// FindByProductName retrieves a Product entity from the Postgres database by the given product name.
// It returns a Product object if found, or nil if no product is found.
func (r *ProductPostgresRepository) FindByProductName(ctx context.Context, productName string) (*e.Product, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT *
        FROM product
        WHERE name = $1`
	rows, err := db.QueryContext(ctx, query, productName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product *e.Product
	if rows.Next() {
		product := &e.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.CreatorID, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
