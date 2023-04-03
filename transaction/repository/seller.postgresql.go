package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ralphbaer/hubla/common"
	e "github.com/Ralphbaer/hubla/transaction/entity"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type SellerPostgresRepository struct {
	connection *common.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewSellerPostgreSQLRepository(c *common.PostgresConnection) *SellerPostgresRepository {
	return &SellerPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *SellerPostgresRepository) Save(ctx context.Context, s *e.Seller) (string, error) {
	db, err := r.connection.Connect()
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrFailedToBeginTransaction, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO sellers(id, name, seller_type, created_at) VALUES ($1, $2, $3, DEFAULT) RETURNING id`
	var sellerID string
	if err := tx.QueryRowContext(ctx, query, s.ID, s.Name, e.SellerTypeMap[s.SellerType]).Scan(&sellerID); err != nil {
		return "", fmt.Errorf("%w: %v", ErrFailedToInsertSeller, err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("%w: %v", ErrFailedToCommitTransaction, err)
	}

	return "", nil
}

func (r *SellerPostgresRepository) Find(ctx context.Context, sellerName string) (*e.Seller, error) {
	db, err := r.connection.Connect()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFailedToConnectToDatabase, err)
	}

	var seller e.Seller
	if err := db.QueryRow(`
		 SELECT *
		 FROM sellers
		 WHERE name = $1`, sellerName).Scan(&seller.ID, &seller.Name, &seller.SellerType, &seller.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Seller not found
		}
		return nil, err // Other error
	}

	return &seller, nil
}
