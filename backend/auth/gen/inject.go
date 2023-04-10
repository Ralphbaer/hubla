//go:build wireinject
// +build wireinject

//golint:ignore

package gen

import (
	"net/http"
	"sync"

	"github.com/Ralphbaer/hubla/backend/auth/app"
	h "github.com/Ralphbaer/hubla/backend/auth/handler"
	r "github.com/Ralphbaer/hubla/backend/auth/repository"
	uc "github.com/Ralphbaer/hubla/backend/auth/usecase"
	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	"github.com/Ralphbaer/hubla/backend/common/jwt"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

var onceConfig sync.Once

func setupPostgreSQLConnection(cfg *app.Config) *hpostgres.PostgresConnection {
	return &hpostgres.PostgresConnection{
		ConnectionString: cfg.PostgreSQLConnectionString,
	}
}

func setupJWTAuth(cfg *app.Config) *jwt.JWTAuth {
	return &jwt.JWTAuth{
		AccessTokenPrivateKey: cfg.AccessTokenPrivateKey,
		AccessTokenPublicKey:  cfg.AccessTokenPublicKey,
	}
}

var applicationSet = wire.NewSet(
	common.InitLocalEnvConfig,
	setupPostgreSQLConnection,
	setupJWTAuth,
	app.NewConfig,
	app.NewRouter,
	app.NewServer,
	r.NewUserPostgreSQLRepository,
	wire.Struct(new(h.LoginHandler), "*"),
	wire.Struct(new(uc.UserUseCase), "*"),
	wire.Bind(new(r.UserRepository), new(*r.UserPostgresRepository)),
	wire.Bind(new(http.Handler), new(*mux.Router)),
)

// InitializeApp setup the dependencies and returns a new *app.App instance
func InitializeApp() *app.App {
	wire.Build(
		applicationSet,
		wire.Struct(new(app.App), "*"),
	)
	return nil
}
