// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gen

import (
	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hpostgres"
	"github.com/Ralphbaer/hubla/backend/transaction/app"
	"github.com/Ralphbaer/hubla/backend/transaction/handler"
	"github.com/Ralphbaer/hubla/backend/transaction/repository"
	"github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

// Injectors from inject.go:

// InitializeApp setup the dependencies and returns a new *app.App instance
func InitializeApp() *app.App {
	config := app.NewConfig()
	postgresConnection := setupPostgreSQLConnection(config)
	sellerBalancePostgresRepository := repository.NewSellerBalancePostgreSQLRepository(postgresConnection)
	sellerPostgresRepository := repository.NewSellerPostgreSQLRepository(postgresConnection)
	sellerUseCase := &usecase.SellerUseCase{
		SellerBalanceRepo: sellerBalancePostgresRepository,
		SellerRepo:        sellerPostgresRepository,
	}
	sellerHandler := &handler.SellerHandler{
		UseCase: sellerUseCase,
	}
	fileMetadataPostgresRepository := repository.NewFileMetadataPostgreSQLRepository(postgresConnection)
	productPostgresRepository := repository.NewProductPostgreSQLRepository(postgresConnection)
	transactionPostgresRepository := repository.NewTransactionPostgreSQLRepository(postgresConnection)
	fileTransactionPostgresRepository := repository.NewFileTransactionPostgreSQLRepository(postgresConnection)
	transactionUseCase := &usecase.TransactionUseCase{
		FileMetadataRepo:    fileMetadataPostgresRepository,
		SellerRepo:          sellerPostgresRepository,
		ProductRepo:         productPostgresRepository,
		TransactionRepo:     transactionPostgresRepository,
		FileTransactionRepo: fileTransactionPostgresRepository,
		SellerBalanceRepo:   sellerBalancePostgresRepository,
	}
	transactionHandler := &handler.TransactionHandler{
		UseCase: transactionUseCase,
	}
	router := app.NewRouter(sellerHandler, transactionHandler)
	server := app.NewServer(config, router)
	appApp := &app.App{
		Server: server,
	}
	return appApp
}

// inject.go:

var onceConfig sync.Once

func setupPostgreSQLConnection(cfg *app.Config) *hpostgres.PostgresConnection {
	return &hpostgres.PostgresConnection{
		ConnectionString: cfg.PostgreSQLConnectionString,
	}
}

var applicationSet = wire.NewSet(common.InitLocalEnvConfig, setupPostgreSQLConnection, app.NewConfig, app.NewRouter, app.NewServer, repository.NewTransactionPostgreSQLRepository, repository.NewSellerPostgreSQLRepository, repository.NewSellerBalancePostgreSQLRepository, repository.NewProductPostgreSQLRepository, repository.NewFileMetadataPostgreSQLRepository, repository.NewFileTransactionPostgreSQLRepository, wire.Struct(new(usecase.TransactionUseCase), "*"), wire.Struct(new(usecase.SellerUseCase), "*"), wire.Struct(new(handler.TransactionHandler), "*"), wire.Struct(new(handler.SellerHandler), "*"), wire.Bind(new(repository.TransactionRepository), new(*repository.TransactionPostgresRepository)), wire.Bind(new(repository.SellerRepository), new(*repository.SellerPostgresRepository)), wire.Bind(new(repository.SellerBalanceRepository), new(*repository.SellerBalancePostgresRepository)), wire.Bind(new(repository.ProductRepository), new(*repository.ProductPostgresRepository)), wire.Bind(new(repository.FileMetadataRepository), new(*repository.FileMetadataPostgresRepository)), wire.Bind(new(repository.FileTransactionRepository), new(*repository.FileTransactionPostgresRepository)), wire.Bind(new(http.Handler), new(*mux.Router)))
