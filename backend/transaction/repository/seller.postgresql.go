package repository

import (
	"context"
	"database/sql"

	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
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
func (r *SellerPostgresRepository) Save(ctx context.Context, s *e.Seller) (*string, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, hpostgres.WithError(err)
	}
	defer tx.Rollback()

	query := `INSERT INTO sellers(id, name, seller_type, created_at) VALUES ($1, $2, $3, DEFAULT) RETURNING id`
	var sellerID string
	if err := tx.QueryRowContext(ctx, query, s.ID, s.Name, e.SellerTypeMap[s.SellerType]).Scan(&sellerID); err != nil {
		return nil, hpostgres.WithError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, hpostgres.WithError(err)
	}

	return &sellerID, nil
}

func (r *SellerPostgresRepository) Find(ctx context.Context, sellerName string) (*e.Seller, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	var seller e.Seller
	query := `SELECT id, name, seller_type, created_at FROM sellers WHERE name = $1`
	err = db.QueryRowContext(ctx, query, sellerName).Scan(&seller.ID, &seller.Name, &seller.SellerType, &seller.CreatedAt)
	if err != nil {
		return nil, hpostgres.WithError(err)
	}

	return &seller, nil
}
