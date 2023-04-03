//go:build wireinject
// +build wireinject

//golint:ignore

package gen

import (
	"net/http"
	"sync"

	"github.com/Ralphbaer/hubla/common"
	"github.com/Ralphbaer/hubla/sales/app"
	h "github.com/Ralphbaer/hubla/sales/handler"
	r "github.com/Ralphbaer/hubla/sales/repository"
	uc "github.com/Ralphbaer/hubla/sales/usecase"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

var onceConfig sync.Once

func setupPostgreSQLConnection(cfg *app.Config) *common.PostgresConnection {
	return &common.PostgresConnection{
		ConnectionString: cfg.PostgreSQLConnectionString,
	}
}

var applicationSet = wire.NewSet(
	common.InitLocalEnvConfig,
	setupPostgreSQLConnection,
	app.NewConfig,
	app.NewRouter,
	app.NewServer,
	r.NewTransactionPostgreSQLRepository,
	r.NewSellerPostgreSQLRepository,
	r.NewSellerBalancePostgreSQLRepository,
	r.NewProductPostgreSQLRepository,
	wire.Struct(new(uc.SalesUseCase), "*"),
	wire.Struct(new(uc.SellerUseCase), "*"),
	wire.Struct(new(h.SalesHandler), "*"),
	wire.Bind(new(r.TransactionRepository), new(*r.TransactionPostgresRepository)),
	wire.Bind(new(r.SellerRepository), new(*r.SellerPostgresRepository)),
	wire.Bind(new(r.SellerBalanceRepository), new(*r.SellerBalancePostgresRepository)),
	wire.Bind(new(r.ProductRepository), new(*r.ProductPostgresRepository)),
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
