package app

import (
	"time"

	"github.com/Ralphbaer/hubla/backend/common"
)

// Config is the top level configuration struct for entire application
type Config struct {
	EnvName                        string `env:"ENV_NAME"`
	PostgreSQLConnectionString     string `env:"POSTGRES_CONNECTION_STRING"`
	PostgreSQLConnectionStringTest string `env:"POSTGRES_SQL_CONNECTION_TEST"`
	ServerAddress                  string `env:"SERVER_ADDRESS"`
	HydraAdminURL                  string `env:"HYDRA_ADMIN_URL"`
	SpecURL                        string `env:"SPEC_URL"`

	Origin                 string        `env:"ORIGIN"`
	AccessTokenPrivateKey  string        `env:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `env:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `env:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `env:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `env:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `env:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `env:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `env:"REFRESH_TOKEN_MAXAGE"`
}

// NewConfig creates a instance of Config
func NewConfig() *Config {
	cfg := &Config{}
	common.SetConfigFromEnvVars(cfg)
	return cfg
}
