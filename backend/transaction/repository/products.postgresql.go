package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
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
		return hpostgres.WithError(err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return hpostgres.WithError(err)
	}
	defer tx.Rollback()

	query := `INSERT INTO products(id, name, creator_id, created_at) VALUES ($1, $2, $3, DEFAULT) RETURNING id`
	var productID string
	err = tx.QueryRowContext(ctx, query, s.ID, s.Name, s.CreatorID).Scan(&productID)

	if err != nil {
		return hpostgres.WithError(err)
	}

	err = tx.Commit()
	if err != nil {
		return hpostgres.WithError(err)
	}

	return nil
}

func (r *ProductPostgresRepository) Find(ctx context.Context, productName string) (*e.Product, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	query := `
        SELECT id, name, creator_id, created_at
        FROM products
        WHERE name = $1`
	var product e.Product
	err = db.QueryRowContext(ctx, query, productName).Scan(&product.ID, &product.Name, &product.CreatorID, &product.CreatedAt)
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	return &product, nil
}
