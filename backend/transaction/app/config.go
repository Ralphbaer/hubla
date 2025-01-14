package app

import "github.com/Ralphbaer/hubla/backend/common"

// Config is the top level configuration struct for entire application
type Config struct {
	EnvName                        string `env:"ENV_NAME"`
	PostgreSQLConnectionString     string `env:"POSTGRES_CONNECTION_STRING"`
	PostgreSQLConnectionStringTest string `env:"POSTGRES_SQL_CONNECTION_TEST"`
	AccessTokenPublicKey           string `env:"ACCESS_TOKEN_PUBLIC_KEY"`
	ServerAddress                  string `env:"SERVER_ADDRESS"`
	SpecURL                        string `env:"SPEC_URL"`
}

// NewConfig creates a instance of Config
func NewConfig() *Config {
	cfg := &Config{}
	if err := common.SetConfigFromEnvVars(cfg); err != nil {
		panic(err)
	}
	return cfg
}
