package repository

import (
	"context"

	"github.com/Ralphbaer/hubla/common"
	e "github.com/Ralphbaer/hubla/sales/entity"

	"github.com/lib/pq"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type SalesPostgresRepository struct {
	connection *common.PostgresConnection
}

/*
func connectToPostgres() (*sql.DB, error) {
	connStr := "user=username password=password dbname=database host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
*/

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewSalesPostgreSQLRepository(c *common.PostgresConnection) *SalesPostgresRepository {
	return &SalesPostgresRepository{
		connection: c,
	}
}

// Save stores the given entity.Sales into PostgreSQL
func (r *SalesPostgresRepository) Save(ctx context.Context, s []e.Sales) (*string, error) {
	//	ctx := context.Background()

	db, err := r.connection.Connect()
	if err != nil {
		return nil, err
	}

	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := txn.Prepare(pq.CopyIn("transactions", "id", "t_type", "t_date", "product_description", "amount", "seller"))
	if err != nil {
		return nil, err
	}

	for _, s := range s {
		_, err = stmt.Exec(s.ID, s.TType, s.TDate, s.ProductDescription, s.Amount, s.Seller)
		if err != nil {
			return nil, err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	/*

		query := `INSERT INTO transactions (id, t_type, t_date, product_description, amount, seller)
		          VALUES ($1, $2, $3, $4, $5, $6)`
		_, err = tx.ExecContext(ctx, query, id, s.TType, s.TDate, s.ProductDescription, s.Amount, s.Seller)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("could not execute insert: %v", err)
		}

		if err = tx.Commit(); err != nil {
			return nil, fmt.Errorf("could not commit transaction: %v", err)
		}

		s.ID = id*/

	return nil, nil
}
