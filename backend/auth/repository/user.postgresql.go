package repository

import (
	"context"
	"database/sql"
	"reflect"

	e "github.com/Ralphbaer/hubla/backend/auth/entity"
	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
)

// UserPostgresRepository implements the UserRepository interface for
// retrieving user data from a PostgreSQL database.
type UserPostgresRepository struct {
	connection *hpostgres.PostgresConnection
}

// NewUserPostgreSQLRepository creates a new UserPostgresRepository instance
// with a given PostgreSQL database connection.
func NewUserPostgreSQLRepository(c *hpostgres.PostgresConnection) *UserPostgresRepository {
	return &UserPostgresRepository{
		connection: c,
	}
}

// FindByEmail retrieves a user by email from the database.
func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*e.User, error) {
	db, err := r.connection.GetDB()
	if err != nil {
		return nil, err
	}

	query := `
        SELECT id, name, email, password, role, created_at, updated_at
        FROM user_account
        WHERE email = $1`
	var user e.User
	err = db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
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

// FindByID retrieves a user by ID from the database.
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
