package repository

import (
	"context"
	"database/sql"
	"reflect"

	e "github.com/Ralphbaer/hubla/backend/auth/entity"
	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type UserPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewSalesPostgreSQLRepository creates an instance of repository.SalesPostgreSQLRepository
func NewUserPostgreSQLRepository(c *hpostgres.PostgresConnection) *UserPostgresRepository {
	return &UserPostgresRepository{
		connection: c,
	}
}

func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*e.User, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, name, email, role, created_at, updated_at
        FROM users
        WHERE email = $1`
	var user e.User
	err = db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.User{}).Name(),
				Err:        err,
			}
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgresRepository) FindByID(ctx context.Context, ID string) (*e.User, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, name, email, role, created_at, updated_at
        FROM users
        WHERE id = $1`
	var user e.User
	err = db.QueryRowContext(ctx, query, ID).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(e.User{}).Name(),
				Err:        err,
			}
		}
		return nil, err
	}

	return &user, nil
}
