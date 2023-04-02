package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/common"
	e "github.com/Ralphbaer/hubla/sales/entity"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type ProductPostgresRepository struct {
	connection *common.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewProductPostgreSQLRepository(c *common.PostgresConnection) *ProductPostgresRepository {
	return &ProductPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *ProductPostgresRepository) Save(ctx context.Context, s *e.Product) error {
	db, err := r.connection.Connect()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO products(id, name, creator_id, created_at) VALUES ($1, $2, $3, DEFAULT) RETURNING id`
	var productID string
	if err := tx.QueryRowContext(ctx, query, s.ID, s.Name, s.CreatorID).Scan(&productID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
