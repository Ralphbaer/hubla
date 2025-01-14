//go:build wireinject
// +build wireinject

//golint:ignore

package gen

import (
	"net/http"
	"sync"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hlogrus"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	"github.com/Ralphbaer/hubla/backend/transaction/app"
	h "github.com/Ralphbaer/hubla/backend/transaction/handler"
	r "github.com/Ralphbaer/hubla/backend/transaction/repository"
	uc "github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

var onceConfig sync.Once

func setupPostgreSQLConnection(cfg *app.Config) *hpostgres.PostgresConnection {
	return &hpostgres.PostgresConnection{
		ConnectionString: cfg.PostgreSQLConnectionString,
	}
}

var applicationSet = wire.NewSet(
	common.InitLocalEnvConfig,
	hlogrus.InitializeLogger,
	setupPostgreSQLConnection,
	app.NewConfig,
	app.NewRouter,
	app.NewServer,
	r.NewTransactionPostgreSQLRepository,
	r.NewSellerPostgreSQLRepository,
	r.NewSellerBalancePostgreSQLRepository,
	r.NewProductPostgreSQLRepository,
	r.NewFileMetadataPostgreSQLRepository,
	r.NewFileTransactionPostgreSQLRepository,
	wire.Struct(new(uc.TransactionUseCase), "*"),
	wire.Struct(new(uc.SellerUseCase), "*"),
	wire.Struct(new(h.TransactionHandler), "*"),
	wire.Struct(new(h.SellerHandler), "*"),
	wire.Bind(new(r.TransactionRepository), new(*r.TransactionPostgresRepository)),
	wire.Bind(new(r.SellerRepository), new(*r.SellerPostgresRepository)),
	wire.Bind(new(r.SellerBalanceRepository), new(*r.SellerBalancePostgresRepository)),
	wire.Bind(new(r.ProductRepository), new(*r.ProductPostgresRepository)),
	wire.Bind(new(r.FileMetadataRepository), new(*r.FileMetadataPostgresRepository)),
	wire.Bind(new(r.FileTransactionRepository), new(*r.FileTransactionPostgresRepository)),
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
